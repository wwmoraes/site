#!/usr/bin/env node

import process from 'node:process';
import { readFileSync } from "node:fs";
import { extname } from "node:path";
import { hash as blake3hash } from "blake3-wasm";

if (process.argv.length < 3) {
  console.log("usage: cfhash <FILE-1>[..<FILE-N]");
  process.exit(2);
}

process.argv.slice(2).forEach(filepath => {
  const contents = readFileSync(filepath);
  const base64Contents = contents.toString("base64");
  const extension = extname(filepath).substring(1);
  const hash = blake3hash(base64Contents + extension)
    .toString("hex")
    .slice(0, 32);

  console.log("%s\tbase64:%s", filepath, base64Contents);
  console.log("%s\t%s", hash, filepath);
});
