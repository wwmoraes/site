{
  fetchurl,
  fetchgit,
  linkFarm,
  runCommand,
  gnutar,
}:
rec {
  offline_cache = linkFarm "offline" packages;
  packages = [
    {
      name = "blake3_wasm___blake3_wasm_2.1.5.tgz";
      path = fetchurl {
        name = "blake3_wasm___blake3_wasm_2.1.5.tgz";
        url = "https://registry.yarnpkg.com/blake3-wasm/-/blake3-wasm-2.1.5.tgz";
        sha512 = "F1+K8EbfOZE49dtoPtmxUQrpXaBIl3ICvasLh+nJta0xkz+9kF/7uet9fLnwKqhDrmj6g+6K3Tw9yQPUg2ka5g==";
      };
    }
  ];
}
