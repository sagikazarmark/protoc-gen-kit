# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

# Main targets
include main.mk

PROTOC_VERSION = 3.12.2

bin/protoc: bin/protoc-${PROTOC_VERSION}
	@ln -sf protoc-${PROTOC_VERSION}/bin/protoc bin/protoc
bin/protoc-${PROTOC_VERSION}:
	@mkdir -p bin/protoc-${PROTOC_VERSION}
ifeq (${OS}, darwin)
	curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-osx-x86_64.zip > bin/protoc.zip
endif
ifeq (${OS}, linux)
	curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip > bin/protoc.zip
endif
	unzip bin/protoc.zip -d bin/protoc-${PROTOC_VERSION}
	rm bin/protoc.zip

bin/protoc-gen-go: go.mod
	@mkdir -p bin
	go build -o bin/protoc-gen-go google.golang.org/protobuf/cmd/protoc-gen-go

bin/protoc-gen-go-grpc: gotools.mod
	@mkdir -p bin
	go build -modfile gotools.mod -o bin/protoc-gen-go-grpc google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: testproto
testproto: bin/protoc bin/protoc-gen-go bin/protoc-gen-go-grpc build
	protoc -I bin/protoc-${PROTOC_VERSION} -I test --plugin=build/protoc-gen-kit --go_out=paths=source_relative:test/ --go-grpc_out=paths=source_relative:test/ --kit_out=paths=source_relative:test/ test/test.proto test/subtest/subtest.proto
