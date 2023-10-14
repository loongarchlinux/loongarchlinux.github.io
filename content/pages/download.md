+++
title = "Loong Arch Linux 下载"
date = 2023-06-10T11:57:24+08:00
draft = false
+++

## 发布信息

安装镜像可以刻录到光盘、挂载为 ISO 文件或直接写入优盘。它仅适用于从新安装；现有的 Arch Linux 系统始终可以用 `pacman -Syu` 进行更新。

- 最新版本：<div id="version" style="display:inline">2023.05.08</div>
- 内核版本：<div id="kernel" style="display:inline">6.3.0</div>
- ISO 大小：<div id="size" style="display:inline">590 MB</div>
- [安装指南](/pages/install/)

## 现有用户

如果您已经是 Loong Arch Linux 用户，则无需下载新的 ISO 来更新现有系统。您可能正在寻找更新的[镜像列表](/pages/mirrorlist/)。

## HTTP 直接下载(推荐)

您可从以下镜像仓库下载iso：

<div>
<ul id="ul_download">
</ul>
</div>

如果您为 Loong Arch Linux 提供了仓库镜像，请联系我们将其添加到这个列表中。

## BitTorrent 下载

如果您愿意，请在下载完成后保持客户端打开状态，以便您可以将其种子重新播种给其他人。

**需要支持 DHT 的客户端。** 建议使用支持 WebSeed 的客户端以获得最快的下载速度。

<ul>
    <li><img width="12" height="12" src="/images/magnet.png" alt=""/>
    磁力链接：<a id="a_magnet" href="#" title="打开磁力链接">2023.05.08 </a></li>
    <li><img width="12" height="12" src="/images/download.png" alt=""/>
    种子文件：<a id="a_torrent" href="#" title="下载种子文件">2023.05.08</a></li>
</ul>

## 网络引导

如果您有有线连接，则可以直接通过网络启动最新版本。

- [Loong Arch Linux 网络引导](/pages/netboot/)

## Docker 镜像

Github Packages 提供了官方 Docker 镜像，您可以使用以下命令获取：

```
$ docker pull ghcr.io/loongarchlinux/archlinux:latest
```

## 虚拟机镜像

您可从以下镜像仓库下载 QEMU qcow2 镜像，您可能希望了解如何从[虚拟机运行](/pages/vmrun/)系统。

- [https://mirrors.wsyu.edu.cn](https://mirrors.wsyu.edu.cn/loongarch/archlinux/images/)
- [https://mirrors.pku.edu.cn](https://mirrors.pku.edu.cn/loongarch/archlinux/images/)
- [https://mirrors.nju.edu.cn](https://mirrors.nju.edu.cn/loongarch/archlinux/images/)
- [https://mirrors.iscas.ac.cn](https://mirrors.iscas.ac.cn/loongarch/archlinux/images/)

<script src="https://cdn.bootcdn.net/ajax/libs/jquery/3.6.4/jquery.min.js"></script>
<script>
    function getfilesize(size) {
        if (!size)
            return "";

        var num = 1024.00;

        if (size < num)
            return size + "B";
        if (size < Math.pow(num, 2))
            return (size / num).toFixed(2) + "K";
        if (size < Math.pow(num, 3))
            return (size / Math.pow(num, 2)).toFixed(2) + "M";
        if (size < Math.pow(num, 4))
            return (size / Math.pow(num, 3)).toFixed(2) + "G";
        return (size / Math.pow(num, 4)).toFixed(2) + "T";
    }
	$(document).ready(function() {
		var baseurl = "https://archapi.zhcn.cc/api/v1";
		var url = baseurl + "/version/";
		$.ajax({
			url: url,
			dataType: "json",
			success:function(result) {
                $('#version').text(result.version);
                $('#kernel').text(result.kernel);
                $('#size').text(getfilesize(result.size));
                $("#a_torrent").attr("href", "https://mirrors.wsyu.edu.cn/loongarch/archlinux/iso/latest/" + result.iso_file + ".torrent")
                .html(result.version);
                $("#a_magnet").attr("href", "magnet:?xt=urn:btih:" + result.bthash +"&dn="+ result.iso_file)
                .html(result.version);
                for(var i=0; i<result.mirrors.length; i++) {
                    let mirror = result.mirrors[i];
                    let uri = new URL(mirror);
                    let url = uri.protocol + "//" + uri.host;
                    let iso_file = mirror + "/iso/" + result.version + "/" + result.iso_file;
                    let livecd_file = mirror + "/iso/" + result.version + "/" + result.livecd_file;
                    $li_url = $("<li><a href='"+ mirror + "/iso/latest/' target='_blank'>"+ url +"</a>&nbsp;&nbsp;<a href='"+ iso_file +"' target='_blank'><img width='12' height='12' src='/images/download.png' alt=''/>ISO</a>&nbsp;&nbsp;<a href='" + livecd_file + "' target='_blank'><img width='12' height='12' src='/images/download.png' alt=''/>LiveCD</a></li>");
                    $("#ul_download").append($li_url);
                }
			}
		});
	});
</script>
