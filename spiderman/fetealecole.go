package main

import (
	"fmt"
)

type datax struct {
	Nom  string
	Age  int
	Jour string
}

func maitresse(x int, singleData datax) string {
	var anniversaire string

	if singleData.Age <= x {

		anniversaire = fmt.Sprintf("Il y a un anniversaire.\nLa fete de %s, est le %s\n", singleData.Nom, singleData.Jour)
	}

	if singleData.Age > x {
		anniversaire = fmt.Sprintln("Il n'y a pas danniversaire")

	}

	return anniversaire

}

func main() {

	/*data1 := datax{
		Nom:  "A",
		Age:  10,
		Jour: "31janvier2013",
	}

	data2 := datax{
		Nom:  "Y",
		Age:  7,
		Jour: "2Mars2016",
	}

	data3 := datax{
		Nom:  "C",
		Age:  10,
		Jour: "29Avril2013",
	}

	data4 := datax{
		Nom:  "Al",
		Age:  5,
		Jour: "2juin2018",
	}*/

	dataAll := []datax{
		{
			Nom:  "A",
			Age:  10,
			Jour: "31janvier2013",
		},
		{
			Nom:  "Y",
			Age:  7,
			Jour: "2Mars2016",
		},
		{
			Nom:  "C",
			Age:  10,
			Jour: "29Avril2013",
		},
		{
			Nom:  "Al",
			Age:  5,
			Jour: "2juin2018",
		},
		{
			Nom:  "Az",
			Age:  15,
			Jour: "2Aout2018",
		},
		{
			Nom:  "Aw",
			Age:  9,
			Jour: "2Avril2019",
		},
		{
			Nom:  "Ay",
			Age:  4,
			Jour: "2juin2011",
		},
		{
			Nom:  "Av",
			Age:  9,
			Jour: "2janvier2013",
		},
		{
			Nom:  "Ax",
			Age:  25,
			Jour: "2septembre2008",
		},
	}

	//dataGathered := []datax{data1, data2, data3, data4}

	//fmt.Printf("Here is the total data %v\n", dataGathered)

	for _, singleData := range dataAll {

		fmt.Print(maitresse(20, singleData))

	}

}
