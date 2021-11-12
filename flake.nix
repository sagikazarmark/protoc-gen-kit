{
  description = "Go Kit Protoc Compiler";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        buildDeps = with pkgs; [ git go gnumake ];
        devDeps = with pkgs;
          buildDeps ++ [
            golangci-lint
            gotestsum
            goreleaser
            protobuf
            protoc-gen-go
            protoc-gen-go-grpc
          ];
      in { devShell = pkgs.mkShell { buildInputs = devDeps; }; });
}
