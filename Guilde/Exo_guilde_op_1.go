package piscine

func Exo_bonus_1(tab []int, cible int) []int {
	sortie := []int{}
	if len(tab) >= 2 {
		for i := 0; i < len(tab); i++ {
			for j := i + 1; j <= len(tab); j++ {
				if tab[i]+tab[j] == cible {
					sortie = append(sortie, i)
					sortie = append(sortie, j)
					return sortie
				}
			}
		}
	}
	return sortie
}
