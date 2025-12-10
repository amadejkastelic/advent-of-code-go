{ pkgs, version, ... }:
pkgs.buildGoModule {
  pname = "advent-of-code";
  version = version;
  src = ./..;
  vendorHash = "sha256-AR7Nb4Nmi+clGkAnY5HOnCeIfNW2txfJ7j66iaVxzyE=";
  doCheck = true;
  buildInputs = [ pkgs.lp_solve ];
  subPackages = [ "cmd/advent-of-code" ];
}
