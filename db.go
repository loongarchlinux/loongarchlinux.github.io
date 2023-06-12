package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Jguer/go-alpm/v2"
	"gorm.io/gorm"
)

type Attr struct {
	ID    uint   `gorm:"primarykey"`
	Name  string `gorm:"uniqueIndex"`
	Value string
}

type Repo struct {
	ID          uint       `gorm:"primarykey"`
	Name        string     `gorm:"uniqueIndex"`
	PackageList []*Package `gorm:"foreignKey:RepoID"`
}

type Arch struct {
	ID          uint       `gorm:"primarykey"`
	Name        string     `gorm:"uniqueIndex"`
	PackageList []*Package `gorm:"foreignKey:ArchID"`
}

type Base struct {
	ID          uint       `gorm:"primarykey"`
	Name        string     `gorm:"uniqueIndex"`
	PackageList []*Package `gorm:"foreignKey:BaseID"`
}

type Package struct {
	ID             uint      `gorm:"primarykey" json:"id"`
	Name           string    `gorm:"unique_index:anr" json:"pkgname"`
	Base           Base      `json:"pkgbase"`
	Repo           Repo      `gorm:"unique_index:anr" json:"repo"`
	Arch           Arch      `gorm:"unique_index:anr" json:"arch"`
	Version        string    `json:"version"`
	Description    string    `json:"pkgdesc"`
	URL            string    `json:"url"`
	FileName       string    `json:"filename"`
	CompressedSize int64     `json:"compressed_size"`
	InstalledSize  int64     `json:"installed_size"`
	BuildDate      time.Time `json:"build_date"`
	Packager       string    `json:"packager"`
	Groups         Strlist   `json:"groups"`
	Licenses       Strlist   `json:"licenses"`
	Conflicts      Strlist   `json:"conflicts"`
	Provides       Strlist   `json:"provides"`
	Replaces       Strlist   `json:"replaces"`
	Depends        Strlist   `json:"depends"`
	OptDepends     Strlist   `json:"optdepends"`
	MakeDepends    Strlist   `json:"makedepends"`
	CheckDepends   Strlist   `json:"checkdepends"`
	RepoID         int
	ArchID         int
	BaseID         int
}

type PackageFiles struct {
	Name       string  `gorm:"unique_index:anr" json:"pkgname"`
	Repo       string  `gorm:"unique_index:anr" json:"repo"`
	Arch       string  `gorm:"unique_index:anr" json:"arch"`
	Version    string  `json:"version"`
	FilesCount int     `json:"files_count"`
	DirCount   int     `json:"dir_count"`
	Files      Strlist `json:"files"`
}

var (
	repos = []string{"core",
		"core-testing",
		"extra",
		"extra-testing",
		"laur",
		"laur-testing",
	}
	archs = []string{
		"any",
		"loong64",
	}
)

func (al *ArchLoong) DBUpdate() bool {

	for _, r := range repos {
		al.db.FirstOrCreate(&Repo{}, Repo{Name: r})
	}

	for _, a := range archs {
		al.db.FirstOrCreate(&Arch{}, Arch{Name: a})
	}

	// 永久清空表
	//al.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&Base{})
	//al.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&Package{})
	al.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Base{})
	al.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Package{})

	UpdatePackages(al)

	// update lastsync
	remote := al.getLastUpdate()
	str_sync := strconv.FormatInt(remote.Unix(), 10)

	lastsync := Attr{
		Name:  "lastsync",
		Value: str_sync,
	}

	result := al.db.First(&lastsync)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		al.db.Create(&lastsync)
	} else {
		al.db.Save(&lastsync)
	}

	return true
}

func DownloadRepos(al *ArchLoong) {
	sync_dir := filepath.Join(al.cfg.RepoDir, "sync")

	for _, r := range repos {
		rpath := fmt.Sprintf("/%s/os/loong64/%s.db", r, r)
		lpath := fmt.Sprintf("%s/%s.db", sync_dir, r)
		fmt.Printf("download %s to %s\n", rpath, lpath)
		if !al.fetchFile(rpath, lpath) {
			fmt.Printf("fetch %s failes", rpath)
		}

		rpath = fmt.Sprintf("/%s/os/loong64/%s.files", r, r)
		lpath = fmt.Sprintf("%s/%s.files", sync_dir, r)
		fmt.Printf("download %s to %s\n", rpath, lpath)

		if !al.fetchFile(rpath, lpath) {
			fmt.Printf("fetch %s failes", rpath)
		}
	}
}

func (al *ArchLoong) getPkgFiles(repo, pkg string) ([]string, error) {
	var files []string
	repo_files := filepath.Join(al.cfg.RepoDir, "sync", repo+".files")
	srcFile, err := os.Open(repo_files)
	if err != nil {
		return files, err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return files, err
	}
	defer gr.Close()

	buf := new(strings.Builder)
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return files, err
		}

		if hdr.Typeflag == tar.TypeReg && hdr.Name == pkg+"/files" {
			io.Copy(buf, tr)
			break
		}
	}
	if buf.Len() > 0 {
		files = strings.Split(buf.String(), "\n")
		if files[0] == "%FILES%" {
			files = files[1:]
		}
		if files[len(files)-1] == "" {
			files = files[:len(files)-1]
		}
	}
	return files, nil
}

func UpdatePackages(al *ArchLoong) {
	DownloadRepos(al)
	h, er := alpm.Initialize("/", al.cfg.RepoDir)
	if er != nil {
		fmt.Println(er)
		return
	}
	defer h.Release()

	for _, r := range repos {
		db, _ := h.RegisterSyncDB(r, 0)
		for _, pkg := range db.PkgCache().Slice() {
			var repo Repo
			var arch Arch
			var base Base
			al.db.Where("name = ?", pkg.DB().Name()).First(&repo)
			al.db.Where("name = ?", pkg.Architecture()).First(&arch)
			al.db.FirstOrCreate(&base, Base{Name: pkg.Base()})

			p := Package{
				Name:           pkg.Name(),
				Base:           base,
				Repo:           repo,
				Arch:           arch,
				Version:        pkg.Version(),
				Description:    pkg.Description(),
				URL:            pkg.URL(),
				FileName:       pkg.FileName(),
				CompressedSize: pkg.Size(),
				InstalledSize:  pkg.ISize(),
				BuildDate:      pkg.BuildDate(),
				//LastUpdate:
				Packager:     pkg.Packager(),
				Groups:       pkg.Groups().Slice(),
				Licenses:     pkg.Licenses().Slice(),
				Conflicts:    dep2str_list(pkg.Conflicts().Slice()),
				Provides:     dep2str_list(pkg.Provides().Slice()),
				Replaces:     dep2str_list(pkg.Replaces().Slice()),
				Depends:      dep2str_list(pkg.Depends().Slice()),
				OptDepends:   dep2str_list(pkg.OptionalDepends().Slice()),
				MakeDepends:  dep2str_list(pkg.MakeDepends().Slice()),
				CheckDepends: dep2str_list(pkg.CheckDepends().Slice()),
				//FilesLastUpdate:
			}
			result := al.db.Model(Package{}).Create(&p)
			if result.Error != nil {
				fmt.Printf("insert %s failed, %s\n", pkg.Name(), result.Error)
			}
		}
	}
}
