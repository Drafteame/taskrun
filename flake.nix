{
  description = "taskrun - create jobs from configurable commands for different stages";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        taskrunVersion = "0.1.5";
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = "taskrun";
          version = taskrunVersion;

          subPackages = [ "." ];

          src = ./.;

          ldflags = [
            "-s"
            "-w"
            "-X 'github.com/Drafteame/taskrun/cmd/commands.Version=nix-v${taskrunVersion}'"
          ];

          env.CGO_ENABLED = false;
          env.GOWORK = "off";

          vendorHash = "sha256-xVXGIEgPvzSUI+nxNe7Q6mBLTA6q5548rMlOH5iNFGg=";

          meta = {
            description = "Simple cli that helps to create jobs from configurable commands for different stages";
            mainProgram = "taskrun";
          };
        };

        apps.default = flake-utils.lib.mkApp {
          drv = self.packages.${system}.default;
        };
      }
    );
}
