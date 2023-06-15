+++
title = "Debuginfod 服务"
author = "yetist"
description = "社区提供 debuginfod 服务"
date = 2022-10-25T18:51:47+08:00
draft = false
+++

Loong Arch Linux 现在提供调试符号（debug）包和 debuginfod 服务。

debuginfod 服务将为 Loong Arch Linux 上的软件提供调试符号信息和源码列表，这些可以被调试器比如 gdb 和 delve 利用。

系统中已经默认设置好 `DEBUGINFOD_URLS` 环境变量，无需自行设置，您需要做的就是直接运行调试器，比如：

```
$ gdb curl
(gdb) b main
(gdb) l
(gdb) r
(gdb) l
(gdb)
```
更多信息请参阅 [Debuginfod](https://wiki.archlinux.org/title/Debuginfod) 维基页面。
