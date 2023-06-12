package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func parseOffset(page string) (limit, offset int) {
	limit = 100
	p, e := strconv.Atoi(page)
	if e != nil {
		return limit, 0
	}

	offset = (p - 1) * limit

	return
}

func parseSort(c *gin.Context) (order string) {
	sort := c.DefaultQuery("sort", "pkgname")
	if sort == "pkgname" {
		order = "packages.name"
	} else if sort == "-pkgname" {
		order = "packages.name desc"
	} else if sort == "repo" {
		order = "Repo__name"
	} else if sort == "-repo" {
		order = "Repo__name desc"
	} else if sort == "arch" {
		order = "Arch__name"
	} else if sort == "-arch" {
		order = "Arch__name desc"
	} else if sort == "last_update" {
		order = "packages.build_date"
	} else if sort == "-last_update" {
		order = "packages.build_date desc"
	}
	return
}

func pageCount(total int64, limit int) int {
	var num int
	count := int(total)
	remainder := count % limit
	if remainder > 0 {
		num = (count / limit) + 1
	} else {
		num = count / limit
	}
	return num
}

func MirrorList(c *gin.Context) {
	al, ok := c.MustGet("archloong").(*ArchLoong)
	if !ok {
		fmt.Printf("%#v\n", al.db)
	}
	c.JSON(http.StatusOK, al.cfg.Mirrors)
}

