package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var re = regexp.MustCompile("^(.*?) ([0-9]{4}) [(]([0-9]+) of ([0-9]+)[)][.](.+?)$")
var replaceString = "$2 - $1 - $3 of $4.$5"

func main() {
	var dry bool
	flag.BoolVar(&dry, "dry", true, "whether or not this should be a real or dry run")
	flag.Parse()

	// 1. Collect all the files to rename.
	walkDir := "sample"
	var toRename []string

	// 1. Collect all the files to rename.
	// look over 'walkDir' recursively
	filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			// do not rename a directory
			return nil
		}

		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, path)
		}

		return nil
	})

	// 2. Iterate over files to generate new name and invoke rename
	// with a system call.
	for _, oldPath := range toRename {
		dir := filepath.Dir(oldPath)
		filename := filepath.Base(oldPath)
		newFilename, _ := match(filename)
		newPath := filepath.Join(dir, newFilename)

		fmt.Printf("mv %s => %s\n", oldPath, newPath)
		if !dry {
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println("Error renaming:", oldPath, newPath, err.Error())
			}
		}
	}
}

// match takes a fileName and tries to match is with 're' regex,
// if matched, it replaces/renames it with replaceString
// throws an error if regex doesn't match.
func match(fileName string) (string, error) {
	if !re.MatchString(fileName) {
		return "", fmt.Errorf("%s didn't match our pattern", fileName)
	}

	return re.ReplaceAllString(fileName, replaceString), nil
}
