package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func UniqueFilename(tpath string, filename string) string {
	fullpath := path.Join(tpath, filename)
	_, err := os.Stat(fullpath)

	if os.IsNotExist(err) {
		return fullpath
	}

	i := 0
	ext := filepath.Ext(filename)
	wefn := strings.Replace(filename, ext, "", -1)

	for os.IsExist(err) {
		newfn := fmt.Sprintf("%s_%d%s", wefn, i, ext)
		newpath := path.Join(tpath, newfn)
		fullpath = newpath
		_, err = os.Stat(newpath)
		i++
	}

	return fullpath
}

func init() {

}
