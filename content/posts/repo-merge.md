+++
title = "软件包仓库合并"
author = "yetist"
description = "软件包仓库合并"
date = 2023-06-13T20:13:29+08:00
draft = false
+++

Loong Arch Linux 仓库合并工作已完成。

## GIT 仓库

Loong Arch Linux 完成了[community](https://github.com/loongarchlinux/community) 仓库内容到 [extra](https://github.com/loongarchlinux/extra) 仓库合并，同时 [community](https://github.com/loongarchlinux/community) 仓库已经归档，不再接受 PR。

## 二进制软件包仓库

合并：[community] 仓库合并到 [extra]。

新增：[core-testing]、[extra-testing] 及 [laur-testing] 仓库。

弃用：[testing]、[community] 及 [community-testing] 仓库。

本次系统更新完成之后，请合并 pacman 的 `/etc/pacman.conf.pacnew` 配置文件。

如果没有 `/etc/pacman.conf.pacnew` 文件，请确认您的 pacman 版本 >= 6.0.2-8：

```
$ pacman -Syu "pacman>=6.0.2-8"
```

同时已经完成了以下清理工作：清理旧iso、2022.03 仓库，旧虚拟机镜像。
