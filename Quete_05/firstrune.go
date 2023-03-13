package piscine

func FirstRune(s string) rune {
	for index, character := range s {
		if index == 0 {
			return character
		}
	}
	return '\n'
}
