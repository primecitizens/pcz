# win32 platform codegen

## Maintenance Notes

To download Windows SDKs, go to `https://developer.microsoft.com/en-us/windows/downloads/sdk-archive/` and find

- Windows SDK for Windows 11 (10.0.22621.1778)
- Windows 10 SDK version 2104 (10.0.20348.0)
- Windows SDK for Windows 7 and .NET Framework 4

DO NOT USE Windows 11 for development, release note of Windows 10 SDK version 2104 says:

> Clang/LLVM for Windows v11 targeting ARM64 is not compatible with the latest winnt.h
> As a workaround, use the previous version of the Windows 10 SDK (build 19041), or clang/LLVM for Windows v10 when targeting ARM64 platforms

## References

- [microsoft/win32metadata](https://github.com/microsoft/win32metadata)
