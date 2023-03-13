package piscine

import "github.com/01-edu/z01"

func PrintNbr(nombre int) {
	if nombre > 0 && nombre <= 9223372036854775807 {
		new_nombre := nombre
		compteur := 1
		for new_nombre > 9 {
			new_nombre = new_nombre / 10
			compteur *= 10
		}
		new_nombre = nombre
		for compteur+1 > 1 {
			rest := new_nombre / compteur
			new_nombre = new_nombre % compteur
			if rest == 1 {
				z01.PrintRune('1')
			} else if rest == 2 {
				z01.PrintRune('2')
			} else {
				z01.PrintRune(rune(rest + 48))
			}
			compteur /= 10
		}
	} else if nombre == 0 {
		z01.PrintRune('0')
	} else if nombre < 0 && nombre > -9223372036854775808 {
		nombre = nombre * -1
		z01.PrintRune('-')
		PrintNbr(nombre)
	} else if nombre == -9223372036854775808 {
		nombre = (nombre + 8) / 10 * -1
		z01.PrintRune('-')
		PrintNbr(nombre)
		z01.PrintRune('8')
	}
}
