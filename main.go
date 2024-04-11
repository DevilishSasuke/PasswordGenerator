package main

import (
	"fmt"
	"math/rand"
	"pswrdgen/passgen"
)

func main() {
	length := rand.Intn(20) + 10
	charset := passgen.ComposeAlphabet(true, true)
	pswrd := passgen.GeneratePassword(length, charset)
	fmt.Println(string(pswrd))
}
