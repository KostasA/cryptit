package main

import (
	"fmt"

	"github.com/KostasA/cryptit/decrypt"
	"github.com/KostasA/cryptit/encrypt"
)

func main() {
	x := encrypt.Nimbus("Kodekloud")
	fmt.Println(x)
	fmt.Println(decrypt.Nimbus(x))

}
