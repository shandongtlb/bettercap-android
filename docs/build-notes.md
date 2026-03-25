# Build notes

This repository stores the Android ARM64 bettercap packaging assets and a short trace of the toolchain inputs.

## Important local source paths used during reconstruction

- bettercap source: `/root/research/bettercap`
- android libusb source/assets: `/root/research/android-libusb`
- output package: `/root/build-out/bettercap-android-arm64-v2415-package`

## Minimal package contents

- `bettercap`
- `libusb1.0.so`
- `libusb-1.0.so`
- `README.txt`

## Packaging

Use:

```bash
./scripts/package-bettercap-android.sh
```

This rebuilds a clean packaging directory from the checked-in packaging assets.
