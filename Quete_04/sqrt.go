package piscine

func Sqrt(nb int) int {
	i := 1
	for i = 1; i <= nb/2+nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}
