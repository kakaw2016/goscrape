package Splintscraper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"regexp"
	"strings"

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
	//fmt.Println(noduplicate(filesdata))
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

func scrapesource(url string) map[string][]string {

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

	//fmt.Println("We have something DOCFIND", sourcescrape)

	reg := regexp.MustCompile("..communit.+result......")

	//sourcescrape = reg.ReplaceAllLiteralString(sourcescrape, "")

	sourcescrape1 := reg.FindAllString(sourcescrape, 1)

	//fmt.Println("We have something", sourcescrape1)

	var groupdata brutedata

	if len(sourcescrape1) != 0 {
		err2 := json.Unmarshal([]byte(sourcescrape1[0]), &groupdata)

		if err2 != nil {
			fmt.Println("JSON unmarshal Error", err2)
		}

	}

	formatedata := jsonToMap(groupdata)

	//fmt.Println("We have something", formatedata)

	return formatedata

}

type product struct {
	Title   string
	Url     string
	PostAge []string
	Images  string
	Voters  string
	Authors string
}

func filterByTime1(events []string, t time.Time) []string {
	filtered := make([]string, 0)
	//timeFiltered, _ := time.Parse("2006-01-02T15:04:05", filtered[0])
	var timeFiltered time.Time
	for _, e := range events {
		timeFiltered, _ = time.Parse("2006-01-02T15:04:05", e)
		if timeFiltered.After(t) {
			filtered = append(filtered, e)
		}
	}
	return noduplicate(filtered)
}

func filterByTime2(events []string, t time.Time) []string {
	filtered := make([]string, 0)
	//timeFiltered, _ := time.Parse("2006-01-02T15:04:05", filtered[0])
	var timeFiltered time.Time
	for _, e := range events {
		timeFiltered, _ = time.Parse("2006-01-02T15:04:05", e)
		if timeFiltered.Before(t) {
			filtered = append(filtered, e)
		}
	}
	return noduplicate(filtered)
}

func intersection(a, b []string) []string {
	m := make(map[string]bool)
	c := []string{}
	for _, item := range a {
		m[item] = true
	}
	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return noduplicate(c)

}

