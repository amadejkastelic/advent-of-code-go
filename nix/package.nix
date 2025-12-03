{ pkgs, version, ... }:
pkgs.buildGoModule {
  pname = "advent-of-code";
  version = version;
  src = ./..;
  vendorHash = "sha256-zdGqpZ8hPk65XQ73hV9pTf/HSMrU29/WjMECQbg70J8=";
  doCheck = true;
  subPackages = [ "cmd/advent-of-code" ];
}
