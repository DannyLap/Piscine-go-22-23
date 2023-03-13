package main

import (
	"fmt"
)

func ft_coin(coins []int, amount int) int {
	valeur_max := valeur_max(coins)
	new_tab := TrieEx1(coins, valeur_max)
	fmt.Println(new_tab)
	if amount == 0 {
		return 0
	}

	if len(coins) == 1 {
		if coins[0] == amount {
			return 1
		} else {
			return -1
		}
	}

	for i := 0; i < len(coins); i++ {
		for i := len(coins) - 1; i >= 0; i-- {
			if coins[i] <= amount {

			}
		}
	}
	return 1
}

func valeur_max(tab []int) int {
	valeur_max := tab[0]
	for i := 0; i < len(tab); i++ {
		if tab[i] >= valeur_max {
			valeur_max = tab[i]
		}
	}

	return valeur_max
}

func TrieEx1(tab []int, valeur_max int) []int {
	var new_tab []int
	for i := 0; i <= valeur_max; i++ {
		for j := 0; j < len(tab); j++ {
			if i == tab[j] {
				new_tab = append(new_tab, i)
			}
		}
	}
	return new_tab
}

func main() {
	args := []int{1, 5, 4, 7, 8, 99, 60, 23, 5}
	fmt.Println(ft_coin(args, 15))
}
