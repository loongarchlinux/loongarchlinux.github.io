+++
title = "TeX Live 打包改变了组织结构"
author = ""
date = 2023-06-30T14:12:29+08:00
draft = false
+++

从版本 2023.66594-5 开始，TeX Live 的打包改变了组织结构，更接近上游的集合(collections)。尽管新的 texlive-basic 包替换了 texlive-core 包，很多原本属于 texlive-core 包的内容（包括一些特定语种的文件）现在被拆分到了别的包中去。如果想了解哪个 Arch 软件包中提供了特定 CTAN 宏包，可以使用 tlmgr 工具，比如:

```
$ tlmgr info euler | grep collection
collection:  collection-latexrecommended
```

这个的意思是说 euler CTAN 宏包包含在了 texlive-latexrecommended 软件包中。你也可以使用 pacman -F 命令查询特定文件的归属。

我们提供了新的 texlive-meta 元包用来安装所有子包（除了特定语种的），还有新的 texlive-doc 包提供了完整的文档，用以离线查阅。
