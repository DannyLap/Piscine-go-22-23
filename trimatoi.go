package piscine

func TrimAtoi(s string) int {
	compteur := true
	compteur2 := false
	var new_nombre int

	for i := 0; i < len(s); i++ {
		if rune(s[i]) == '-' && compteur {
			compteur2 = true
		} else if IsNumeric(string(s[i])) {
			compteur = false
		}
	}

	for i := 0; i < len(s); i++ {
		if IsNumeric(string(s[i])) {
			new_nombre = (new_nombre + (int(s[i]) - 48)) * 10
		}
	}
	new_nombre = new_nombre / 10

	if compteur2 {
		return new_nombre * -1
	} else {
		return new_nombre
	}
}
