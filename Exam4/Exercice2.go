package piscine

func ft_missing(nums []int) int {
	new_tab := TrieEx2(nums)
	var missing int

	if new_tab[len(new_tab)-1] != 0 {
		return 0
	}

	for i := 0; i < len(new_tab)-1; i++ {
		if new_tab[i] != new_tab[i+1]+1 {
			return new_tab[i] - 1
		}
	}

	return missing
}

func TrieEx2(tab []int) []int {
	valeur_max := len(tab)
	var new_tab []int
	for j := valeur_max; j >= 0; j-- {
		for i := 0; i < len(tab); i++ {
			if tab[i] == j {
				new_tab = append(new_tab, tab[i])
			}
		}
	}
	return new_tab
}
