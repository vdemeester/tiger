with import <nixpkgs> {};
stdenv.mkDerivation  {
    name = "tiger";
    buildInputs = [
        pkgs.go_1_10
        pkgs.vndr
        pkgs.gnumake
	      pkgs.gotools
      ];
}
