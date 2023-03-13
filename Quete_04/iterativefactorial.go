package piscine

func IterativeFactorial(nb int) int {
	if nb >= 0 && nb < 30 {
		resultat := 1
		for i := 1; i <= nb; i++ {
			resultat = resultat * i
		}
		return resultat
	} else {
		return 0
	}
}
