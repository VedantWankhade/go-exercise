let
    nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-23.11";
    pkgs = import nixpkgs { config = {}; overlays = []; };
in

pkgs.mkShell {
    packages = with pkgs; [
        go
        gotools
        go-tools
    ];

    shellHook = ''
        echo "============= go-exercise ==================="
        echo "Entered go-exercise environment."
        echo "Dependencies"
        go version
        echo "------------------------"
        echo "Git status"
        git status
    '';
}

