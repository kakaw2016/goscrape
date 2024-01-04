package main

import (
	"fmt"
)

func selection1(x int, y []string, z []string) string {

	d1 := fmt.Sprintln("Entrer le chiffre:", x)

	d2 := fmt.Sprintln("Voici la voiture:", y[x-1])

	d3 := fmt.Sprintln("Votre couleur est celle-ci:", z[x-1])

	return d1 + d2 + d3

}

func main() {

	listeVoiture := []string{"Ford", "Hundai", "Toyota", "Ferrari"}

	couleurVoiture := []string{"Jaune", "Blanc", "Bleu", "Noir"}

	resultat := selection1(1, listeVoiture, couleurVoiture)

	fmt.Print(resultat)

}
