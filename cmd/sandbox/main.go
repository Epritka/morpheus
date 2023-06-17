package main

import (
	"fmt"

	"github.com/Epritka/morpheus/example"
)

func main() {
	fmt.Println()
	err := example.Test()
	fmt.Println(err)
	fmt.Println()
	// fmt.Println(example.CreatePerson())
}
