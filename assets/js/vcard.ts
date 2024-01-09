const str2ab = (data: string): Uint8Array =>
  new Uint8Array(data.split("").map(char => char.charCodeAt(0)));

const ab2str = (bytes: Uint8Array): string =>
  Array.from(bytes).map(byte => String.fromCharCode(byte)).join("");

const gunzip = async (bits: (ArrayBuffer | ArrayBufferView) | string | Blob): Promise<string> =>
  (new Blob([bits], { type: " application/gzip; charset=binary" }))
    .stream()
    .pipeThrough(new DecompressionStream("gzip"))
    .getReader()
    .read()
    .then(d => d.value)
    .then(ab2str);

// const node = (tag: string, text: string): HTMLElement => {
//   const node = document.createElement(tag);
//   node.appendChild(document.createTextNode(text));
//   return node;
// };

const result = (success: boolean): void =>
  document.querySelector("#vcard")?.classList.add(success ? "success" : "failure");

const downloadBlob = (blob: Blob, name: string): void => {
  const link = document.createElement("a");
  link.href = window.URL.createObjectURL(blob);
  link.download = name;
  link.click();
};

(async () => Promise.resolve()
  .then(_ => {
    const data = location.hash.substring(1);
    location.hash = "";
    return data;
  })
  .then(d => d.length > 0 ? d : Promise.reject("no data provided"))
  .then(decodeURIComponent)
  .then(atob)
  .then(str2ab)
  .then(gunzip)
  .then(v => typeof v !== "undefined" ? v : Promise.reject("empty hash data"))
  .then(v => v.length > 0 ? v : Promise.reject("empty hash data"))
  .then(v => v.startsWith("BEGIN:VCARD") ? v : Promise.reject("invalid vCard data"))
  .then(str2ab)
  .then(v => new Blob([v], { type: "text/vcard" }))
  .then(b => {
    result(true);
    downloadBlob(b, "artero.vcf");
  })
  .catch(reason => {
    console.info(reason);
    result(false);
    return;
  }))();
