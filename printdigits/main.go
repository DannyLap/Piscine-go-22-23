package main

import "github.com/01-edu/z01"

func main() {
	for tmp := '0'; tmp <= '9'; tmp++ {
		z01.PrintRune(tmp)
	}
	z01.PrintRune('\n')
}
