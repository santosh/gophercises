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
	// fileName := "birthday_001.txt"
	// // => Birthday - 1 of 4.txt
	// newName, err := match(fileName, 4)
	// if err != nil {
	// fmt.Println("no match")
	// os.Exit(1)
	// }
	// fmt.Println(newName)
	dir := "sample"
	var toRename []file
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			// do not rename a directory
			return nil
		}
		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, file{
				name: info.Name(),
				path: path,
			})
		}
		return nil
	})
	for _, orig := range toRename {
		var n file
		var err error
		n.name, err = match(orig.name)
		if err != nil {
			fmt.Println("Error matching:", orig.path, err.Error())
		}
		n.path = filepath.Join(dir, n.name)
		fmt.Printf("mv %s => %s\n", orig.path, n.path)
		err = os.Rename(orig.path, n.path)
		if err != nil {
			fmt.Println("Error renaming:", orig.path, err.Error())
		}
	}
}

// match returns new file name, or an error if the filename
// didn't match our pattern
func match(fileName string) (string, error) {
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

	// get rid of any padding
	number, err := strconv.Atoi(pieces[len(pieces)-1])

	if err != nil {
		return "", fmt.Errorf("%s didn't match our pattern")
	}

	return fmt.Sprintf("%s - %d.%s", strings.Title(name), number, ext), nil
}
