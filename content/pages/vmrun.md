+++
title = "Loong Arch Linux 虚拟机运行"
date = 2023-06-10T11:57:24+08:00
draft = false
+++

## 虚拟机运行

Loong Arch Linux 发布了用于 QEMU 的虚拟机镜像文件，用户可选择合适的[镜像](https://mirrors.pku.edu.cn/loongarch/archlinux/images/)通过虚拟机来运行和体验。

- archlinux-minimal-YYYY.MM.DD-loong64.qcow2.zst: 可引导到命令行界面
- archlinux-xfce4-YYYY.MM.DD-loong64.qcow2.zst: 可引导到 XFCE 桌面环境
- archlinux-mate-YYYY.MM.DD-loong64.qcow2.zst: 可引导到 Mate 桌面环境

注：为节省网络流量，镜像发布为 zst 压缩格式，下载之后需要解压出 qcow2 文件。

## 运行要求

- QEMU 需要使用 7.2 及以上版本
- 需要准备虚拟机固件

## 运行方法

请参考以下命令运行：

```
qemu-system-loongarch64 \
    -m 4G \
    -cpu la464-loongarch-cpu \
    -machine virt \
    -smp 4 \
    -bios ./QEMU_EFI_x.y.fd \
    -serial stdio \
    -device virtio-gpu-pci \
    -net nic -net user \
    -device nec-usb-xhci,id=xhci,addr=0x1b \
    -device usb-tablet,id=tablet,bus=xhci.0,port=1 \
    -device usb-kbd,id=keyboard,bus=xhci.0,port=2 \
    -hda archlinux-xxx-yyyy-loong64.qcow2
```

注：

- `QEMU_EFI_x.y.fd`: 虚拟机固件文件，请选择 QEMU 的对应版本下载，[QEMU_EFI_7.2.fd](https://mirrors.pku.edu.cn/loongarch/archlinux/images/QEMU_EFI_7.2.fd)、[QEMU_EFI_8.0.fd](https://mirrors.pku.edu.cn/loongarch/archlinux/images/QEMU_EFI_8.0.fd)。或者从 [edk2-loongarch64](/package/?repo=extra&arch=any&name=edk2-loongarch64) 软件包中解压适用于此[QEMU](/package/?repo=extra&arch=loong64&name=qemu-system-loongarch64)版本的固件文件。


## 虚拟机中的帐号

管理员名称：root

管理员密码：loongarch

用户名称：loongarch

用户密码：loongarch
