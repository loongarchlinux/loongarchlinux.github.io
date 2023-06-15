+++
title = "仓库路径变更及aur仓库改名"
author = "yetist"
description = "仓库路径变更及aur仓库改名"
date = 2023-03-15T21:42:18+08:00
draft = false
+++

经过一段时间的过渡，今天正式完成仓库变更。

1. 删除 `2022.09` 仓库路径

请修改 `/etc/pacman.d/mirrorlist` 文件中的仓库地址，将以下内容

```
Server = https://mirrors.wsyu.edu.cn/loongarch/2022.09/$repo/os/$arch
```

修改为：

```
Server = https://mirrors.wsyu.edu.cn/loongarch/archlinux/$repo/os/$arch
```

2. 删除 [aur] 软件包仓库

由于 [aur] 仓库已经改名为 [laur]，请将 `/etc/pacman.conf` 文件中的以下内容

```
[aur]
SigLevel = Optional TrustAll
Include = /etc/pacman.d/mirrorlist
```

修改为：

```
[laur]
SigLevel = Optional TrustAll
Include = /etc/pacman.d/mirrorlist
```
