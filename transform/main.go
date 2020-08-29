package main

import (
	"io"
	"os"

	"github.com/santosh/gophercises/transform/primitive"
)

func main() {
	f, err := os.Open("in.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	out, err := primitive.Transform(f, 300)
	if err != nil {
		panic(err)
	}

	os.Remove("out.png")
	outFile, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	io.Copy(outFile, out)
}
