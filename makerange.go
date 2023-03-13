package piscine

func MakeRange(min, max int) []int {
	var tab []int
	if min < max {
		size := max - min
		tab = make([]int, size)
		for i := 0; i < size; i++ {
			tab[i] = i + min
		}
	}
	return tab
}
