package piscine

func UltimateDivMod(a *int, b *int) {
	tmp1 := *a / *b
	tmp2 := *a % *b
	*a = tmp1
	*b = tmp2
}