func Version(c *gin.Context) {
	al, ok := c.MustGet("archloong").(*ArchLoong)
	if !ok {
		fmt.Printf("%#v\n", al.db)
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	type Release struct {
		Version string   `yaml:"version" json:"version"`
		IsoFile string   `yaml:"isofile" json:"iso_file"`
		TarFile string   `yaml:"tarfile" json:"tar_file"`
		Kernel  string   `yaml:"kernel" json:"kernel"`
		Size    int      `yaml:"size" json:"size"`
		BtHash  string   `yaml:"bthash" json:"bthash"`
		Mirrors []string `json:"mirrors"`
	}
	var release *Release

	data := al.getURLBytes("/iso/latest/arch/version.yaml")
	err := yaml.Unmarshal(data, &release)
	if err != nil {
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	release.Mirrors = al.cfg.Mirrors
	c.JSON(http.StatusOK, release)
}

func OnePackage(c *gin.Context) {
	al, ok := c.MustGet("archloong").(*ArchLoong)
	if !ok {
		fmt.Printf("%#v\n", al.db)
	}
	var pkg Package
	id := c.Query("id")
	repo := c.Query("repo")
	arch := c.Query("arch")
	name := c.Query("name")
	if id != "" {
		al.db.Joins("Arch").Joins("Repo").Joins("Base").First(&pkg, id)
	} else if repo != "" && arch != "" && name != "" {
		var r Repo
		var a Arch
		al.db.Where("name = ?", arch).First(&a)
		al.db.Where("name = ?", repo).First(&r)
		pkg.Name = name
		pkg.Repo = r
		pkg.RepoID = int(r.ID)
		pkg.Arch = a
		pkg.ArchID = int(a.ID)
		al.db.Joins("Arch").Joins("Repo").Joins("Base").Where(pkg).First(&pkg)
	}
	c.JSON(http.StatusOK, gin.H{
		"id":              pkg.ID,
		"pkgname":         pkg.Name,
		"pkgbase":         pkg.Base.Name,
		"repo":            pkg.Repo.Name,
		"arch":            pkg.Arch.Name,
		"pkgver":          pkg.Version,
		"pkgdesc":         pkg.Description,
		"url":             pkg.URL,
		"filename":        pkg.FileName,
		"compressed_size": pkg.CompressedSize,
		"installed_size":  pkg.InstalledSize,
		"build_date":      pkg.BuildDate,
		"packager":        pkg.Packager,
		"groups":          pkg.Groups,
		"licenses":        pkg.Licenses,
		"conflicts":       pkg.Conflicts,
		"provides":        pkg.Provides,
		"replaces":        pkg.Replaces,
		"depends":         pkg.Depends,
		"optdepends":      pkg.OptDepends,
		"makedepends":     pkg.MakeDepends,
		"checkdepends":    pkg.CheckDepends,
	})
}

func OnePackageFiles(c *gin.Context) {
	al, ok := c.MustGet("archloong").(*ArchLoong)
	if !ok {
		fmt.Printf("%#v\n", al.db)
	}
	repo := c.Query("repo")
	arch := c.Query("arch")
	name := c.Query("name")
	var pkg Package
	if repo != "" && arch != "" && name != "" {
		var r Repo
		var a Arch
		al.db.Where("name = ?", arch).First(&a)
		al.db.Where("name = ?", repo).First(&r)
		pkg.Name = name
		pkg.Repo = r
		pkg.RepoID = int(r.ID)
		pkg.Arch = a
		pkg.ArchID = int(a.ID)
		al.db.Joins("Arch").Joins("Repo").Joins("Base").Select("packages.version").Where(pkg).First(&pkg)
		fmt.Printf("%#v\n", pkg.Version)

		pkgver := pkg.Name + "-" + pkg.Version
		if flist, err := al.getPkgFiles(r.Name, pkgver); err == nil {
			files_count := len(flist)
			var dir_count int
			for i := 0; i < files_count; i++ {
				if strings.HasSuffix(flist[i], "/") {
					dir_count += 1
				}
			}
			files := PackageFiles{
				Name:       name,
				Repo:       r.Name,
				Arch:       a.Name,
				Version:    pkg.Version,
				FilesCount: files_count,
				DirCount:   dir_count,
				Files:      flist,
			}
			c.JSON(http.StatusOK, files)
		}
	}
}

func PackageDownload(c *gin.Context) {
	al, ok := c.MustGet("archloong").(*ArchLoong)
	if !ok {
		fmt.Printf("%#v\n", al.db)
	}
	var pkg Package
	repo := c.Param("repo")
	arch := c.Param("arch")
	name := c.Param("name")
	if repo != "" && arch != "" && name != "" {
		var r Repo
		var a Arch
		al.db.Where("name = ?", arch).First(&a)
		al.db.Where("name = ?", repo).First(&r)
		pkg.Name = name
		pkg.Repo = r
		pkg.RepoID = int(r.ID)
		pkg.Arch = a
		pkg.ArchID = int(a.ID)
		al.db.Joins("Arch").Joins("Repo").Joins("Base").Select("packages.file_name").Where(pkg).First(&pkg)

		randomIndex := rand.Intn(len(al.cfg.Mirrors))
		mirror := al.cfg.Mirrors[randomIndex]
		url := fmt.Sprintf("%s/%s/os/loong64/%s", mirror, repo, pkg.FileName)
		c.Redirect(http.StatusFound, url)
	}
}

func PackageList(c *gin.Context) {
	al, ok := c.MustGet("archloong").(*ArchLoong)
	if !ok {
		fmt.Printf("%#v\n", al.db)
	}

	type Line struct {
		ID          uint      `json:"id"`
		Name        string    `json:"pkgname"`
		Repo        string    `json:"repo"`
		Arch        string    `json:"arch"`
		Version     string    `json:"version"`
		Description string    `json:"pkgdesc"`
		BuildDate   time.Time `json:"build_date"`
	}

	q := c.Query("q")
	arch := c.Query("arch")
	repo := c.Query("repo")

	fmt.Printf("repo: %#v\n", repo)

	page := c.DefaultQuery("page", "1")
	limit, offset := parseOffset(page)
	order := parseSort(c)

	var exacts []Line
	var lists []Line
	var count int64

	if q != "" {
		var pkgs []Package
		var pkg Package
		pkg.Name = q
		if repo != "" || arch != "" {
			var r Repo
			var a Arch
			if repo != "" {
				al.db.Where("name = ?", repo).First(&r)
				pkg.Repo = r
				pkg.RepoID = int(r.ID)
			}
			if arch != "" {
				al.db.Where("name = ?", arch).First(&a)
				pkg.Arch = a
				pkg.ArchID = int(a.ID)
			}
		}
		al.db.Debug().Joins("Arch").Joins("Repo").Select("packages.id", "packages.name", "packages.version", "packages.description", "packages.build_date").Where(pkg).Find(&pkgs)
		for i := 0; i < len(pkgs); i++ {
			line := Line{
				ID:          pkgs[i].ID,
				Name:        pkgs[i].Name,
				Repo:        pkgs[i].Repo.Name,
				Arch:        pkgs[i].Arch.Name,
				Version:     pkgs[i].Version,
				Description: pkgs[i].Description,
				BuildDate:   pkgs[i].BuildDate,
			}
			exacts = append(exacts, line)
		}
		pkg.Name = ""

		al.db.Debug().Joins("Arch").Joins("Repo").Select("packages.id", "packages.name", "packages.version", "packages.description", "packages.build_date").Where(pkg).Where("packages.name like ?", "%"+q+"%").Find(&pkgs).Count(&count)
		al.db.Debug().Joins("Arch").Joins("Repo").Select("packages.id", "packages.name", "packages.version", "packages.description", "packages.build_date").Where(pkg).Where("packages.name like ?", "%"+q+"%").Order(order).Limit(limit).Offset(offset).Find(&pkgs)

		for i := 0; i < len(pkgs); i++ {
			line := Line{
				ID:          pkgs[i].ID,
				Name:        pkgs[i].Name,
				Repo:        pkgs[i].Repo.Name,
				Arch:        pkgs[i].Arch.Name,
				Version:     pkgs[i].Version,
				Description: pkgs[i].Description,
				BuildDate:   pkgs[i].BuildDate,
			}
			lists = append(lists, line)
		}

	} else {
		var pkgs []Package
		var pkg Package
		if repo != "" {
			var r Repo
			if repo != "" {
				al.db.Where("name = ?", repo).First(&r)
				pkg.Repo = r
				pkg.RepoID = int(r.ID)
			}
		}
		if arch != "" {
			var a Arch
			if arch != "" {
				al.db.Where("name = ?", arch).First(&a)
				pkg.Arch = a
				pkg.ArchID = int(a.ID)
			}
		}
		al.db.Debug().Joins("Arch").Joins("Repo").Select("packages.id", "packages.name", "packages.version", "packages.description", "packages.build_date").Where(&pkg).Find(&pkgs).Count(&count)
		al.db.Debug().Joins("Arch").Joins("Repo").Select("packages.id", "packages.name", "packages.version", "packages.description", "packages.build_date").Where(&pkg).Order(order).Limit(limit).Offset(offset).Find(&pkgs)

		for i := 0; i < len(pkgs); i++ {
			line := Line{
				ID:          pkgs[i].ID,
				Name:        pkgs[i].Name,
				Repo:        pkgs[i].Repo.Name,
				Arch:        pkgs[i].Arch.Name,
				Version:     pkgs[i].Version,
				Description: pkgs[i].Description,
				BuildDate:   pkgs[i].BuildDate,
			}
			lists = append(lists, line)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"page_num":   page,
		"page_count": pageCount(count, limit),
		"results":    count,
		"exacts":     exacts,
		"packages":   lists,
	})
}

// ArchLoongMiddleware will add the ArchLoong to the context
func ArchLoongMiddleware(al *ArchLoong) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("archloong", al)
		c.Next()
	}
}

func (al *ArchLoong) WebServer() {
	if gin.Mode() == gin.DebugMode {
		al.route.Use(cors.Default())
	} else {
		al.route.Use(cors.New(cors.Config{
			AllowOrigins:     al.cfg.Origins,
			AllowMethods:     []string{"GET", "PATCH", "HEAD", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
		}))
	}
	al.route.Use(ArchLoongMiddleware(al))

	v1 := al.route.Group("/api/v1")
	{
		v1.GET("/mirrors/", MirrorList)
		v1.GET("/version/", Version)
		v1.GET("/packages/", PackageList)
		v1.GET("/package/", OnePackage)
		v1.GET("/package/files/", OnePackageFiles)
		v1.GET("/:repo/:arch/:name/download/", PackageDownload)
	}
	al.route.Run(":" + strconv.Itoa(al.cfg.Port))
}
