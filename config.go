package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/yetist/xmppbot/utils"
)

type Config struct {
	AppName    string
	AppVersion string
	AppConfig  string
	Mirrors    []string `toml:"mirrors"`
	Origins    []string `toml:"origins"`
	RepoDir    string   `toml:"repodir"`
	Sqlite     string   `toml:"sqlite"`
	Port       int      `toml:"web_port"`
	Mode       string   `toml:"web_mode"`
}

var DefaultConfig = Config{
	Mirrors: []string{
		"https://mirrors.wsyu.edu.cn/loongarch/archlinux",
		"https://mirrors.pku.edu.cn/loongarch/archlinux",
		"https://mirrors.nju.edu.cn/loongarch/archlinux",
		"https://mirror.iscas.ac.cn/loongarch/archlinux",
	},
	Origins: []string{
		"https://loongarchlinux.org",
	},
	RepoDir: utils.ExpandUser("~/.cache/archloong"),
	Sqlite:  utils.ExpandUser("~/.cache/archloong.db"),
	Port:    8080,
	Mode:    "debug",
}

func selfConfigDir() string {
	if dir, err := utils.GetExecDir(); err != nil || strings.HasSuffix(dir, "_obj/exe") {
		wd, _ := os.Getwd()
		return wd
	} else {
		return dir
	}
}

func userConfigDir(name, version string) (pth string) {
	if pth = os.Getenv("XDG_CONFIG_HOME"); pth == "" {
		pth = utils.ExpandUser("~/.config")
	}

	if name != "" {
		pth = filepath.Join(pth, name)
	}

	if version != "" {
		pth = filepath.Join(pth, version)
	}

	return pth
}

func sysConfigDir(name, version string) (pth string) {
	if pth = os.Getenv("XDG_CONFIG_DIRS"); pth == "" {
		pth = "/etc/xdg"
	} else {
		pth = utils.ExpandUser(filepath.SplitList(pth)[0])
	}
	if name != "" {
		pth = filepath.Join(pth, name)
	}

	if version != "" {
		pth = filepath.Join(pth, version)
	}
	return pth
}

func LoadConfig(name, version, cfgname string) (*Config, error) {
	config := DefaultConfig

	sysconf := path.Join(sysConfigDir(name, version), cfgname)
	userconf := path.Join(userConfigDir(name, version), cfgname)
	selfconf := path.Join(selfConfigDir(), cfgname)
	cwdconf := path.Join(utils.CwdDir(), cfgname)
	defer func() {
		config.AppName = name
		config.AppVersion = version
		config.AppConfig = cfgname
	}()

	if utils.IsFile(cwdconf) {
		if _, err := toml.DecodeFile(cwdconf, &config); err != nil {
			print(err)
		}
	} else if utils.IsFile(selfconf) {
		if _, err := toml.DecodeFile(selfconf, &config); err != nil {
			print(err)
		}
	} else if utils.IsFile(userconf) {
		if _, err := toml.DecodeFile(userconf, &config); err != nil {
			print(err)
		}
	} else if utils.IsFile(sysconf) {
		if _, err := toml.DecodeFile(sysconf, &config); err != nil {
			print(err)
		}
	}
	config.RepoDir = utils.ExpandUser(config.RepoDir)
	config.Sqlite = utils.ExpandUser(config.Sqlite)
	return &config, nil
}
