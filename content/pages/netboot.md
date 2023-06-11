+++
title = "网络启动"
author = ""
date = 2023-06-11T16:19:48+08:00
draft = false
+++

## 关于网络启动

![](/pages/images/ipxe.png)

网络引导(Netboot)镜像是个非常小的文件（大约1~2M），可用于在系统启动时从网络下载最新的 Arch Linux 版本。 网络引导(Netboot)镜像本身无需更新，最新发布版本将自动可用。

网络引导(Netboot) 是使用自定义 iPXE 构建的。 live 系统的 Linux 内核、initramfs 和 squashfs 文件是从 Arch Linux 镜像下载的。

## 要求

要使用 netboot，必须满足以下要求：

- 具有 DHCP 自动配置的有线（以太网）互联网连接
- 足够的内存来存储和运行系统

## 下载

- [ipxe-arch.efi](https://mirrors.wsyu.edu.cn/loongarch/archlinux/netboot/ipxe-arch.efi) - loong64 UEFI 网络引导镜像
- [ipxe-arch.iso](https://mirrors.wsyu.edu.cn/loongarch/archlinux/netboot/ipxe-arch.iso) - loong64 UEFI iPXE iso

## 用法说明

ipxe.efi 镜像可用于在 UEFI 模式下启动 Arch Linux 网络引导。

ipxe.efi 镜像可以通过 efibootmgr 添加为启动选项，从 systemd-boot、grub 或 rEFInd 等启动管理器启动，或者直接从 UEFI shell 启动。

下载文件： ipxe-arch.efi，并保存到 EFI 系统分区 (ESP) 的 `/EFI/arch_netboot` 目录中。

假设您的 (ESP) 分区挂载到了 /boot/efi 目录，则参考以下命令：

```
# mkdir /boot/efi/EFI/arch_netboot
# sudo cp ipxe-arch.efi /boot/efi/EFI/arch_netboot/arch_netboot.efi
```

### 1. 添加为固件启动选项

首先安装 efibootmgr 包，然后下载 UEFI 网络引导镜像。

假设您的 EFI 系统分区 (ESP) 是 /dev/sda1，在操作系统下使用 efibootmgr 命令，为UEFI 增加一个新的菜单项：

```
# efibootmgr --create --disk /dev/sda --part 1 --loader /EFI/arch_netboot/arch_netboot.efi --label "Arch Linux Netboot" --unicode
```

之后开机时，通过按 F12 调出快捷菜单，并选择Arch Linux Netboot 来启动网络安装。

### 2. 通过Grub 启动管理器启动

假设您的<ESP> 分区的UUID为1234-5678，则创建 /boot/grub/custom.cfg 文件，内容参考以下内容：

```
menuentry 'Arch Linux Netboot' {
        insmod part_gpt
        insmod fat
        if [ x$feature_platform_search_hint = xy ]; then
          search --no-floppy --fs-uuid --set=root  1234-5678
        else
          search --no-floppy --fs-uuid --set=root 1234-5678
        fi
        chainloader /EFI/arch_netboot/arch_netboot.efi
}
```

开机进入 Grub 菜单时，将多出一项 Arch Linux Netboot 启动菜单，选择之后，将进入网络安装环境。

### 3. 从UEFI Shell 手动启动

重启按F12，在界面中选择 UEFI Shell，进入UEFI Shell 环境，输入以下命令启动：

```
Shell> fs0:
Shell> EFI\arch_netboot\arch_netboot.efi
```

