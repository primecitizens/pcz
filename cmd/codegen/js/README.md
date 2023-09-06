# js platform codegen

Upstream Web APIs:

- [w3c/webref@curated ed/idl](https://github.com/w3c/webref)
- [mdn/browser-compat-data](https://github.com/mdn/browser-compat-data)

## Maintenance Notes

The BCD (Browser Compat Data) can be downloaded from the release area of the official MDN repository (`https://github.com/mdn/browser-compat-data/releases`, find `data.json`, about 12MB in size) or download from CDN sites like `https://unpkg.com/@mdn/browser-compat-data/data.json`

To update WebIDL files, you will need to have firefox source code avaiable (more specifically, its `dom/webidl` and `dom/chrome-webidl` subdirs)

```bash
# to download (from their git mirror)
git clone --depth=1 -b release https://github.com/mozilla/gecko-dev
```
