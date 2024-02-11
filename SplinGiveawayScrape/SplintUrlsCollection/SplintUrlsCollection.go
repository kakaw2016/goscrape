package SplintUrlsCollection

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func noDuplicateArray(contents []string) []string {
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

func scrapesource(url string) map[string][]string {

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	sourcescrape := doc.Find("#content").Text()
	//fmt.Println("We have something Urlscollection", sourcescrape)

	//reg := regexp.MustCompile("^.+Sig.+\n\n.+\n")
	//sourcescrape = reg.ReplaceAllLiteralString(sourcescrape, "")

	reg := regexp.MustCompile("..communit.+result......")

	sourcescrape1 := reg.FindAllString(sourcescrape, 1)

	//fmt.Println("We have something", sourcescrape)

	var groupdata brutedata

	err2 := json.Unmarshal([]byte(sourcescrape1[0]), &groupdata)

	if err2 != nil {
		fmt.Println(err)
	}

	formatedata := jsonToMap(groupdata)

	//fmt.Println("We have something", formatedata)

	return formatedata

}

type product struct {
	Url string
}

func (stock *product) collectedata(data map[string][]string) []string {
	var groupUrls []string

	for key, value := range data {

		if key == "url" && len(data["url"]) != 0 {

			value = data["url"]

			reg := regexp.MustCompile("(ipfs+)|(https.+images.ecency.com.+)")
			//reg1 := regexp.MustCompile("https.+")
			for _, authValue := range value {

				//if reg.MatchString(authValue) || reg1.MatchString(authValue) {
				//if reg.MatchString(authValue) {

				//	continue
				//} else if !reg.MatchString(authValue) && authValue != "" {
				if !reg.MatchString(authValue) && authValue != "" {
					reg2 := regexp.MustCompile(".+@")
					v1 := reg2.ReplaceAllString(authValue, "https://hive.blog/@")

					if v1 != "" {

						stock.Url = v1

						groupUrls = append(groupUrls, stock.Url)

					}

				}
			}

		} else if len(data["url"]) == 0 {
			fmt.Println("There is an error of KEYURL")
		}

	}
	//fmt.Println("We have something GroupUrls", groupUrls)
	return noDuplicateArray(groupUrls)

}

func Initialized() {

	var collInfo product

	var blockUrls []string

	urlsPages := []string{

		"https://hive.blog/created/giveaways",
	}

	for _, urlPage := range urlsPages {
		postCodeSource := scrapesource(urlPage)
		blockUrls = collInfo.collectedata(postCodeSource)
		blockUrls = append(blockUrls, blockUrls...)

	}

	blockUrls = noDuplicateArray(blockUrls)

	fileStoredata, err := os.OpenFile("/home/kakashinaruto/go/src/github.com/kakaw2016/goscrape/SplinGiveawayScrape/SplintConnectLinkScrape.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fileStoredata.Close()

	w := bufio.NewWriter(fileStoredata)

	var dataToStore2 string

	for _, splinterArticle := range blockUrls {

		dataToStore2 = fmt.Sprintln("\n" + splinterArticle)

		_, _ = w.WriteString(dataToStore2)

	}

	w.Flush()

}

/*func CronSchedule() {
	ch := gocron.Start()

	gocron.Every(20).Minutes().Do(Initialized)

	//go test(ch)

	_, time := gocron.NextRun()
	fmt.Println("Next Schedule In: ", time)

	<-ch

}*/
