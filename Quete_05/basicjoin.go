package piscine

func BasicJoin(elems []string) string {
	var s string
	for i := 0; i < len(elems); i++ {
		s = s + string(elems[i])
	}
	return s
}
