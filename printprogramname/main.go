package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := []rune(os.Args[0])[2:]
	for _, c := range name {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
