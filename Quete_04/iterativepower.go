package piscine

func IterativePower(nb int, power int) int {
	resultat := 1
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		for i := 0; i < power; i++ {
			resultat = resultat * nb
		}
		return resultat
	}
}
