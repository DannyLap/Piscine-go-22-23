package piscine

func IsAlpha(s string) bool {
	compteur := 0
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= rune(65) && rune(s[i]) <= rune(90) || rune(s[i]) >= rune(97) && rune(s[i]) <= rune(122) || rune(s[i]) >= rune(48) && rune(s[i]) <= rune(57) {
			compteur += 0
		} else {
			return false
		}
	}
	return true
}
