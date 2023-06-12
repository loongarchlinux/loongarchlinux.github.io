package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ISO struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	Kernel  string `json:"kernel_version"`
	Size    string `json:"iso_size"`
	Sha256  string `json:"sha256_sum"`
}

type Updater interface {
	Update()
}

func init() {
}

func newISO() *ISO {
	return &ISO{}
}

func (i *ISO) Update() {
	i.Version = "aaa"
	i.Kernel = "aaa"
	i.Size = "aaa"
	i.Sha256 = "aaa"
}

type Module interface {
}

func releaseISO(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version":        "2028.05.08",
		"kernel_version": "6.3.0",
		"iso_size":       "592Mb",
		"sha256_sum":     "fc5def8d5f78f5b9699072f6c4396241de0f9922866db8049a5d785457d30400",
	})
}

func releaseVM(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"path": c.FullPath(),
	})
}
