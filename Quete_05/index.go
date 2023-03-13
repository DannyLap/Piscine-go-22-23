package piscine

func Index(ss string, toFind string) int {
	for index := 0; index < len(ss)-len(toFind); index++ {
		if ss[index:index+len(toFind)] == toFind {
			return index
		}
	}
	return -1
}
