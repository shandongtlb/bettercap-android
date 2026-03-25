#!/usr/bin/env bash
set -euo pipefail
ROOT="${ROOT:-$(pwd)}"
OUTDIR="${OUTDIR:-$ROOT/dist/bettercap-android-arm64-package}"
rm -rf "$OUTDIR"
mkdir -p "$OUTDIR"
cp packaging/bettercap/README.txt "$OUTDIR/"
cp packaging/bettercap/bettercap "$OUTDIR/"
cp packaging/bettercap/libusb1.0.so "$OUTDIR/"
cp -a packaging/bettercap/libusb-1.0.so "$OUTDIR/"
( cd "$(dirname "$OUTDIR")" && tar -czf bettercap-android-arm64-package.tar.gz "$(basename "$OUTDIR")" )
printf "[ok] package dir: %s\n" "$OUTDIR"
ls -lah "$OUTDIR"
