package main

import (
	"fmt"
)

func main() {

	var prenom []string

	prenom = []string{"Yohann", "Allan", "Olivier"}

	var x, y, z int

	fmt.Println("Le premier Numbre:")

	fmt.Scanln(&x)

	fmt.Println("Le deuxieme Numbre:")

	fmt.Scanln(&y)

	z = x - y

	if z%2 == 0 {

		fmt.Println(prenom[0])

	} else if z%3 == 0 {

		fmt.Println(prenom[1])

	} else if z%5 == 0 {

		fmt.Println(prenom[2])
	} else {
		fmt.Println("Une personne inconnue")
	}

}
