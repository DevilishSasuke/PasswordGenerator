package main

import (
	"fmt"
	"math/rand"
	"pswrdgen/passgen"
)

func main() {
	length := rand.Intn(150) + 10
	charset := passgen.ComposeAlphabet(true, true)
	pswrd, err := passgen.GeneratePassword(length, &charset)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(pswrd))
}
