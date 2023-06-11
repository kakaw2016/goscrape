package main

import (
	"fmt"
)

func main() {

	i := 0

	for i <= 100 {
		if i%2 == 0 {
			fmt.Println(i, "bonjour yohann et allan")
		}
		if i%7 == 0 {
			fmt.Println(i, "bonjour yohann")
		}
		if i%4 == 0 {
			fmt.Println(i, "bonjour allan")
		}

		if i%34 == 0 {
			fmt.Println(i, "bonjour tonton Olivier")
		}

		i = i + 1

	}
}
