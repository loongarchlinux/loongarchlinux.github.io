---
title: "关于 Loong Arch Linux"
date: 2021-06-19T19:12:53+08:00
---

Loong Arch Linux 是一个针对 LoongArch 平台移植的 Arch Linux 发行版，遵循轻量、简洁、优雅的开发原则，借灵活的架构应用于各种环境。Arch 安装后只提供最基本的系统，用户可以根据自己的需求来搭建不同的系统环境。官方并不提供图形化的配置工具，多数系统配置是通过修改文本文件来进行的，并尽力提供最新稳定版本的软件。

Loong Arch Linux 使用 Pacman 作为包管理器，它在提供了一个简单的包管理器同时，也提供了一个易用的包构建系统，使用户能够轻松地管理和定制官方提供的、用户自己制作的、甚至是来自第三方的各种软件包。仓库系统也能够让用户轻松的构建和维护自己的编译脚本、软件包和仓库，这样将有助于社区的成长和建设。

Loong Arch Linux 的基本安装包由 [core] 软件库提供。此外 [extra] 和 [laur] 软件库则提供了大量的的高品质软件以满足你的需求。Loong Arch Linux 同时也支持 Arch 用户软件仓库(AUR)提供的 [unsupported] 软件库，里面有大量的编译脚本，用户可以通过 `makepkg` 工具轻松地从源码中编译软件。

Loong Arch Linux 采用“滚动升级”策略，这样可以实现“一次安装，永久更新”。升级到下一个“版本”的 Loong Arch Linux 几乎不需要重新安装系统，只需一行命令，你就能轻松的享受到最新的 Loong Arch Linux。

Loong Arch Linux 努力和上游软件源码保持一致，只有使程序能够在 Loong Arch Linux 正常编译运行的补丁才会被加入更新中。

总之，Loong Arch Linux 是一个灵活、简洁的、满足有一定经验的 Linux® 用户的需求的发行版。它强大且易于管理的特性，使其成为可以完美胜任服务器和工作站的发行版。它可以变成任何你想要的样子。如果你也认为这是一个 GNU/Linux 发行版该做的，欢迎你来自由使用并参与其中，为社区做出贡献，欢迎来到 Loong Arch Linux！

## 为什么叫 Loong Arch Linux ?

Loong Arch Linux 表示 LoongArch ArchLinux，将重复的 Arch 部分去掉，则为 Loong ArchLinux 或 LoongArch Linux, 给 Arch 前后都加上空格，则为 Loong Arch Linux。

没有考虑叫 ArchLinux LoongArch？嗯，确实考虑过。由于 Archlinux 的 Arm 版本叫 ArchLinuxARM, Riscv 版本叫 ArchRiscv, 按这种命名规则，发行版的中文名叫 **阿龙(ArchLoong)**，也许可以设计个吉祥物？😁

## Loong Arch Linux 历史

- 2021年06月19日，发布 [Alpha](http://archlinux.oukan.online/alpha/index.html) 版本，提供 bootstrap 镜像，带 [core]、[extra]、[community] 和 [aur] 软件包仓库。
- 2022年05月13日，发布 [2022.03](https://bbs.loongarch.org/d/67-loongarchlinux-202203) 版本，提供 iso 镜像，支持 `xfce`、`mate` 桌面环境，带 [core]、[extra]、[community] 及 [aur] 软件包仓库。
- 2022年09月29日，发布 [2022.09](https://bbs.loongarch.org/d/126-archlinux-loong64-202209) 版本, 提供 iso 镜像，增加 `lxde`、`lxqt` 等桌面环境，带 [core]、[extra]、[community] 及 [aur] 软件包仓库。发行版架构名称采用 loong64。
- 2022年10月21日，首次启用[testing仓库](https://bbs.loongarch.org/d/126-archlinux-loong64-202209/40)，从现在开始，LoongArch 上游社区 ABI 趋于稳定，开始采用滚动升级模式，用户不再需要使用 iso 重新安装系统。
- 2022年10月24日，发布新 [iso](https://bbs.loongarch.org/d/126-archlinux-loong64-202209/66) ，其中集成了 `archinstall` TUI 安装程序，并提供带中文界面的安装程序 `setup`，降低了新手用户安装系统的难度。
- 2022年10月25日，[debuginfod 服务](https://bbs.loongarch.org/d/126-archlinux-loong64-202209/76) 正常工作。
- 2022年11月04日，仓库新增 [cinnamon](https://bbs.loongarch.org/d/126-archlinux-loong64-202209/90) 桌面环境。
- 2022年12月03日，发布[虚拟机安装镜像](https://bbs.loongarch.org/d/126-archlinux-loong64-202209/120)。
- 2022年12月28日，经过一个季度滚动升级测试，上游 ABI 已经完全稳定。移除仓库地址中的 `2022.09` 版本号路径，将 [aur] 仓库改名为 [laur]，清理 python2，新增 `lat` 二进制翻译测试功能，openssl 升级到 3.0等。见[这里](https://bbs.loongarch.org/d/126-archlinux-loong64-202209/128)
- 2023年01月28日，社区开发者木棉为 Loong Arch Linux 制作了采用 `Calamares` 安装程序的图形化 [LiveCD](https://bbs.loongarch.org/d/176-calamareslivecdarchlinux)
- 2023年02月24日，为 Loong Arch Linux 增加[网络启动](https://bbs.loongarch.org/d/179-archlinux)支持，采用 ipxe 可从互联网启动并安装系统。
- 2023年03月15日，正式删除 `2022.09` 仓库路径，正式删除 [aur] 软件包仓库。
- 2023年03月26日，为 Loong Arch Linux 设计了自己的 [logo](https://avatars.githubusercontent.com/u/84459977)。
- 2023年04月05日，新增 `Gnome` 和 `KDE` 桌面环境。
- 2023年04月08日，北大镜像站为 Loong Arch Linux 提供仓库镜像。
- 2023年04月10日，新增 Loong Arch Linux 的`Docker` 镜像。
- 2023年06月03日，开始合并 [community] 到 [extra] 仓库，开始清理旧的iso、旧的2022.03仓库、旧的qemu qcow2镜像等。
- 2023年06月07日，完成旧iso、2022.03仓库、qemu qcow2镜像的清理，完成 git 仓库 [community] 到 [extra] 合并， [community] git 仓库已归档，优化开发流程，git 仓库可接受 PR 贡献。
- 2023年06月14日，从20230614版本开始，新增磁力链iso下载方式。
- 2023年06月26日，[网站](https://loongarchlinux.org) 完工，正式上线。
