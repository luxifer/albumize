package main

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"time"
)

var result = make([]string, 0)

func main() {
	if len(os.Args) <= 1 {
		usage()
	}

	_, err := os.Stat(os.Args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Path %s does not exist\n", os.Args[1])
		os.Exit(2)
	}

	ipath := os.Args[1]
	files, _ := ioutil.ReadDir(ipath)

	for _, f := range files {
		if !f.IsDir() {
			result = append(result, path.Join(ipath, f.Name()))
		}
	}

	layout := "2006:01:02 15:04:05"

	for _, file := range result {
		fh, err := os.Open(file)

		if err != nil {
			panic(err)
		}

		x, err := exif.Decode(fh)

		if err == nil {
			date, _ := x.Get(exif.DateTimeOriginal)
			if date != nil {
				pdate, _ := time.Parse(layout, date.StringVal())
				sy := strconv.Itoa(pdate.Year())
				sd := strconv.Itoa(pdate.Day())
				sm := pdate.Month().String()
				ppath := path.Join(ipath, sy, sm, sd)
				_, fn := path.Split(file)
				os.MkdirAll(ppath, 0755)

				dest := unique(ppath, fn)
				//os.Rename(file, dest)
				fmt.Fprintf(os.Stdout, "Move file %s to %s\n", file, dest)
			}
		}

		fh.Close()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <path>\n", os.Args[0])
	os.Exit(2)
}

func unique(tpath string, filename string) string {
	fullpath := path.Join(tpath, filename)
	_, err := os.Stat(fullpath)

	if os.IsNotExist(err) {
		return fullpath
	}

	return fullpath
}
