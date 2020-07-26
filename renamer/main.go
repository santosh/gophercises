package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var dry bool
	flag.BoolVar(&dry, "dry", true, "whether or not this should be a real or dry run")
	flag.Parse()

	// 1. Collect all the files to rename.
	// 2. Iterate over files to generate new name and invoke rename
	// with system call.

	walkDir := "sample"
	toRename := make(map[string][]string)

	// 1. Collect all the files to rename.
	// look over 'walkDir' recursively
	filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			// do not rename a directory
			return nil
		}

		curDir := filepath.Dir(path)
		// collect all files and their absolute path from the directory to
		// toRename slice
		if m, err := match(info.Name()); err == nil {
			// example key: sample/picture.jpg
			// example key: sample/subdir/picture.jpg
			// example val: [picture_001.txt picture_002.txt picture_003.txt]
			key := filepath.Join(curDir, fmt.Sprintf("%s.%s", m.base, m.ext))
			toRename[key] = append(toRename[key], info.Name())
		}
		return nil
	})

	// 2. Iterate over files to generate new name and invoke rename
	// with system call.
	for key, files := range toRename {
		dir := filepath.Dir(key)
		n := len(files)
		sort.Strings(files)

		// rename on the basis of index
		for fileIndex, filename := range files {
			res, _ := match(filename)
			newFilename := fmt.Sprintf("%s - %d of %d.%s", res.base, (fileIndex + 1), n, res.ext)
			oldPath := filepath.Join(dir, filename)
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
}

type matchResult struct {
	base  string
	index int
	ext   string
}

// match takes a fileName and returns a struct of segments which will
// used to compose a new file name.
// returns an error if the filename doesn't match our pattern
func match(fileName string) (*matchResult, error) {
	// split by '.'; note that there can be more than
	// two parts at this point. e.g. 'san.tosh.k.jpg'
	pieces := strings.Split(fileName, ".")

	// take the last element of the slice
	ext := pieces[len(pieces)-1]

	// join the rest of the elements back
	fileName = strings.Join(pieces[0:len(pieces)-1], ".")

	// this time split by '_'
	pieces = strings.Split(fileName, "_")

	// join back all elements except the last
	name := strings.Join(pieces[0:len(pieces)-1], "_")

	// remove any padding from the last element of
	number, err := strconv.Atoi(pieces[len(pieces)-1])

	if err != nil {
		return nil, fmt.Errorf("%s didn't match our pattern", fileName)
	}

	return &matchResult{strings.Title(name), number, ext}, nil
}
