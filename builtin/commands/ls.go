package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Ls(argv []string) error {
	location := "."
	var locations []string
	switch len(argv) {
	case 0, 1:
	case 2:
		location = argv[1]
	default:
		locations = argv[1:]
	}
	if locations != nil {
		for i, location := range locations {
			if err := printLocation(location, true); err != nil {
				return nil
			}
			if i != len(locations)-1 {
				fmt.Println()
			}
		}
		return nil
	}
	return printLocation(location, false)
}

func printLocation(location string, printDirLabel bool) (err error) {
	fi, err := os.Stat(location)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		printFileInfo(fi)
		return nil
	}
	if printDirLabel {
		fmt.Println(fi.Name() + "/:")
	}
	files, err := ioutil.ReadDir(location)
	if err != nil {
		return err
	}
	for _, file := range files {
		if !strings.HasPrefix(file.Name(), ".") {
			printFileInfo(file)
		}
	}
	return nil
}

func printFileInfo(fi os.FileInfo) {
	name := fi.Name()
	if fi.IsDir() {
		name += "/"
	}
	fmt.Println(name)
}
