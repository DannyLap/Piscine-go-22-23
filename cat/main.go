package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		var fullText [][]byte
		for j := 1; j < len(os.Args); j++ {
			text, err := ioutil.ReadFile(os.Args[j])
			if err != nil {
				s1 := "ERROR: " + err.Error()
				for _, letter := range s1 {
					z01.PrintRune(letter)
				}
				z01.PrintRune('\n')
				os.Exit(1)
			}
			fullText = append(fullText, text)
		}
		for i := 0; i < len(fullText); i++ {
			for k := 0; k < len(fullText[i]); k++ {
				z01.PrintRune(rune(fullText[i][k]))
			}
		}
	}
}
