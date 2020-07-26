package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type file struct {
	name string
	path string
}

func main() {
	// 1. Collect all the files to rename.
	// 2. Iterate over files to generate new name and invoke rename
	// with system call.

	dir := "sample"
	var toRename []file

	// 1. Collect all the files to rename.
	// look over 'dir' recursively
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			// do not rename a directory
			return nil
		}
		// append all files and their absolute path from the directory to
		// toRename slice
		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, file{
				name: info.Name(),
				path: path,
			})
		}
		return nil
	})

	// 2. Iterate over files to generate new name and invoke rename
	// with system call.
	for _, orig := range toRename {
		var n file
		var err error
		n.name, err = match(orig.name)
		if err != nil {
			fmt.Println("Error matching:", orig.path, err.Error())
		}
		n.path = filepath.Join(dir, n.name)
		fmt.Printf("mv %s => %s\n", orig.path, n.path)
		// err = os.Rename(orig.path, n.path)
		// if err != nil {
		// 	fmt.Println("Error renaming:", orig.path, err.Error())
		// }
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
		return nil, fmt.Errorf("%s didn't match our pattern")
	}

	return &matchResult{strings.Title(name), number, ext}, nil
}