func (stock *product) collectedata(data map[string][]string) {

	for key, value := range data {

		if key == "parent_author" && len(data["parent_author"]) != 0 {

			value = data["parent_author"]

			reg1 := regexp.MustCompile("((onthemountain)|(freed99)|(hiveborgminer)|(ydaiznfts)|(bokica80)|(spectrumecons)|(bechibenner)|(phortun)|(pundito)|(oadissin))")
			for _, authValue := range value {
				if reg1.MatchString(authValue) {

					stock.Authors = ""
					break

				} else if !reg1.MatchString(authValue) {

					value = data["pathname"]

					reg1 := regexp.MustCompile("(@[a-zA-Z0-9-]+)")

					v1 := reg1.FindString(value[0])

					v1 = strings.TrimPrefix(v1, "@")

					stock.Authors = strings.ToUpper(v1)

				}

			}
		}

		if key == "title" && len(data["title"]) != 0 {

			if _, ok := data["title"]; ok {
				value = data["title"]

				//reg1 := regexp.MustCompile("(.+(gondek).+|.+(abilitie).+|.+(land).+|.+(focus).+|.+(chest).+|.+(art).+|.+(league).+|.+(quest).+|.+(challenge).+|.+(mage).+|.+(report).+|.+(lolz).+|.+(powerup).+|.+(hive).+|.+(risingstar).+|.+(virtualgrowth).+|.+(star).+|.+(rising).+|.+(crop).+|.+(woo).+|.+(alive).+|.+(chess).+|.+(phortun).+|.+(woogame).+|.+(euro).+|.+(goldquizblog).+|.+(hashkings).+)")

				reg1 := regexp.MustCompile("(.*(hh guild).*|.*(terracor).*|.*(weed).*|.*(gimmi).*|.*(mage).*|.*(art).*|.*(lolz).*|.*(powerup).*|.*(psyber).*|.*(risingstar).*|.*(virtualgrowth).*|.*(star).*|.*(rising).*|.*(crop).*|.*(woo).*|.*(alive).*|.*(chess).*|.*(phortun).*|.*(woogame).*|.*(euro).*|.*(goldquizblog).*|.*(hashkings).*)")

				for _, value2 := range value {

					value2 = strings.ToLower(value2)

					if reg1.MatchString(value2) {
						stock.Title = ""
						break

						//} else if !reg1.MatchString(value2) && value2 != "Splinterlands" && value2 != "RE:" && value2 != "" {
					} else if !reg1.MatchString(value2) && value2 != "" {

						//if value2 != "Splinterlands" && value2 != "RE:" && value2 != "" {

						titleCorrection := strings.ToUpper(value2)

						stock.Title, _ = strings.CutPrefix(titleCorrection, "RE: ")

						//fmt.Println("Title gathered", stock.Title)

					}

				}

			} else {
				fmt.Println("There is an error of KEYTITLE")
			}

		}

		if (key == "updated" || key == "created") && (len(data["updated"]) != 0 || len(data["created"]) != 0) {

			valueA := data["updated"]
			valueB := data["created"]
			var valueC, valueD []string

			dayAgoTime1 := time.Now().AddDate(0, 0, -1).Format("2006-01-02T15:04:05")

			dayAgoTime, _ := time.Parse("2006-01-02T15:04:05", dayAgoTime1)

			ageFilter1 := filterByTime1(valueA, dayAgoTime)

			timeCurrent1 := time.Now().Format("2006-01-02T15:04:05")

			timeCurrent, _ := time.Parse("2006-01-02T15:04:05", timeCurrent1)

			ageFilter2 := filterByTime2(valueB, timeCurrent)

			valueC = intersection(ageFilter1, ageFilter2)

			if len(valueC) != 0 && len(valueC) >= 2 {

				valueD = append(valueD, valueC[0], valueC[1])

				stock.PostAge = valueD

			} else if len(valueC) != 0 && len(valueC) < 2 {

				valueD = append(valueD, valueC...)

				stock.PostAge = valueD
			} else {
				stock.PostAge = []string{}

			}

			//fmt.Println(stock.Title)

			if len(data["updated"]) == 0 && len(data["created"]) == 0 {
				fmt.Println("There is an error of KEYAGE")
			}
		}

		//stock.PostAge = []string{"In progress", "will come"}

		if key == "pathname" && len(data["pathname"]) != 0 {

			value = data["pathname"]

			//reg1 := regexp.MustCompile("(.+(mage).+|.+(art).+|.+(report).+|.+(risingstar).+|.+(virtualgrowth).+|.+(star).+|.+(rising).+|.+(dcrops).+|.+(crop).+|.+(hivechess).+|.+(phortun).+|.+(wrestling).+|.+(woo).+|.+(euro).+|.+(goldquizblog).+|.+(hashkings).+)")

			reg1 := regexp.MustCompile("(.+(giveaw).+|.+(delegat).+|.+(free ).+|.+(hsbi).+|.+(raffle).+|.+(kojiri).+)")

			reg2 := regexp.MustCompile("(.+@)")

			if reg1.MatchString(value[0]) {

				//v1 := strings.ReplaceAll(value[0], "/hive-13323/@", "https://ecency.com/@")
				v1 := reg2.ReplaceAllString(value[0], "https://ecency.com/@")

				stock.Url = v1

			} else {
				stock.Url = ""

			}

		} else if len(data["pathname"]) == 0 {
			fmt.Println("There is an error of KEYURL")
		}

		if key == "image" && len(data["image"]) != 0 {

			value = data["image"]

			valueC := []string{}

			for _, valueA := range value {

				//reg1 := regexp.MustCompile("(.+(risingstar).+|.+(virtualgrowth).+|.+(star).+|.+(rising).+)")

				reg2 := regexp.MustCompile("https.+(i.imgur|DQmXJux2kvzMb5n2ooQKy7f3rzx9fcsHPsUddWaGjXehL3f/giveaway5k|DQmaJ7PZXJZJu2vJ3eGrXQ4tQLamvUnLHhqjkkbVEdEDpzb/902c291e9d0c9f48fa80e85523c6401de12caf65efc761c2dad4adefa4bae1eb|23w2cNG3rWNgyzkbBCCKxZMuaPHhadpoJHXqjB7ZitTK6ZcjMzAad6wSXs9BSgu2YWgU2|(lolztoken.+lolz)).+(webp|jpg|jpeg|png|gif|JPG|JPEG|PNG|GIF)")

				if !reg2.MatchString(valueA) {
					reg3 := regexp.MustCompile("https.+" + "((webp)|(jpg)|(jpeg)|(png)|(gif)|(JPG)|(JPEG)|(PNG)|(GIF))")

					valueB := reg3.FindString(valueA)

					if valueB != "" {
						valueC = append(valueC, valueB)

					}
				}

			}
			valueC = noduplicate(valueC)
			if len(valueC) != 0 {

				//fmt.Println("Result ValueC", valueC)
				stock.Images = valueC[0]

			}
		}

		if key == "total_votes" && len(data["total_votes"]) != 0 {

			value = data["total_votes"]

			for _, rightValue := range value {

				if rightValue != "0" {

					stock.Voters = rightValue

				}

			}

		} /* else if len(data["total_votes"]) == 0 {
			fmt.Printf("There is an error of KEYVOTERS on the %s\n", stock.Url)
		}*/

	}

}

