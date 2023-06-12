package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Jguer/go-alpm/v2"
)

func str2int64(str string) int64 {
	s := strings.TrimSpace(str)
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func getURL(url string) (body []byte, err error) {
	var resp *http.Response

	resp, err = http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("http get request error:", err)
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body error:", err)
		return
	}
	return
}

func dep2str_list(deps []alpm.Depend) []string {
	slice := []string{}
	for _, v := range deps {
		if v.Mod > 1 {
			slice = append(slice, fmt.Sprintf("%s%s%s", v.Name, v.Mod.String(), v.Version))
		} else {
			slice = append(slice, v.Name)
		}
	}
	return slice
}

func file2str_list(files []alpm.File) []string {
	slice := []string{}
	for _, v := range files {
		slice = append(slice, v.Name)
	}
	return slice
}

func strlst_contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
