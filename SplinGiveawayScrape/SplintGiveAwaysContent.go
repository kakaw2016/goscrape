package main

import (
	"fmt"

	"github.com/claudiu/gocron"
	"github.com/kakaw2016/goscrape/SplinGiveawayScrape/SplintUrlsCollection"
	"github.com/kakaw2016/goscrape/SplinGiveawayScrape/Splintscraper"
)

func main() {

	SplintUrlsCollection.Initialized()

	Splintscraper.Initialized()

	fmt.Println("Work is done")
	ch := gocron.Start()

	gocron.Every(20).Minutes().Do(SplintUrlsCollection.Initialized)

	gocron.Every(20).Minutes().Do(Splintscraper.Initialized)

	_, time := gocron.NextRun()
	fmt.Println("Next Schedule In: ", time)

	<-ch
}

//.*\n.*THGAMING.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n
