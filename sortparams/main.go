package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		tmp := os.Args[i]
		runee := rune(tmp[0])
		if runee == ' ' {
			for j := 0; j < len(tmp); j++ {
				z01.PrintRune(rune(tmp[j]))
			}
			z01.PrintRune('\n')
		}
	}

	for i := 0; i <= 9; i++ {
		for j := 1; j < len(os.Args); j++ {
			tmp := os.Args[j]
			runee := rune(tmp[0])
			if rune(i+48) == runee {
				for k := 0; k < len(tmp); k++ {
					z01.PrintRune(rune(tmp[k]))
				}
				z01.PrintRune('\n')
			}
		}
	}

	for i := 65; i < 91; i++ {
		for j := 1; j < len(os.Args); j++ {
			tmp := os.Args[j]
			runee := rune(tmp[0])
			if rune(i) == runee {
				for k := 0; k < len(tmp); k++ {
					z01.PrintRune(rune(tmp[k]))
				}
				z01.PrintRune('\n')
			}
		}
	}

	for i := 97; i < 123; i++ {
		for j := 1; j < len(os.Args); j++ {
			tmp := os.Args[j]
			runee := rune(tmp[0])
			if rune(i) == runee {
				for k := 0; k < len(tmp); k++ {
					z01.PrintRune(rune(tmp[k]))
				}
				z01.PrintRune('\n')
			}
		}
	}
}
