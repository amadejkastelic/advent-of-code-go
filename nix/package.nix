{ pkgs, version, ... }:
pkgs.buildGoModule {
  pname = "advent-of-code";
  version = version;
  src = ./..;
  vendorHash = "sha256-PW8ZywwJV3eYDQkohXRohjhAocWK+sj08S+RfOt18F4=";
  doCheck = true;
  buildInputs = [ pkgs.lp_solve ];
  subPackages = [ "cmd/advent-of-code" ];
}
