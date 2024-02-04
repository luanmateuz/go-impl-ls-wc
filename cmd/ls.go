package cmd

import (
	"fmt"
	"os"
)

var (
	reset = "\033[0m"
	blue  = "\033[34m"
)

func Ls(path string, all bool) {
	if path == "" {
		path = "./"
	}

	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if f.Name()[0] == '.' && !all {
			continue
		}
		if f.IsDir() {
			fmt.Printf(blue+"%s ", f.Name())
		} else {
			fmt.Printf(reset+"%s ", f.Name())
		}
	}
	fmt.Println()
}
