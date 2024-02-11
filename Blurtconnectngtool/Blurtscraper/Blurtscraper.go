package Blurtscraper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"regexp"
	"strings"
	"time"

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

func scrapesource(url string) map[string][]string {

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("GET URLS", err)
	}

	defer resp.Body.Close()

	/*if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s %v", resp.StatusCode, resp.Status, resp.Request.URL)

	}*/

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	sourcescrape := doc.Find("#content").Text()

	reg := regexp.MustCompile("^.+Sig.+LoginSign up")

	sourcescrape = reg.ReplaceAllLiteralString(sourcescrape, "")

	//fmt.Println("We have something", sourcescrape)

	var groupdata brutedata

	_ = json.Unmarshal([]byte(sourcescrape), &groupdata)

	/*if err2 != nil {
		fmt.Println("JSON Unmarshal", err2)
	}*/

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

		if key == "name" && len(data["name"]) != 0 {

			value = data["name"]

			reg1 := regexp.MustCompile("(blurtofficial)|(andgon99)|(blurtconnect-ng)|(alejos7ven)|(onchain-curator)|(clixmoney)|(tekraze)|(saboin)|(joviansummer)|(lucylin)|(phusionphil)")
			for _, authValue := range value {
				if !reg1.MatchString(authValue) && authValue != "post" {
					//if authValue != "post" {

					stock.Authors = "@" + authValue

				} else if reg1.MatchString(authValue) || authValue == "post" {
					//} else if authValue == "post" {

					stock.Authors = ""
				}

			}

		} else if len(data["name"]) == 0 {
			fmt.Println("There is an error of KEYAUTHORS")
		}

		/*if key == "root_title" && len(data["root_title"]) != 0 {

			if _, ok := data["root_title"]; ok {
				value = data["root_title"]
				reg1 := regexp.MustCompile("(app available on Google Play.+)")

				if !reg1.MatchString(value[0]) && value[0] != "" {

					stock.Title = strings.ToUpper(value[0])

					//fmt.Println("Title gathered", stock.Title)

				} else if reg1.MatchString(value[0]) || (value[0] == "") {
					stock.Title = ""
					//fmt.Println("Title gathered Empty", stock.Title)

				}

			} else {
				fmt.Println("There is an error of KEYTITLE")
			}

		}*/

		if key == "pathname" && len(data["pathname"]) != 0 {

			value = data["pathname"]

			regCorrection := regexp.MustCompile("^.+/")

			v1 := regCorrection.ReplaceAllLiteralString(value[0], "")

			regCorrection = regexp.MustCompile("-")

			v2 := regCorrection.ReplaceAllLiteralString(v1, " ")

			stock.Title = strings.ToUpper(v2)

		} else if len(data["pathname"]) == 0 {
			fmt.Println("There is an error of TITLE")
		}

		if (key == "time" || key == "last_update") && (len(data["time"]) != 0 || len(data["last_update"]) != 0) {

			valueA := data["last_update"]
			valueB := data["time"]
			var valueC, valueD []string

			dayAgoTime1 := time.Now().AddDate(0, 0, -3).Format("2006-01-02T15:04:05")

			dayAgoTime, _ := time.Parse("2006-01-02T15:04:05", dayAgoTime1)

			ageFilter1 := filterByTime1(valueA, dayAgoTime)

			timeCurrent1 := time.Now().Format("2006-01-02T15:04:05")

			timeCurrent, _ := time.Parse("2006-01-02T15:04:05", timeCurrent1)

			ageFilter2 := filterByTime2(valueB, timeCurrent)

			valueC = intersection(ageFilter1, ageFilter2)

			//fmt.Println("testing element", valueA)

			if len(valueC) != 0 && len(valueC) >= 2 {

				valueD = append(valueD, valueC[0], valueC[1])

				stock.PostAge = valueD

			} else if len(valueC) != 0 && len(valueC) < 2 {

				valueD = append(valueD, valueC...)

				stock.PostAge = valueD
			} else {
				stock.PostAge = []string{}

			}
			//fmt.Println("testing element", stock.PostAge)

			if len(data["last_update"]) == 0 && len(data["time"]) == 0 {
				fmt.Println("There is an error of KEYAGE")
			}

		}
		if key == "pathname" && len(data["pathname"]) != 0 {

			value = data["pathname"]
			//v1 := strings.ReplaceAll(value[0], "/@", "https://blurtlatam.intinte.org/@")
			regCorrection := regexp.MustCompile("^.+@")

			v1 := regCorrection.ReplaceAllLiteralString(value[0], "https://blurt.blog/@")

			stock.Url = v1

		} else if len(data["pathname"]) == 0 {
			fmt.Println("There is an error of KEYURL")
		}

		if key == "body" && len(data["body"]) != 0 {

			value = data["body"]

			valueC := []string{}

			for _, valueA := range value {

				reg2 := regexp.MustCompile("https.+(i.imgur|51f67e7fe072b0ad0fb02f079493b62ad3965f04|fb1a8a788360e7f39bd770b6ecfbe60f1364285b|pastormike/162aa91c5cb0e5ef78f0ad07e388f0d2fe53d87|blurtlatam.+d9667a3dcb3a4323|nalexadre|alejos7ven|onchain-curator|andgon99|symbion|dianaventas|bichotaclan|clixmoney|tekraze|saboin|joviansummer|lucylin|phusionphil).+(webp|jpg|jpeg|png|gif|JPG|JPEG|PNG|GIF)")

				if !reg2.MatchString(valueA) {
					reg3 := regexp.MustCompile("https.+" + "((webp)|(jpg)|(jpeg)|(png)|(gif)|(JPG)|(JPEG)|(PNG)|(GIF))")

					val1 := strings.Split(valueA, "\"")
					for _, val2 := range val1 {
						valueB := reg3.FindString(val2)

						if valueB != "" {
							valueC = append(valueC, valueB)

						}

					}

				}

			}
			valueC = noduplicate(valueC)
			if len(valueC) != 0 {

				//fmt.Println("Result ValueC", valueC)
				stock.Images = valueC[0]

			}

		} else if key == "json_metadata" && len(data["json_metadata"]) != 0 {

			ytvalue := data["json_metadata"]

			var ytvalue2, ytvalue5 []string
			reg4 := regexp.MustCompile("https.+yout.+(webp|jpg|jpeg|png|gif|JPG|JPEG|PNG|GIF)")
			reg5 := regexp.MustCompile("https.+S2ZQ3XA2OBM.+(webp|jpg|jpeg|png|gif|JPG|JPEG|PNG|GIF)")

			for _, ytvalue1 := range ytvalue {

				ytvalue2 = strings.Split(ytvalue1, "\",\"")

				for _, ytvalue3 := range ytvalue2 {

					if !reg5.MatchString(ytvalue3) {

						ytv1 := strings.Split(ytvalue3, "\",\"")

						for _, ytv2 := range ytv1 {

							ytvalue4 := reg4.FindString(ytv2)

							if ytvalue4 != "" {
								ytvalue5 = append(ytvalue5, ytvalue4)

							}

						}

					}
				}

			}
			ytvalue5 = noduplicate(ytvalue5)
			if len(ytvalue5) != 0 {

				stock.Images = ytvalue5[0]

				//fmt.Println("The result are", stock.Images)

			}

		} else if len(data["body"]) == 0 {
			fmt.Println("There is an error of KEYIMAGE")
		}

		if key == "total_votes" && len(data["total_votes"]) != 0 {

			value = data["total_votes"]

			stock.Voters = value[0]

		} else if len(data["total_votes"]) == 0 {
			fmt.Println("There is an error of KEYVOTERS")
		}

	}

}

