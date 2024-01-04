package main

import (
	"fmt"
	"time"
)

func multiplication(total int) {

	var multiplicant int

	dividerlist := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, total}

	var goodresult []int

	for _, divider := range dividerlist {

		if total%divider == 0 {
			multiplicant = total / divider

			//goodresult[0] = divider

			//goodresult[1] = multiplicant

			goodresult = []int{divider, multiplicant}

			fmt.Printf("%d X %d = %d\n", goodresult[0], goodresult[1], total)

		}

	}

	return

}

func main() {

	fmt.Println("Debut de l'excercice.")

	fmt.Println("Entre la valeur avec laquelle tu veux t'excercer")

	var valeurEtude int

	fmt.Scanln(&valeurEtude)

	fmt.Println("Voici les resultats")

	time.Sleep(2 * time.Second)

	multiplication(valeurEtude)

	time.Sleep(1 * time.Second)

	fmt.Println("Merci pour votre attention")

}
