+++
title = "Loong Arch Linux 安装指南"
author = ""
date = 2023-06-11T17:44:03+08:00
draft = false
+++

## TUI安装（推荐）

自2023.07 版本之后，会自动进入安装界面。

如果您使用早期版本，可运行以下命令来进入安装界面：

```
# archinstall
```

对于中文用户来说，可参考以下配置建议：

- **镜像区域** 请设置为 `China`, 以提高软件包下载速度
- **本地语言** 请设置为 `zh_CN`，以自动安装中文字体和输入法
- **磁盘布局** 中的文件系统格式，建议选择为 `ext4`（测试中发现内核的 `xfs` 文件系统有bug)
- **引导加载程序** 建议设置为 `grub-install`
- **交换分区** 请根据实际情况决定是否启用
- **配置文件** 可以选择桌面(desktop)、最小化(minimal)、服务器(server)及图形(xorg)环境安装，桌面推荐 `xfce4` 或 `mate`，显卡驱动请选择应选择 `AMD / ATI (open-source)`
- **音频** 推荐设置为 `pipewire`
- **内核** 请设置为 `linux`
- **网络配置** 推荐设置为 `NetworkManager`
- **时区** 推荐设置为 `Asia/Shanghai`
- **自动时间同步 (NTP)**  建议启用

设置完成之后，界面大致会显示如下：

![](/pages/images/archinstall.png)

## 手工安装

参考文档：
- [Archlinux Wiki 安装指南](https://wiki.archlinuxcn.org/wiki/%E5%AE%89%E8%A3%85%E6%8C%87%E5%8D%97)
- [LA UOSC 安装指南](https://bbs.loongarch.org/d/88-archlinux)

### 1. 分区

使用 cfdisk 对磁盘分区，假定要被分区的磁盘为 _/dev/the_disk_to_be_partitioned_

```
# cfdisk /dev/the_disk_to_be_partitioned  #要被分区的磁盘

```
分区表类型：**gpt**

分区最基本要求：
- 一个 EFI 分区， 类型为 "EFI 系统" ，大小不小于200M，建议500M，假定为 _/dev/efi_system_partition_
- 一个根分区，类型为 “Linux root (LoongArch-64)" ，假定为 _/dev/root_partition_

### 2. 格式化

```
# mkfs.ext4 /dev/root_partition            #格式化根分区
# mkfs.fat -F 32 /dev/efi_system_partition #格式化EFI系统分区
```

### 3. 挂载分区

将根磁盘卷挂载到 /mnt，例如：

```
# mount /dev/root_partition /mnt
```

挂载 EFI 系统分区：

```
# mount --mkdir /dev/efi_system_partition /mnt/boot/efi
```

### 4. 安装软件包

```
# pacstrap /mnt base base-devel linux linux-firmware networkmanager grub efibootmgr
```

### 5. 生成新系统的 Fstab

用以下命令生成 fstab 文件 (用 -U 或 -L 选项设置 UUID 或卷标)：

```
# genfstab -U /mnt > /mnt/etc/fstab
```

### 6. Chroot

chroot 到新安装的系统：

```
# arch-chroot /mnt
```

### 7. 配置系统

```
# ## 时区
# ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# hwclock --systohc

# ## 本地化
# sed -i -e 's/#zh_CN.UTF-8/zh_CN.UTF-8/' -e 's/#en_US.UTF-8/en_US.UTF-8/' /etc/locale.gen
# locale-gen
# echo LANG=zh_CN.UTF-8 > /etc/locale.conf

# ## 主机名（如myhostname）
# echo myhostname > /etc/hostname
# echo "127.0.0.1 localhost" >> /etc/hosts
# echo "::1       localhost" >> /etc/hosts
# echo "127.0.1.1 myhostname.localdomain myhostname" >> /etc/hosts

# ## initramfs
# mkinitcpio -P

# ## 创建帐户（如user)
# useradd -m -G wheel user
# passwd user

# ## 设置root密码
# passwd

# ## 网络服务
# systemctl enable NetworkManager.service

# ## 引导器
# grub-install
# grub-mkconfig -o /boot/grub/grub.cfg
```

### 8. 桌面环境

```
# ## 安装基本图形系统 xorg
# pacman -Sy --needed xorg xorg-server xf86-video-amdgpu xf86-video-ati xf86-video-loongson

# ## 安装桌面环境(xfce4/mate/gnome/kde)
# pacman -Sy --needed xfce4
# pacman -Sy --needed mate
# pacman -Sy --needed gnome
# pacman -Sy --needed plasma-meta

# ## 安装显示管理器(lightdm/gdm/sddm)
# pacman -Sy --needed lightdm lightdm-gtk-greeter
# pacman -Sy --needed gdm
# pacman -Sy --needed sddm
# systemctl enable <DM> # DM 为 lightdm、gdm、sddm 三选一
```

### 9. 字体及输入法

```
# ## 安装中文字体
# pacman -Sy wqy-microhei wqy-microhei-lite wqy-zenhei
# fc-cache -fv

# ## 安装输入法
# pacman -Sy --needed fcitx5 fcitx5-chinese-addons fcitx5-configtool fcitx5-gtk fcitx5-qt
# cat > /etc/X11/xinit/xinitrc.d/50-input.sh << EOF
export XIM=fcitx
export GTK_IM_MODULE=fcitx
export QT_IM_MODULE=fcitx
export XMODIFIERS="@im=fcitx"
EOF
# chmod +x /etc/X11/xinit/xinitrc.d/50-input.sh
```

## 常见问题

1. 签名无效

进入archlinux 安装环境之后，运行 `pacman -Sy` 命令，遇到如下错误：
```
错误：core: 来自 "65D4986C7904C6DBF2C4DD9A4E4E02B70BA5C468" 的签名是未知信任的
错误：extra: 密钥 "65D4986C7904C6DBF2C4DD9A4E4E02B70BA5C468" 未知
错误：密钥环不可写
错误：未能同步所有数据库（无效或已损坏的数据库（PGP 签名))
```

从iso启动之后，会自动运行 `pacman-key` 来初始化签名数据库，您可以重启系统，并等待后台 `pacman-key` 进程结束即可。

或者，在不重启系统的情况下，分别运行以下命令，重建本地签名数据库即可恢复正常：

```
pacman-key --init
pacman-key --populate archlinux
```

2. 如何关闭 `archinstall` 测试镜像连接？

运行 `archinstall` 时，会长时间显示如下内容：
```
Testing connectivity to the Arch Linux mirrors ...
```

请耐心等待其结束，或者使用 `archinstall --skip-mirror-check` 命令重新运行。
