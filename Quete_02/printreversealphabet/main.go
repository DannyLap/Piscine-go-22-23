package main

import "github.com/01-edu/z01"

func main() {
	for tmp := 'z'; tmp >= 'a'; tmp-- {
		z01.PrintRune(tmp)
	}
	z01.PrintRune('\n')
}
