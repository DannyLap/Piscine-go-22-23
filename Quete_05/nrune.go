package piscine

func NRune(s string, n int) rune {
	for index, character := range s {
		if index == n-1 {
			return character
		}
	}
	return '\x00'
}
