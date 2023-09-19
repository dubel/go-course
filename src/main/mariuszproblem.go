package main

import (
	"fmt"
	"os"

	memguard "github.com/awnumar/memguard"
)

func main() {
	a := memguard.NewBuffer(100)

	fmt.Print("AAAA", a)
	os.Exit(0)
}
