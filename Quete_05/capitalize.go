package piscine

func Capitalize(s string) string {
	var s2 string
	if s[0] >= 'a' && s[0] <= 'z' {
		s2 = s2 + (string(s[0] - 32))
	} else {
		s2 = s2 + string(s[0])
	}

	for i := 1; i < len(s); i++ {
		if rune(s[i]) >= 'a' && rune(s[i]) <= 'z' {
			if (s[i-1] >= 'a' && s[i-1] <= 'z') || (s[i-1] >= 'A' && s[i-1] <= 'Z') || (s[i-1] >= '0' && s[i-1] <= '9') {
				s2 = s2 + string(s[i])
			} else {
				s2 = s2 + (string(s[i] - 32))
			}
		} else if (rune(s[i]) >= 'A' && rune(s[i]) <= 'Z') && ((rune(s[i-1]) >= 'a' && rune(s[i-1]) <= 'z') || (rune(s[i-1]) >= 'A' && rune(s[i-1]) <= 'Z') || (rune(s[i-1]) >= '0' && rune(s[i-1]) <= '9')) {
			s2 = s2 + (string(s[i] + 32))
		} else {
			s2 = s2 + string(s[i])
		}
	}
	return s2
}
