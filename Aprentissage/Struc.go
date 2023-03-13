package main

import "fmt"

func main() {

	type joueur struct {
		nom             string
		age             int
		couleur_cheveux string
	}

	var player joueur

	player.nom = "Lola"

	player.age = 19

	player.couleur_cheveux = "vert"

	fmt.Println(player.nom)

}
