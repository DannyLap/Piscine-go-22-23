package piscine

func AlphaCount(s string) int {
	compteur := 0
	for index, character := range s {
		if rune(character) >= 65 && rune(character) <= 90 || rune(character) >= 97 && rune(character) <= 122 {
			compteur++
			if index == len(s) {
				return compteur
			}
		}
	}
	return compteur
}
