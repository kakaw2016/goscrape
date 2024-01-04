package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/claudiu/gocron"

	"github.com/PuerkitoBio/goquery"
)

func noduplicate(contents []string) []string {
	contentdata := make(map[string]bool)
	correctlist := []string{}
	for _, data := range contents {
		if _, value := contentdata[data]; !value {
			contentdata[data] = true
			correctlist = append(correctlist, data)
		}
	}
	return correctlist
}

func readfile(flocation string) []string {
	var filesdata []string

	blurtlinks, error := os.Open(flocation)
	if error != nil {
		log.Fatal(error)
	}
	defer blurtlinks.Close()
	scanner := bufio.NewScanner(blurtlinks)

	for scanner.Scan() {
		filedata := fmt.Sprintln(scanner.Text())

		filedata = strings.Trim(filedata, "\n")

		if filedata != "" {
			filesdata = append(filesdata, filedata)

		}

	}
	error = scanner.Err()
	if error != nil {
		log.Fatal(error)
	}
	//fmt.Println("test length", len(noduplicate(filesdata)))
	return noduplicate(filesdata)

}

func scrapesource(url string) map[string]string {

	formatedata := make(map[string]string)

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("GET URLS", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s %v", resp.StatusCode, resp.Status, resp.Request.URL)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	sourcescrape := doc.Find("#content").Text()
	var sourcescrape4 string
	//fmt.Println("We have something map", sourcescrape)

	reg := regexp.MustCompile(`[0-9]{1,}[.,]?\d+[^a-z]+SAVINGS`)
	sourcescrape4 = reg.FindString(sourcescrape)
	//fmt.Println("We have something sourcescrape4", sourcescrape4)

	reg2 := regexp.MustCompile(`@[a-z09.-]+`)

	accounts := reg2.FindString(url)

	if sourcescrape4 != "" {

		formatedata[accounts] = sourcescrape4
	}

	//fmt.Println("We have something map", formatedata)

	return formatedata

}

func initialized() {

	currentDir, _ := os.Getwd()
	configPath := filepath.Join(currentDir, "EcoSynsUrlFile.txt")
	relativePath, _ := filepath.Rel(currentDir, configPath)
	fmt.Println("Reading config file at relative path:", relativePath)

	ecosynBlurtUrls := readfile(relativePath)

	fileStoredata, err := os.OpenFile("BlurtPowerDataScrape.txt", os.O_CREATE|os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		log.Fatal("Storage of data encounter some malfunctioning", err)
	}
	defer fileStoredata.Close()

	fmt.Println("Total URL", len(ecosynBlurtUrls))

	w := bufio.NewWriter(fileStoredata)

	var dataToStore3 string

	for _, blurtPost := range ecosynBlurtUrls {
		//fmt.Println(blurtPost)

		z := scrapesource(blurtPost)
		for k1, z1 := range z {
			//fmt.Println(k1, "======>", z1)
			dataToStore3 = fmt.Sprintln(k1, "==>", z1)
			_, _ = w.WriteString(dataToStore3)

		}
		w.Flush()

	}

}

/*func test(stop chan bool) {
	time.Sleep(20 * time.Second)
	gocron.Clear()
	fmt.Println("All task removed")
	close(stop)
}*/

func cronjob() {

	ch := gocron.Start()

	gocron.Every(30).Minutes().Do(initialized)

	//go test(ch)

	_, time := gocron.NextRun()
	fmt.Println("Next Schedule In: ", time)

	<-ch

}

func main() {

	initialized()
	cronjob()
}
