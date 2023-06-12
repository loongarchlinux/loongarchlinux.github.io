package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/yetist/xmppbot/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ArchLoong struct {
	cfg   *Config
	cron  *gocron.Scheduler
	route *gin.Engine
	db    *gorm.DB
}

func NewArchLoong(cfg *Config) *ArchLoong {
	al := &ArchLoong{
		cfg:  cfg,
		cron: gocron.NewScheduler(time.Local),
	}
	db, err := gorm.Open(sqlite.Open(cfg.Sqlite), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Attr{})
	db.AutoMigrate(&Repo{})
	db.AutoMigrate(&Arch{})
	db.AutoMigrate(&Base{})
	db.AutoMigrate(&Package{})
	db.AutoMigrate(&PackageFiles{})
	al.db = db

	return al
}

func (al *ArchLoong) GetConfig() *Config {
	return al.cfg
}

func (al *ArchLoong) getURLBytes(path string) []byte {
	var text []byte
	var err error

	for _, host := range al.cfg.Mirrors {
		url := fmt.Sprintf("%s/%s", host, path)
		if text, err = getURL(url); err == nil {
			break
		}
	}
	return text
}

func (al *ArchLoong) getURLString(path string) string {
	var b = al.getURLBytes(path)
	return string(b)
}

func (al *ArchLoong) fetchFile(rpath, lpath string) bool {
	ldir := filepath.Dir(lpath)
	if !utils.IsDir(ldir) {
		if err := os.MkdirAll(ldir, os.ModePerm); err != nil {
			return false
		}
	}

	bytes := al.getURLBytes(rpath)
	if err := os.WriteFile(lpath, bytes, 0644); err != nil {
		return false
	}
	return true
}

func (al *ArchLoong) getLastUpdate() time.Time {
	i := str2int64(al.getURLString("/lastupdate"))
	return time.Unix(i, 0)
}

func (al *ArchLoong) ShouldUpdate() bool {
	var sync Attr
	al.db.FirstOrCreate(&sync, Attr{Name: "lastsync"})
	i := str2int64(sync.Value)
	if i == 0 {
		return true
	}
	local := time.Unix(i, 0)
	remote := al.getLastUpdate()

	if local.Before(remote) {
		return true
	}
	return false
}

func (al *ArchLoong) UpdateMirrors() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	body, err := getURL("https://github.com/loongarchlinux/core/raw/main/pacman-mirrorlist/mirrorlist")
	if err != nil {
		println("error")
		return
	}

	lines := strings.Split(string(body), "\n")
	for _, l := range lines {
		if strings.Index(l, "Server") == -1 {
			continue
		}
		url := strings.Split(l, "=")[1]
		server := strings.Split(url, "$")[0]
		mirror := strings.Trim(server, " /")
		if !strlst_contains(al.cfg.Mirrors, mirror) {
			al.cfg.Mirrors = append(al.cfg.Mirrors, mirror)
		}
	}
}

// 更新数据库
func (al *ArchLoong) Update() {
	if al.ShouldUpdate() {
		al.DBUpdate()
		al.UpdateMirrors()
	}
}

// 运行web server
func (al *ArchLoong) Server() {
	al.cron.Every(1).Day().At("03:00").Do(al.Update)
	al.cron.StartAsync()

	// 启动web server
	if al.cfg.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	al.route = gin.Default()
	al.WebServer()
}
