(async () =>
  Promise.resolve(document.querySelectorAll<HTMLLinkElement>("link[rel=preload][as=style]"))
    .then(links => links.forEach(link => {
      link.rel = "stylesheet";
      link.removeAttribute("as");
    })))();
