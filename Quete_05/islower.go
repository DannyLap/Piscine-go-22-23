package piscine

func IsLower(s string) bool {
	compteur := 0
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= rune(97) && rune(s[i]) <= rune(122) {
			compteur += 0
		} else {
			return false
		}
	}
	return true
}
