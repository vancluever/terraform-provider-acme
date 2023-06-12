{
  description = "terraform-provider-acme: project development environment";

  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; };
      in {
        inherit pkgs;
        devShell = pkgs.mkShell {
          packages = [
            pkgs.buf
            pkgs.buf-language-server
            pkgs.go
            pkgs.golangci-lint
            pkgs.golangci-lint-langserver
            pkgs.gopls
            pkgs.nixfmt
            pkgs.protoc-gen-go
            pkgs.protoc-gen-go-grpc
            pkgs.protoc-gen-go-grpc
            pkgs.rnix-lsp
          ];
        };
      });
}
