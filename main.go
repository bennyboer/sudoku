package main

import (
	"fmt"
	emptyGenerator "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/generator/empty"
)

func main() {
	generator1 := emptyGenerator.EmptyGenerator{}
	empty := generator1.Generate(0.0)

	fmt.Printf("%v", empty)
}
