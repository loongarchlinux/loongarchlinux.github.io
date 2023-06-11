+++
title = "Loong Arch Linux 下载"
date = 2023-06-10T11:57:24+08:00
draft = false
+++

## 发布信息

安装镜像可以刻录到光盘、挂载为 ISO 文件或直接写入优盘。它仅适用于从新安装；现有的 Arch Linux 系统始终可以用 `pacman -Syu` 进行更新。

- 最新版本：2023.05.08
- 内核版本：6.3.0
- ISO 大小：590 MB
- [安装指南](/pages/install/)

## 现有用户

如果您已经是 Loong Arch Linux 用户，则无需下载新的 ISO 来更新现有系统。您可能正在寻找更新的[镜像列表](/pages/mirrorlist/)。

## 网络引导

如果您有有线连接，则可以直接通过网络启动最新版本。

- [Loong Arch Linux 网络引导](/pages/netboot/)

## Docker 镜像

Github Packages 提供了官方 Docker 镜像，您可以使用以下命令获取：

```
docker pull ghcr.io/loongarchlinux/archlinux:latest
```

## 虚拟机镜像

您可从以下镜像仓库下载 QEMU qcow2 镜像，您可能希望了解如何从[虚拟机运行](/pages/vmrun/)系统。

- [https://mirrors.wsyu.edu.cn](https://mirrors.wsyu.edu.cn/loongarch/archlinux/iimages/)
- [https://mirrors.pku.edu.cn](https://mirrors.pku.edu.cn/loongarch/archlinux/images/)
- [https://mirrors.nju.edu.cn](https://mirrors.nju.edu.cn/loongarch/archlinux/images/)
- [https://mirrors.iscas.ac.cn](https://mirrors.iscas.ac.cn/loongarch/archlinux/images/)


## HTTP 直接下载

您可从以下镜像仓库下载iso：

- [https://mirrors.wsyu.edu.cn](https://mirrors.wsyu.edu.cn/loongarch/archlinux/iso/latest/)
- [https://mirrors.pku.edu.cn](https://mirrors.pku.edu.cn/loongarch/archlinux/iso/latest/)
- [https://mirrors.nju.edu.cn](https://mirrors.nju.edu.cn/loongarch/archlinux/iso/latest/)
- [https://mirrors.iscas.ac.cn](https://mirrors.iscas.ac.cn/loongarch/archlinux/iso/latest/)

如果您为 Loong Arch Linux 提供了仓库镜像，请联系我们将其添加到这个列表中。
