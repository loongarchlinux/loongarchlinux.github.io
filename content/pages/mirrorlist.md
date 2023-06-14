+++
title = "Loong Arch Linux 镜像列表"
date = 2023-06-10T11:57:24+08:00
draft = false
+++

Loong Arch Linux 目前包含以下仓库镜像:

- [https://mirrors.wsyu.edu.cn](https://mirrors.wsyu.edu.cn/loongarch/archlinux/)
- [https://mirrors.pku.edu.cn](https://mirrors.pku.edu.cn/loongarch/archlinux/)
- [https://mirrors.nju.edu.cn](https://mirrors.nju.edu.cn/loongarch/archlinux/)
- [https://mirror.iscas.ac.cn](https://mirror.iscas.ac.cn/loongarch/archlinux/)

软件包 `pacman-mirrorlist` 提供了预配置好的仓库镜像，您可通过以下命令来安装/升级：

```
# pacman -Syu pacman-mirrorlist
```

## 仓库配置

请选择离您地理位置较近区域的镜像使用，通常能获得较高的下载速度。

如需调整仓库，请编辑 `/etc/pacman.d/mirrorlist` 文件，对您想使用的镜像取消注释，
并将其置于 mirrorlist 文件的最上方。

典型的仓库配置文件如下：

```
## China
Server = https://mirrors.pku.edu.cn/loongarch/archlinux/$repo/os/$arch
Server = https://mirrors.nju.edu.cn/loongarch/archlinux/$repo/os/$arch
Server = https://mirrors.wsyu.edu.cn/loongarch/archlinux/$repo/os/$arch
Server = https://mirror.iscas.ac.cn/loongarch/archlinux/$repo/os/$arch
```
