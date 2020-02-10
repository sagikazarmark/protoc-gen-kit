package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"emperror.dev/errors"
	"github.com/dave/jennifer/jen"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func main() {
	err := Main(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

// Main does the hard work. It is called by the main func.
func Main(in io.Reader, out io.Writer) error {
	req, err := readCodeGeneratorRequest(in)
	if err != nil {
		return err
	}

	param := req.GetParameter()

	params := make(map[string]string)
	for _, p := range strings.Split(param, ",") {
		if i := strings.Index(p, "="); i < 0 {
			params[p] = ""
		} else {
			params[p[0:i]] = p[i+1:]
		}
	}

	pluginList := "none" // Default list of plugin names to enable (empty means all).

	for k, v := range params {
		// nolint: gocritic
		switch k {
		case "plugins":
			pluginList = v
		}
	}

	_ = pluginList

	resp := new(plugin.CodeGeneratorResponse)

	for _, file := range req.GetProtoFile() {
		if len(file.GetService()) < 1 {
			continue
		}

		respFile, err := generateFile(file)
		if err != nil {
			resp.Error = proto.String(err.Error())

			break
		}

		resp.File = append(resp.File, respFile)
	}

	err = writeCodeGeneratorResponse(out, resp)
	if err != nil {
		return err
	}

	return nil
}

func readCodeGeneratorRequest(in io.Reader) (*plugin.CodeGeneratorRequest, error) {
	data, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read input")
	}

	req := new(plugin.CodeGeneratorRequest)
	if err = proto.Unmarshal(data, req); err != nil {
		return nil, errors.Wrap(err, "cannot parse input proto")
	}

	if len(req.FileToGenerate) == 0 {
		return nil, errors.Wrap(err, "no files to generate")
	}

	return req, nil
}

func writeCodeGeneratorResponse(out io.Writer, resp *plugin.CodeGeneratorResponse) error {
	data, err := proto.Marshal(resp)
	if err != nil {
		return errors.Wrap(err, "cannot serialize output proto")
	}

	_, err = out.Write(data)
	if err != nil {
		return errors.Wrap(err, "cannot write output")
	}

	return nil
}

func generateFile(file *descriptor.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error) {
	code := jen.NewFile(file.GetOptions().GetGoPackage())

	code.ImportAlias("github.com/go-kit/kit/transport/grpc", "kitgrpc")

	for _, service := range file.GetService() {
		fields := []jen.Code{jen.Op("*").Id(fmt.Sprintf("Unimplemented%sServer", service.GetName())), jen.Line()}

		for _, method := range service.GetMethod() {
			fields = append(fields, jen.Id(method.GetName()+"Handler").Qual("github.com/go-kit/kit/transport/grpc", "Handler"))
		}

		serverName := fmt.Sprintf("%sKitServer", service.GetName())
		code.Type().Id(serverName).Struct(fields...)

		for _, method := range service.GetMethod() {
			if !strings.HasPrefix(method.GetInputType(), "."+file.GetPackage()+".") {
				return nil, errors.New("input type is not within the same package as the service")
			}

			if !strings.HasPrefix(method.GetOutputType(), "."+file.GetPackage()+".") {
				return nil, errors.New("output type is not within the same package as the service")
			}

			responseType := strings.TrimPrefix(method.GetOutputType(), "."+file.GetPackage()+".")
			const receiver = "s"

			code.Func().Params(jen.Id(receiver).Id(serverName)).Id(method.GetName()).Params(
				jen.Id("ctx").Qual("context", "Context"),
				jen.Id("req").Op("*").Id(strings.TrimPrefix(method.GetInputType(), "."+file.GetPackage()+".")),
			).Params(
				jen.Op("*").Id(responseType),
				jen.Error(),
			).Block(
				jen.Id("_").Op(",").Id("resp").Op(",").Err().Op(":=").
					Id(receiver).Dot(method.GetName()+"Handler").Dot("ServeGRPC").Call(
					jen.Id("ctx"),
					jen.Id("req"),
				),
				jen.If(jen.Err().Op("!=").Nil()).Block(
					jen.Return(jen.Nil(), jen.Err()),
				),
				jen.Line(),
				jen.Return(
					jen.Id("resp").Assert(jen.Op("*").Id(responseType)),
					jen.Nil(),
				),
			)
		}
	}

	var buf bytes.Buffer

	err := code.Render(&buf)
	if err != nil {
		return nil, err
	}

	content, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, err
	}

	return &plugin.CodeGeneratorResponse_File{
		Name:    proto.String(fileName(file)),
		Content: proto.String(string(content)),
	}, nil
}

func fileName(file *descriptor.FileDescriptorProto) string {
	name := file.GetName()
	if ext := path.Ext(name); ext == ".proto" || ext == ".protodevel" {
		name = name[:len(name)-len(ext)]
	}

	name += ".kit.go"

	return name
}
