Package: bettercap v2.41.5 for Android ARM64

Files:
- bettercap              -> stripped Android ARM64 binary
- libusb1.0.so           -> Android ARM64 libusb runtime library
- libusb-1.0.so          -> symlink name for runtime linker compatibility

Notes:
- This package was cross-compiled for Android ARM64.
- libpcap is already linked into the bettercap binary.
- libusb is provided as an external shared library in this package.
- If the target environment does not preserve symlinks, copy libusb1.0.so
  and create/duplicate it as libusb-1.0.so on device.
