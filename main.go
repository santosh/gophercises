package main

import (
	"log"

	hackerrank "github.com/santosh/gophercises/hr"
)

func main() {
	encrypted := hackerrank.CaesarCipher("santosh kumar", 1)
	log.Println(encrypted)
}
