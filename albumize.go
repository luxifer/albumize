package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var (
	result = make([]string, 0)
	layout = "2006:01:02 15:04:05"
)

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

	Organize(result, ipath)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <path>\n", os.Args[0])
	os.Exit(2)
}
