# darwin platform api codegen

## Maintenance Notes

To generate darwin platform apis, you need to have both `Xcode.app` (more specifically, the SDKs comes with it) and `llvm` installed

To download Xcode, go to (for example) `https://download.developer.apple.com/Developer_Tools/Xcode_14.3.1/Xcode_14.3.1.xip`

**HINT**: Since we only need to have SDK header files in XCode.app, if you don't want to create/login your apple account in order to download the Xcode (which is just another absurd apple thing), you may find images built by cirruslabs for Tart very helpful (for example `https://github.com/cirruslabs/macos-image-templates/pkgs/container/macos-ventura-xcode`)

## Paths

- XCode.app
  - Installed from App Store or `.xip`: `/Applications/Xcode.app`
- llvm root
  - Installed via `brew`: `$(brew --prefix)/Cellar/llvm/16.0.6` (change `16.0.6` to the version you installed)
