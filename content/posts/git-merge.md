+++
title = "git 仓库合并"
author = "yetist"
description = "git 仓库合并"
date = 2023-06-03T19:32:06+08:00
draft = false
+++

ArchLinux(x86_64) 于2023年05月19日开始，将源码仓库从 SVN 迁移到了 GIT 上，并且完成了仓库拆分与合并。

拆分：[testing] 仓库将拆分为 [core-testing] 和 [extra-testing] ， [staging] 仓库将拆分为 [core-staging] 和 [extra-staging] 。

合并：[community] 仓库将合并为 [extra] ，因此在迁移后将为空。

弃用：[community-testing]、[community-staging] 仓库将被弃用。

Loong Arch Linux 跟随x86_64，完成了 GIT 仓库的合并，[community](https://github.com/loongarchlinux/community) 仓库内容合并到了 [extra](https://github.com/loongarchlinux/extra) 仓库。

同时将完成以下清理工作：

- 清理旧 [iso](https://mirrors.wsyu.edu.cn/loongarch/archlinux/iso/) 文件，未来将仅保留最近 4 个版本。
- 清理旧的 [2022.03](https://mirrors.wsyu.edu.cn/loongarch/2022.03/) 仓库，这个版本使用的是 ABI 1.0，和现有系统并不兼容，且几乎无人使用。
- 清理旧的[虚拟机镜像](https://mirrors.wsyu.edu.cn/loongarch/archlinux/images/)，未来将仅保留最新的镜像。
