{ pkgs, version, ... }:
pkgs.buildGoModule {
  pname = "advent-of-code";
  version = version;
  src = ./..;
  vendorHash = "sha256-+PvrStpf2oKT/IhF4nhCQzxMJEC7+PaCcJ1xmKQGlDM=";
  doCheck = true;
  subPackages = [ "cmd/advent-of-code" ];
}