func Initialized() {

	var blurtUrls []string

	blurtUrls1 := readfile("/home/kakashinaruto/go/src/github.com/kakaw2016/goscrape/SplinGiveawayScrape/SplintConnectLinkScrape.txt")

	blurtUrls2 := readfile("/home/kakashinaruto/go/src/github.com/kakaw2016/goscrape/SplinGiveawayScrape/SplinterPosts")

	blurtUrls = append(blurtUrls, blurtUrls1...)
	blurtUrls = append(blurtUrls, blurtUrls2...)
	blurtUrls = noduplicate(blurtUrls)

	fileStoredata, err := os.OpenFile("/home/kakashinaruto/go/src/github.com/kakaw2016/goscrape/SplinGiveawayScrape/SplintLiveScrape.txt", os.O_CREATE|os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fileStoredata.Close()

	fmt.Println("Total URL", len(blurtUrls))

	const headtext string = `
<div class="text-justify">

**Dear friends,**

**Peace,**

</div>

<div class="text-justify">

Beautiful day over the colorful landscape of Pretoria.

How is your adventure progression in the universe of Splinterlands?

I hope that your season reward is great and will remain diverse and rich in powerful cards.

My adventure is exceptional in the Wild format. The progression could be better but it is all about the journey and less about the destination.

Enjoy your campaign and achieve your goals. The season is full of surprises great Warriors on Pretoria Lands.


<div class="text-center">

[![freeimageSource](https://images.hive.blog/500x0/https://files.peakd.com/file/peakd-hive/oadissin/23wgRuXqeKYgchLaNsWcg2CbB8NapA5mTwVVKSfHazXn2xycK2ShKvxvSnBhtkoULU6nv.png)](https://peakd.com/@oadissin/posts)

</div>

--

The Splinterlands directory is a compilation of articles searched and filtered through hundreds of contests, delegations, and giveaway-related publication on Hive.

I can assure you that it can be difficult to navigate the chaotic content feed to target one specific category of publication. Also, it is time-consuming to find an active contest each day.

Splinterlands players have the opportunity to read a brief report on a recent contest posted on Hive. I keep the directory updated so that community members can easily participate in various worthwhile Splinterlands activities hosted on the blockchain.



--

<div class="text-justify">

So I thought I could post and maintain here an updated directory of several contests, giveaways, or raffles related to Splinterlands.

The community members are welcome to participate in updating this list. By doing so this project can achieve its community goals.

*I have been thinking about a "tag" that can be used to easily identify items similar to those in the list below: "#splinterlandsgiveaways"

**I invite authors to suggest a better tag or use #splinterlandsgiveaways in the list of tags before submitting the article


<div class="text-center">

![](https://images.ecency.com/DQmeUyxbYf52V7k5Kw8gS6NJdc5yFh34x3YvifyP97FxzXD/sp3.gif)

--

</div>

<div class="text-center">

Here we go...

</div>

<div class="text-center">

* - * - * - *
</div>

<div class="text-justify">
`
	const endtext string = `
</div>

<div class="text-center">

- * - * - * - * -

</div>

<div class="text-center">


![](https://images.ecency.com/DQmNr7UUfiRZPVRkGpPDAhwLku3inZ23UVfGKkyG1AzY5Me/18.gif)


</div>

<div class="text-justify"> 

The main page of the author's account may present up-to-date contests that were not mentioned in this publication.


Please check your favorite contest author’s profile to place your entry as early as possible.
</div>


<div class="text-center">


[![freeimageSource](https://images.hive.blog/500x0/https://images.ecency.com/DQmUfC6jgWDuLBcFAQiDAWkPPqAPJuuQhvHZh8iN7yeS8eo/splinterlandsdirectorycontest.png)](https://peakd.com/@oadissin/posts)


</div>

Feel free to read more about the contest I share on my blog. Please visit the following blog page.


<div class="text-center"> 

++[Blog Profile](https://ecency.com/@oadissin/posts)++

</div>


<div class="text-center"> 


## Warm regards
</div>

Delegation - Raffles - Giveaways: 40+ Updated Contests // Directory #️⃣ 200

splinterlands spt play2earn indiaunited ocdb freecompliments thgaming play2own oneup neoxian leo alive pgm pimp proofofbrain meme waiv contest hive-engine archon cent vyb ctp bro hhguild


Description

Splinterlands warriors peace and prosperity. Here are some giveaways and contest links gathered from Splinterlands articles.
`

	w := bufio.NewWriter(fileStoredata)

	dataToStore1 := fmt.Sprintf("\n%s\n", headtext)

	_, _ = w.WriteString(dataToStore1)
	w.Flush()

	var collinfo product

	for _, blurtPost := range blurtUrls {

		postCodeSource := scrapesource(blurtPost)
		collinfo.collectedata(postCodeSource)

		if collinfo.Title != "" && collinfo.Images != "" && collinfo.Url != "" && collinfo.Voters != "" && collinfo.Authors != "" && len(collinfo.PostAge) != 0 {

			dataToStore2 := fmt.Sprintf("\n☎☎☎\n%s\n\n<div class=\"text-center\">\n\n[![](https://images.hive.blog/450x0/%s)](%s)\n</div>\nPosted Since: %s\n\nVoted By %s Hive Bloggers\nArticle Creator: %s\n", collinfo.Title, collinfo.Images, collinfo.Url, collinfo.PostAge[:], collinfo.Voters, collinfo.Authors)

			//if (collinfo.Images != "" && collinfo.Url != "" && collinfo.Voters != "" && collinfo.Authors != "") || (collinfo.Title != "") {
			//dataToStore2 := fmt.Sprintf("\nN. %d\n%s\n\n<div class=\"text-center\">\n\n [![](%s)](%s)\n</div>\n\nVoted By %s Blurtians\nArticle Creator: %s\n", i+1, collinfo.Title, collinfo.Images, collinfo.Url, collinfo.Voters, collinfo.Authors)
			_, _ = w.WriteString(dataToStore2)
			w.Flush()
		}
	}

	dataToStore3 := fmt.Sprintf("\n%s\n", endtext)

	_, _ = w.WriteString(dataToStore3)
	w.Flush()

	/*for _, blurtPost := range blurtUrls[:1] {
		fmt.Println(blurtPost)
		z := scrapesource(blurtPost)
		for k1, z1 := range z {
			fmt.Println(k1, "======>", z1)

		}

	}*/

}

/*func Cronjob() {

	ch := gocron.Start()

	gocron.Every(30).Minutes().Do(Initialized)

	//go test(ch)

	_, time := gocron.NextRun()
	fmt.Println("Next Schedule In: ", time)

	<-ch

}*/

/*
Filtration

Date to KEEP

\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*2022\-12\-[0][1-2].*2022\-12\-[0][1-2].*\n.*\n.*\n.*
*/
