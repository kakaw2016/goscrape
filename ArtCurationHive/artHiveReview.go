package main

import (
	"fmt"

	"github.com/claudiu/gocron"
	"github.com/kakaw2016/goscrape/ArtCurationHive/Artscraper"
	"github.com/kakaw2016/goscrape/ArtCurationHive/Arturlscollection"
)

func main() {

	Arturlscollection.Initialized()

	Artscraper.Initialized()

	fmt.Println("Work is done")
	ch := gocron.Start()

	gocron.Every(2).Days().Do(Arturlscollection.Initialized)

	gocron.Every(2).Days().Do(Artscraper.Initialized)

	_, time := gocron.NextRun()
	fmt.Println("Next Schedule In: ", time)

	<-ch
}

//.*\n.*THGAMING.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n
