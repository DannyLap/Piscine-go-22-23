package piscine

func StrLen(s string) int {
	tmp := len(s)
	for i := 0; i < tmp; i++ {
		if s[i] == 'Ã' {
			tmp -= 1
		}
	}
	return tmp
}
