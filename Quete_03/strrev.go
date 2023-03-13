package piscine

func StrRev(s string) string {
	var x string
	for i := 0; i <= len(s)-1; i++ {
		x = string(s[i]) + x
	}
	return x
}
