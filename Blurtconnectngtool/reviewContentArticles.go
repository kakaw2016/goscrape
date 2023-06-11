package main

import (
	"fmt"

	"github.com/claudiu/gocron"
	"github.com/kakaw2016/goscrape/Blurtconnectngtool/BlurtUrlsCollection"
	"github.com/kakaw2016/goscrape/Blurtconnectngtool/Blurtscraper"
)

func main() {

	BlurtUrlsCollection.Initialized()

	Blurtscraper.Initialized()

	fmt.Println("Work is done")
	ch := gocron.Start()

	gocron.Every(20).Minutes().Do(BlurtUrlsCollection.Initialized)

	gocron.Every(20).Minutes().Do(Blurtscraper.Initialized)

	_, time := gocron.NextRun()
	fmt.Println("Next Schedule In: ", time)

	<-ch
}
