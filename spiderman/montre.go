package main

import (
	"fmt"
)

func main() {

	var lesmontres []string

	lesmontres = ["colmi","sport","spiderman","ecrivonslebonheur"]

	for _, monchoix:= range lesmontres{
		fmt.Println("J'ai choisi cette montre", monchoix)
	}
	

}