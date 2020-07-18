package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// fileName := "birthday_001.txt"
	// // => Birthday - 1 of 4.txt
	// newName, err := match(fileName, 4)
	// if err != nil {
	// fmt.Println("no match")
	// os.Exit(1)
	// }
	// fmt.Println(newName)
	dir := "./sample"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	count := 0
	type rename struct {
		filename string
		path     string
	}
	var toRename []string

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Dir:", file.Name())
		} else {
			_, err := match(file.Name(), 4)
			if err == nil {
				count++
				toRename = append(toRename, file.Name())
			}
		}
	}
	for _, orig := range toRename {
		origPath := filepath.Join(dir, orig)
		newFilename, err := match(orig, count)
		if err != nil {
			panic(err)
		}
		newPath := filepath.Join(dir, newFilename)
		err := os.Rename(origPath, newPath)
		if err != nil {
			panic(err)
		}
		fmt.Printf("mv %s => %s\n", origPath, newPath)
	}
}

// match returns new file name, or an error if the filename
// didn't match our pattern
func match(fileName string, total int) (string, error) {
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

	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, ext), nil
}
