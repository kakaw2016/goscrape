package CollectApiData

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func NoDuplicateArray(contents []string) []string {
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

func Readfile(flocation string) []string {
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
	return NoDuplicateArray(filesdata)

}

type Battle struct {
	BattleQueueID1       string `json:"battle_queue_id_1"`
	BattleQueueID2       string `json:"battle_queue_id_2"`
	ManaCap              int    `json:"mana_cap"`
	Ruleset              string `json:"ruleset"`
	Winner               string `json:"winner"`
	Player1RatingFinal   int    `json:"player_1_rating_final"`
	Player2RatingFinal   int    `json:"player_2_rating_final"`
	Player1RatingInitial int    `json:"player_1_rating_initial"`
	Player2RatingInitial int    `json:"player_2_rating_initial"`
	Player1              string `json:"player_1"`
	Player2              string `json:"player_2"`

	IsSurrender bool `json:"is_surrender"`
}

type SplinterlandApiBattle struct {
	Player  string   `json:"player"`
	Battles []Battle `json:"battles"`
}

func Scrapesource(UrlA string) (jsonStored SplinterlandApiBattle) {

	resp, err := http.Get(UrlA)

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

	sourcescrape := doc.Find("body").Text()

	err2 := json.Unmarshal([]byte(sourcescrape), &jsonStored)

	if err2 != nil {
		fmt.Println(err)
	}

	//fmt.Println("We have something at jsonStored", jsonStored)

	return jsonStored

}
