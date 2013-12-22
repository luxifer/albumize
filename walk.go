package main

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"os"
	"path"
	"strconv"
	"time"
)

func Organize(result []string, ipath string) {
	for _, file := range result {
		fh, err := os.Open(file)

		defer fh.Close()

		if err != nil {
			panic(err)
		}

		x, err := exif.Decode(fh)

		if err == nil {
			date, _ := x.Get(exif.DateTimeOriginal)
			if date != nil {
				pdate, _ := time.Parse(layout, date.StringVal())
				sy := strconv.Itoa(pdate.Year())
				sd := fmt.Sprintf("%02d", pdate.Day())
				sm := pdate.Month().String()
				ppath := path.Join(ipath, sy, sm, sd)
				_, fn := path.Split(file)
				os.MkdirAll(ppath, 0755)

				dest := UniqueFilename(ppath, fn)
				fmt.Fprintf(os.Stdout, "Move file %s to %s\n", file, dest)
				os.Rename(file, dest)
			}
		}
	}
}

func init() {

}
