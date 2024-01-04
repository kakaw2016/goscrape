package main

import (
	"bufio"
	"encoding/json"
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

type brutedata map[string]interface{}

func jsonToMap(data map[string]interface{}) map[string][]string {
	// final output
	out := make(map[string][]string)

	// check all keys in data
	for key, value := range data {
		// check if key not exist in out variable, add it
		if _, ok := out[key]; !ok {
			out[key] = []string{}
		}

		if valueA, ok := value.(map[string]interface{}); ok { // if value is map
			out[key] = append(out[key], "")
			for keyB, valueB := range jsonToMap(valueA) {
				if _, ok := out[keyB]; !ok {
					out[keyB] = []string{}
				}
				out[keyB] = append(out[keyB], valueB...)
			}
		} else if valueA, ok := value.([]interface{}); ok { // if value is array
			for _, valueB := range valueA {
				if valueC, ok := valueB.(map[string]interface{}); ok {
					for keyD, valueD := range jsonToMap(valueC) {
						if _, ok := out[keyD]; !ok {
							out[keyD] = []string{}
						}
						out[keyD] = append(out[keyD], valueD...)
					}
				} else {
					out[key] = append(out[key], fmt.Sprintf("%v", valueB))
				}
			}
		} else { // if string and numbers and other ...
			out[key] = append(out[key], fmt.Sprintf("%v", value))
		}

	}

	return out
}

type product struct {
	Title   string
	Voters  string
	Authors string
}

func scrapesource(url string) (map[string]string, map[string][]string) {

	formatedata2 := make(map[string][]string)

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
	var sourcescrape1 string
	//fmt.Println("We have something map", sourcescrape)

	reg1 := regexp.MustCompile(`[0-9]{1,}[.,]?\d+[^a-z]+SAVINGS`)
	sourcescrape1 = reg1.FindString(sourcescrape)
	//fmt.Println("We have something sourcescrape4", sourcescrape4)

	reg2 := regexp.MustCompile(`@[a-z09.-]+`)

	accounts := reg2.FindString(url)

	formatedata1 := make(map[string]string)

	if sourcescrape1 != "" {

		formatedata1[accounts] = sourcescrape1
	}
	//fmt.Println("We have something", formatedata1)

	sourcescrape2 := strings.Split(sourcescrape, "\n")

	sourcescrape3 := ""

	for _, rightData := range sourcescrape2 {

		if !strings.Contains(rightData, "global") {

			continue
		} else {

			sourcescrape3 = rightData
		}

		//fmt.Println(sourcescrape3)
	}

	var groupdata brutedata

	err2 := json.Unmarshal([]byte(sourcescrape3), &groupdata)

	if err2 != nil {
		fmt.Println(err)
	}

	formatedata2 = jsonToMap(groupdata)

	//fmt.Println("We have something", formatedata2)

	return formatedata1, formatedata2

}

func (stock *product) collectedata(data map[string][]string) {

	for key, value := range data {

		/*if key == "name" && len(data["name"]) != 0 {

			value = data["name"]

			//reg1 := regexp.MustCompile("(blurtconnect-ng)|(alejos7ven)|(onchain-curator)|(clixmoney)|(tekraze)|(saboin)|(joviansummer)|(lucylin)|(phusionphil)")
			for _, authValue := range value {
				//if !reg1.MatchString(authValue) && authValue != "post" {
				if authValue != "post" {

					stock.Authors = "@" + authValue

					//} else if reg1.MatchString(authValue) || authValue == "post" {
				} else if authValue == "post" {

					stock.Authors = ""
				}

			}

		} else if len(data["name"]) == 0 {
			fmt.Println("There is an error of KEYAUTHORS")
		}*/

		if key == "witnesses_voted_for" && len(data["witnesses_voted_for"]) != 0 {

			value = data["witnesses_voted_for"]

			stock.Voters = value[0]

		} else if len(data["witnesses_voted_for"]) == 0 {
			fmt.Println("There is an error or No vote for witnesses")
		}

	}

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

	var dataToStore2, dataToStore3 string
	var collinfo product

	for _, blurtPost := range ecosynBlurtUrls {
		//fmt.Println(blurtPost)

		y, z := scrapesource(blurtPost)

		collinfo.collectedata(z)
		for k1, z1 := range y {
			//fmt.Println("Key in the Map", k1, "======>", "VALUE", z1)
			dataToStore3 = fmt.Sprint(k1, "\t", z1)
			_, _ = w.WriteString(dataToStore3)

		}
		w.Flush()
		//if collinfo.Title != "" && collinfo.Voters != "" && collinfo.Authors != "" {
		if collinfo.Voters != "" {

			dataToStore2 = fmt.Sprintf("\t Voted %s Witnesses\n", collinfo.Voters)
			_, _ = w.WriteString(dataToStore2)
			w.Flush()
		}

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
