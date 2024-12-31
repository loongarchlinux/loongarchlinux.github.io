+++
title = "新年新起点，迈入社区维护新阶段"
author = "yetist"
description = "维护团队变更公告"
date = 2024-12-01T00:00:00+08:00
draft = false
+++

自 2024 年以来，龙架构开源软件生态持续成熟，仓库中有越来越多的软件包可用，伴随而来的问题是，每次更新系统，要下载、编译的包数量庞大，同时个人工作繁忙，无法保证维护投入（实际已远低于 0.5 人），所以 2024 年采用按季度进行更新。

为了保持滚动更新的特性，Loong Arch Linux 发行版急需更多的维护力量参与。经过和社区多次沟通、对接，北京大学学生 Linux 俱乐部（LCPU）接手了这一项目的维护工作。

LCPU 接手项目后，采用了全新的维护结构设计。在 2024 年 8 月，LCPU 的维护团队借鉴 [archriscv 社区](https://github.com/felixonmars/archriscv-packages) 经验，从头构建了更易于维护的[补丁集式维护仓库](https://github.com/lcpu-club/loongarch-packages)与更强大易用的 [devtools-loong64](https://aur.archlinux.org/packages/devtools-loong64) 开发者工具。经过两个季度的试运行，目前维护工作已经步入正轨。LCPU 的维护团队修复并构建了 KDE 6、Firefox、Chromium、Code - OSS 等广受欢迎的日用应用，积极推动了多个重要项目的龙架构修复上游化，整理出了丰富详尽的[维护文档](https://github.com/lcpu-club/loongarch-packages/wiki)，使得社区参与广泛可行，展现了强大了开发、维护、教育与组织能力。

新的社区团队的目标是：

* 及时跟随 Arch Linux 官方的更新进度，持续维护 Loong Arch Linux 发行版
  * 最终借助 [Arch Linux Ports](https://rfc.archlinux.page/0032-arch-linux-ports/) 平台，推动 Arch Linux 官方增加龙架构支持，同步发行龙架构版本
* 修复上游软件在龙架构上的构建问题，并尽可能将修复上游化
  * 最终推动各软件包上游提高龙架构的维护等级
* 培养更多能够为龙架构生态作出贡献的人才，建设更加开放健康的开源社区

欢迎有维护意愿的同学[加入社区](https://github.com/lcpu-club/loongarch-packages)团队（可联系社区负责人 [wszqkzqk](mailto:wszqkzqk@qq.com)），共同建设龙架构的软件生态，共同维护 Loong Arch Linux 发行版。

目前社区团队维护的系统整体上比当前系统的软件包版本新，用户可通过修改仓库配置，平滑进行切换。

**仓库切换方法**

1. 编辑文件 `/etc/pacman.d/mirrorlist` ，将内容修改为：

```
Server = https://mirrors.pku.edu.cn/loongarch-lcpu/archlinux/$repo/os/$arch
Server = https://loongarchlinux.lcpu.dev/loongarch/archlinux/$repo/os/$arch
```

2. 更新系统，运行命令：

```bash
sudo rm -f /var/lib/pacman/sync/*
sudo pacman -Syuu
```

3. LCPU 维护的新 Loong Arch Linux 项目建议使用 `pacman-mirrorlist-loong64` 作为镜像列表文件以避免和 Arch Linux 上游的冲突，可在**换源与更新完成后**编辑 `/etc/pacman.conf`，将所有的 `Include = /etc/pacman.d/mirrorlist` 替换为 `Include = /etc/pacman.d/mirrorlist-loong64`：

```bash
sudo sed -i 's/^\s*Include\s*=\s*\/etc\/pacman\.d\/mirrorlist\s*$/Include = /etc/pacman.d/mirrorlist-loong64/' /etc/pacman.conf
```

后续使用中可按需编辑 `/etc/pacman.d/mirrorlist-loong64` 中的镜像配置列表。
