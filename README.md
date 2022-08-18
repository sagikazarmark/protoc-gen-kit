> # ⚠ Project moved
> The project has been moved to a new place: https://github.com/sagikazarmark/protoc-gen-go-kit

# Go Kit Protoc Compiler

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/sagikazarmark/protoc-gen-kit/CI?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/sagikazarmark/protoc-gen-kit?style=flat-square)](https://goreportcard.com/report/github.com/sagikazarmark/protoc-gen-kit)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/sagikazarmark/protoc-gen-kit)
[![built with nix](https://img.shields.io/badge/builtwith-nix-7d81f7?style=flat-square)](https://builtwithnix.org)

**Protoc compiler for Go kit code.**


**IMPORTANT: This project is work in progress.**


## Installation

Download prebuilt binaries from the [releases](https://github.com/sagikazarmark/protoc-gen-kit/releases) page or install it from source:

```shell
go get github.com/sagikazarmark/protoc-gen-kit
```


## Development

When all coding and testing is done, please run the test suite:

```shell
make check
```

For the best developer experience, install [Nix](https://builtwithnix.org/) and [direnv](https://direnv.net/).

Alternatively, install Go manually or using a package manager. Install the rest of the dependencies by running `make deps`.


## License

The project is licensed under the [MIT License](LICENSE).
