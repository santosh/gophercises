package main

import (
	"fmt"

	"github.com/santosh/gophercises/secret"
)

func main() {
	v := secret.File("my-fake-key", ".secrets")
	err := v.Set("demo_key", "some crazy value")
	if err != nil {
		panic(err)
	}
	plain, err := v.Get("demo_key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)
}