func Initialized() {

	blurtUrls := readfile("/home/kakashinaruto/go/src/github.com/kakaw2016/goscrape/Blurtconnectngtool/BlurtConnectLinkScrape.txt")

	fileStoredata, err := os.OpenFile("/home/kakashinaruto/go/src/github.com/kakaw2016/goscrape/Blurtconnectngtool/BlurtLiveScrape.txt", os.O_CREATE|os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fileStoredata.Close()

	/*fmt.Println("Total URL", len(blurtUrls))

	for _, blurtPost := range blurtUrls[1:3] {
		fmt.Println(blurtPost)
		z := scrapesource(blurtPost)
		for k1, z1 := range z {
			fmt.Println(k1, "======>", z1)

		}

	}*/

	const headtext string = `
<div class="text-justify">

Peace,

Greetings,

I hope you are doing great my friends.

The curation effort is still on for the pleasure of all Blurtians in the community.

Blurtconnect-ng teams have done successful collaborative works to provide attention to amazing Blurters articles published on a daily basis in the community.

Blurtconnect-ng teams are watching out for any form of plagiarism that may occur after reading your publications. It is important that we cite and properly reference our written words before posting on Blurt. Since it is more difficult to recover from the damage of a confirmed plagiarized article than to respect the rules regarding content publication on Blurt, I will urge every Blurtians to avoid Plagiarism.

Below are a few posts selected for this scope of publications. The articles in the Blurtconnect community are highlighted.

</div>

<div class="text-center">

![i](https://imgp.blurt.world/700x0/https://img.blurt.world/blurtimage/blurtconnect-ng/5277734eccb9ab52b2cf7884feabdadd27296880.png)

</div>

<div class="text-center">

#-#-#-# 
</div>

<div class="text-center">

## Blurtconnect would like to present a 75% upvote opportunity to all Blurtians that leave "3 thoughtful" comments and 3 upvotes (at a friendly percentage) to any of your self-selected articles in the list below. Please share a message in the comment section of this article to announce the completion of your curation effort in the community for potential upvotes from blurtconnect-ng.
</div>

<div class="text-justify">

|Information about the creation of Blurt account|Example of the first page on the portal|
|-|-|
|You can easily create an account for your partner, family members or friends go through the account creation steps on the app "LestBlurt". If you have not installed the app on your smartphone you can download it at the following [Address Link](https://rb.gy/putpdt)the APK version for Android. | ![W](https://imgp.blurt.world/200x0/https://img.blurt.world/blurtimage/thankyouforevery/8e54806cea855a29ff96e2ef977453697144af99.jpeg)|
|In case you have alternative apps or websites offering the registration of a new account please write in the comment your information|

</div>

<div class="text-center">

#-#-#-#
</div>

<div class="text-justify">
`
	const endtext string = `</div>

<div class="text-center">

-- 
</div>

### The motivation of the day: Overcoming Obstacles and Achieving Success

	
<div class="text-center">


# Follow-Up News

</div>

<div class="text-center">

![54ae8e85fb890b0a71db8fce52a4e4b57f3077e5.png](https://imgp.blurt.world/700x0/https://img.blurt.world/blurtimage/blurtconnect-ng/6c83395b21b672a070991832ac0e0d515afc47a7.png)

<br>

<div class="text-center">

https://www.youtube.com/watch?v=olodT2jiNHA

</div>


<div class="text-justify">

Blurtconnect-ng Team Is Running A Witness Node On BLURT.

* https://blurtwallet.com/~witnesses?highlight=blurtconnect-ng

Please kindly click on this link above to Vote Our Witness.


![3505e605d235eea847207bbfd9e475917a54262f.jpg](https://imgp.blurt.world/700x0/https://img.blurt.world/blurtimage/blurtconnect-ng/f4f151738d56b858452e9affcb8c4dee7a0b426f.jpg)


<div class="text-center">

# BLURTCONNECT-NG MOTTO
</div>


### The strength of the wolves is the pack and the strength of the pack is the wolves.

<div class="text-center">

### **All for one and one for all.**
#### *BLURT belongs to all of us.*
##### Let us all join hands and give blurt more value.

</div>

--

<div class="text-center">

### CLICK [_HERE_](https://blurt.blog/introduceyourself/@blurtconnect-ng/we-are-glad-to-finally-announce-our-community-blurtconnect-here-is-our-introduction-post) TO VIEW BLURTCONNECT INTRODUCTION POST
</div>


<div class="text-center">

![libertyblockchainblurt.png](https://imgp.blurt.world/300x0/https://img.blurt.world/blurtimage/blurtconnect-ng/c7323ffd2eb749a3a4114de04d2bd622f1a8b1fd.png)
[src](https://giphy.com/channel/Blurtblog)

</div>

<div class="text-center">

# ***STAY TUNED FOR MORE***
	
</div>

BLURTCONNECT-NG LARGE SCOPE CONTENTS REPORT N# // 2% to Null

Blurtconnect Witness and Curators welcome you all to this brief recapitulation of posts in the community

blurtfirst blurtafrica instablurt r2cornell blurtlatam blurtpak newvisionlife curation blurtindia blurthispano
`

	w := bufio.NewWriter(fileStoredata)

	dataToStore1 := fmt.Sprintf("\n%s\n", headtext)

	_, _ = w.WriteString(dataToStore1)
	w.Flush()

	var collinfo product

	for _, blurtPost := range noduplicate(blurtUrls) {

		postCodeSource := scrapesource(blurtPost)
		collinfo.collectedata(postCodeSource)

		//if collinfo.Title != "" && collinfo.Images != "" && collinfo.Url != "" && len(collinfo.PostAge) == 2 && collinfo.Voters != "" && collinfo.Authors != "" {
		if collinfo.Title != "" && collinfo.Images != "" && collinfo.Url != "" && collinfo.Voters != "" && collinfo.Authors != "" && len(collinfo.PostAge) != 0 {
			dataToStore2 := fmt.Sprintf("\n☀ ☎ ☀ ☂ ☀ ☂ ☎\n%s\n\n<div class=\"text-center\">\n\n[![](https://imgp.blurt.world/550x0/%s)](%s)\n</div>\nPosted Since: %s\n\nVoted By %s Blurtians\nArticle Creator: %s\n", collinfo.Title, collinfo.Images, collinfo.Url, collinfo.PostAge[:], collinfo.Voters, collinfo.Authors)
			_, _ = w.WriteString(dataToStore2)
			w.Flush()
		}
	}

	dataToStore3 := fmt.Sprintf("\n%s\n", endtext)

	_, _ = w.WriteString(dataToStore3)
	w.Flush()

}

/*func test(stop chan bool) {
	time.Sleep(20 * time.Second)
	gocron.Clear()
	fmt.Println("All task removed")
	close(stop)
}

func Cronjob() {

	ch := gocron.Start()

	gocron.Every(30).Minutes().Do(Initialized)

	//go test(ch)

	_, time := gocron.NextRun()
	fmt.Println("Next Schedule In: ", time)

	<-ch

}*/

/*func main() {
	initialized()
	cronjob()

}*/

/*

Filtration

Date to KEEP


\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*2022\-12\-[0][1-2].*2022\-12\-[0][1-2].*\n.*\n.*\n.*

*/
