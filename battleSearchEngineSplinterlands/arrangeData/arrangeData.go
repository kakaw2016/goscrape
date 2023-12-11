package ArrangeData

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	CollectApiData "github.com/kakaw2016/goscrape/battleSearchEngineSplinterlands/collectApiData"
)

type product struct {
	Players    string
	BattleLink string
	Rulesets   struct {
		Rules string
		Mana  int
	}
	PlayerWinNumb          string
	Warrior1RankingInitial int
	Warrior2RankingInitial int
	Warrior1RankingFinal   int
	Warrior2RankingFinal   int
}

var products []product

func (stock *product) categorizeData(data CollectApiData.SplinterlandApiBattle) []product {

	for _, value := range data.Battles {

		//fmt.Println(value)

		if value.Winner == data.Player && !value.IsSurrender && (value.Player1RatingInitial < value.Player1RatingFinal) {

			stock.BattleLink = "https://splinterlands.com/?p=battle&id=" + value.BattleQueueID1

			stock.Rulesets = struct {
				Rules string
				Mana  int
			}{value.Ruleset, value.ManaCap}

			stock.Warrior1RankingFinal = value.Player1RatingFinal

			//stock.Warrior2RankingFinal = value.Player2RatingFinal

			stock.Warrior1RankingInitial = value.Player1RatingInitial

			//stock.Warrior2RankingInitial = value.Player2RatingInitial

			stock.Players = value.Player1

			stock.PlayerWinNumb = "Player1"

			newdata := product{

				Players:       stock.Players,
				BattleLink:    stock.BattleLink,
				Rulesets:      stock.Rulesets,
				PlayerWinNumb: stock.PlayerWinNumb,

				Warrior1RankingFinal: stock.Warrior1RankingFinal,
				Warrior2RankingFinal: stock.Warrior2RankingFinal,
			}

			products = append(products, newdata)

		}

		if value.Winner == data.Player && !value.IsSurrender && (value.Player2RatingInitial < value.Player2RatingFinal) {

			stock.BattleLink = "https://splinterlands.com/?p=battle&id=" + value.BattleQueueID1

			stock.Rulesets = struct {
				Rules string
				Mana  int
			}{value.Ruleset, value.ManaCap}

			//stock.Warrior1RankingFinal = value.Player1RatingFinal

			stock.Warrior2RankingFinal = value.Player2RatingFinal

			//stock.Warrior1RankingInitial = value.Player1RatingInitial

			stock.Warrior2RankingInitial = value.Player2RatingInitial

			stock.Players = value.Player2

			stock.PlayerWinNumb = "Player2"

			newdata := product{

				Players:       stock.Players,
				BattleLink:    stock.BattleLink,
				Rulesets:      stock.Rulesets,
				PlayerWinNumb: stock.PlayerWinNumb,

				Warrior1RankingFinal: stock.Warrior1RankingFinal,
				Warrior2RankingFinal: stock.Warrior2RankingFinal,
			}

			products = append(products, newdata)

		}

	}

	return products

}

var battleDetailed product

func CodeLauncher() {

	currentDir, _ := os.Getwd()
	configPath := filepath.Join(currentDir, "configurationfile.txt")
	fileStoredata := filepath.Join(currentDir, "outputdata.txt")

	fileData, err := os.OpenFile(fileStoredata, os.O_CREATE|os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fileData.Close()

	w := bufio.NewWriter(fileData)

	collectedUrls := CollectApiData.Readfile(configPath)

	for _, playerData := range collectedUrls {

		battleApi := CollectApiData.Scrapesource(playerData)

		//fmt.Println("we have battleapi", battleApi)

		finalProductA := battleDetailed.categorizeData(battleApi)

		//fmt.Println("we have FinalProductA", finalProductA)

		for _, finalProductB := range finalProductA {

			fileData := fmt.Sprintf("\n%v\n-----------\n", finalProductB)

			_, _ = w.WriteString(fileData)

			w.Flush()

		}

	}

}
