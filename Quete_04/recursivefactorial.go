package piscine

func RecursiveFactorial(nb int) int {
	if nb > 0 && nb < 30 {
		return RecursiveFactorial(nb-1) * nb
	} else if nb == 0 {
		return 1
	} else {
		return 0
	}
}
