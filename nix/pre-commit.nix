{
  pkgs,
  preCommitHooks,
  shellHook,
}:

let
  deps =
    (pkgs.buildGoModule {
      pname = "advent-of-code-modules";
      version = "0.0.1";
      src = ../.;
      proxyVendor = true;
      vendorHash = "sha256-RWKQ2VraO1SJCqATCtNDIKmonn+KKm4xurkC+cHRc+A=";
      buildInputs = [ pkgs.lp_solve ];
    }).goModules;

  goWithProxy = pkgs.writeShellScriptBin "go" ''
    export GOPROXY="file://${deps}"
    export GOSUMDB="off"
    exec ${pkgs.go}/bin/go "$@"
  '';

  cgoSetupHook = pkgs.makeSetupHook {
    name = "cgo-setup-hook";
  } (pkgs.writeScript "cgo-setup-hook.sh" shellHook);
in
preCommitHooks.run {
  src = ../.;
  hooks = {
    nixfmt-rfc-style.enable = true;
    gofmt.enable = true;
    golangci-lint = {
      enable = true;
      extraPackages = [ goWithProxy ];
    };
    golines = {
      enable = true;
      settings.flags = "-m 100 --dry-run";
    };
    gotest = {
      enable = true;
      extraPackages = [
        goWithProxy
        cgoSetupHook
      ];
    };
    govet = {
      enable = true;
      extraPackages = [
        goWithProxy
        cgoSetupHook
      ];
    };
  };
}
