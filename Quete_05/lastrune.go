package piscine

func LastRune(s string) rune {
	for index, character := range s {
		if index == len(s)-1 {
			return character
		}
	}
	return '\n'
}
