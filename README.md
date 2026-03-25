# bettercap Android ARM64 简短编译/整理说明

## 目标

整理 Android ARM64 的 bettercap 交付物与相关工具文件，形成可追溯仓库。

## 当前交付结构

```text
packaging/bettercap/
  README.txt
  bettercap
  libusb1.0.so
  libusb-1.0.so
```

说明：
- `bettercap` 为 Android ARM64 主程序
- `libusb1.0.so` 为 Android ARM64 运行时库
- `libusb-1.0.so` 为兼容性软链接名
- `libpcap` 已链接进主程序

## 关键来源

### bettercap 源码
```text
/root/research/bettercap
```

### Android libusb 相关来源
```text
/root/research/android-libusb
```

关键参考文件：
```text
/root/research/android-libusb/libusb/android/libs/arm64-v8a/libusb1.0.so
```

### 最终打包产物来源
```text
/root/build-out/bettercap-android-arm64-v2415-package
```

## 当前仓库保存的内容

- `packaging/bettercap/`：当前已整理好的发布文件
- `refs/`：参考性二进制/来源文件
- `logs/`：构建/打包日志快照
- `scripts/package-bettercap-android.sh`：二次打包脚本

## 为什么单独建仓库

因为这里更像“Android bettercap 交付链路仓库”，不是上游 bettercap 主源码仓库。

建议：
- 上游源码继续看 `bettercap` 官方仓库或本地 research 副本
- 这个仓库用于保留 Android ARM64 交付物、包装方式、日志和说明

## 已知情况

- 交付物是 Android ARM64 可执行文件 + 外置 libusb
- 当前包内不依赖额外 Python runtime
- `libusb-1.0.so` 在目标环境若不保留软链接，可复制 `libusb1.0.so` 生成同名文件替代
