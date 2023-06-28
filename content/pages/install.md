+++
title = "Loong Arch Linux 安装指南"
author = ""
date = 2023-06-11T17:44:03+08:00
draft = false
+++

## TUI安装（推荐）

运行命令：

```
# archinstall
```

设置中文界面：选择 "Archinstall language", 选中 "Simplified Chinese" 回车，界面将以中文显示，后续根据界面进行操作。

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
