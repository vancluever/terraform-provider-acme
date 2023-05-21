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
            pkgs.go
            pkgs.nixfmt
            pkgs.buf
            pkgs.buf-language-server
            pkgs.protoc-gen-go
            pkgs.protoc-gen-go-grpc
          ];
        };
      });
}
