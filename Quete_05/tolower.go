package piscine

func ToLower(s string) string {
	var s2 string
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= rune(65) && rune(s[i]) <= rune(90) {
			s2 = s2 + string(s[i]+32)
		} else {
			s2 = s2 + string(s[i])
		}
	}
	return s2
}
