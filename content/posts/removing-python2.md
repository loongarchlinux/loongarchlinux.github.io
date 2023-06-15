+++
title = "仓库变更及python2清理"
author = "yetist"
description = "从仓库中删除 python2"
date = 2022-12-28T10:15:08+08:00
draft = false
+++

## 仓库变更

由于原来的 [2022.09](https://mirrors.wsyu.edu.cn/loongarch/2022.09/) 目录，已经无法代表当前的 Archlinux 状态，故将其修改为 [archlinux](https://mirrors.wsyu.edu.cn/loongarch/archlinux/)。

由于原来的二进制仓库 [aur] 与 AUR 用户仓库同名，在互相交流时会带来歧义，故将原 [aur] 二进制仓库改名为 [laur]，此仓库会提供部分 LoongArch 架构相关的典型应用，或 AUR 上重要的二进制软件包。

不再提供 [wine-apps] 仓库，如果有需要测试 wine 应用的同学，请安装 `pamac-aur` 软件包，并在其中搜索 `wine-apps` 获取。

## python2 清理

Loong Arch Linux 从本次升级之后，已经完全从仓库中删除了所有依赖 python2 的包。如果您的系统上仍然安装了 python2 ，请考虑删除它及任何依赖于 python2 的包。

如果您仍然需要 python2 软件包，您可以保留它，但请注意不会有安全更新。如果您需要修补的软件包，请查阅 AUR，或使用[非官方的用户仓库](https://wiki.archlinux.org/title/Unofficial_user_repositories)。

## 二进制翻译

本次升级将带来二进制翻译功能，供测试使用，lat软件包位于 [laur] 仓库。如果使用 `pamac-aur` 安装 wine-apps 软件，lat 会被自动安装。

## 升级风险提示（重要）

本次升级，将 openssl 从 1.1 版本升级到了 3.0 版本，影响到大量关键软件包的重构，这些软件包会影响系统的正常运行，因此， **本次升级，必须要使用 `pacman -Syu` 进行整体更新** ，不可以选择性更新软件包，否则会导致系统无法正常使用。

另外，如果更新系统后，仍有个别软件还需要依赖 openssl 1.1 版本，可单独安装 `openssl-1.1` 这个包，这个包将提供 1.1 版本的运行时库。
