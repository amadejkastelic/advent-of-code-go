{ pkgs, version, ... }:
pkgs.buildGoModule {
  pname = "advent-of-code";
  version = version;
  src = ./..;
  vendorHash = "sha256-bhZ5o2XmM06rZtUmT7XGUIHiaLteeBaG+H3lP/JPwiQ=";
  doCheck = true;
  subPackages = [ "cmd/advent-of-code" ];
}
