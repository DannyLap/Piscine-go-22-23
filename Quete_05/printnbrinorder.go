package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(nombre int) {
	tableau := []int{}
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
				tableau = append(tableau, 1)
			} else if rest == 2 {
				tableau = append(tableau, 2)
			} else {
				tableau = append(tableau, rest)
			}
			compteur /= 10
		}
		for i := 0; i <= 9; i++ {
			for j := 0; j < len(tableau); j++ {
				if int(i) == tableau[j] {
					z01.PrintRune(rune(i) + 48)
				}
			}
		}

	} else if nombre == 0 {
		z01.PrintRune('0')
	}
}
