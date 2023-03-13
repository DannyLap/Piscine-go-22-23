package piscine

func IsUpper(s string) bool {
	compteur := 0
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= rune(65) && rune(s[i]) <= rune(90) {
			compteur += 0
		} else {
			return false
		}
	}
	return true
}
