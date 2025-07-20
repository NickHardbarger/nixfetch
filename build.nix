{
  buildGoModule,
}:

buildGoModule {
  pname = "nixfetch";
  version = "0.1";

  src = builtins.path {
    name = "build";
    path = ./.;
  };
}
