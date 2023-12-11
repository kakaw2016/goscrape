package Artscraper

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

			reg1 := regexp.MustCompile("((oadissin)|(casimodo)|(hermansa)|(fikif)|(thalibul)|(ybf))")
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

				//reg1 := regexp.MustCompile("(.*(hh guild).*|.*(weed).*|.*(gimmi).*|.*(mage).*|.*(lolz).*|.*(powerup).*|.*(psyber).*|.*(risingstar).*|.*(virtualgrowth).*|.*(star).*|.*(rising).*|.*(crop).*|.*(woo).*|.*(chess).*|.*(phortun).*|.*(woogame).*|.*(euro).*|.*(goldquizblog).*|.*(hashkings).*)")

				reg1 := regexp.MustCompile("(.*(actifit).*)")
				for _, value2 := range value {

					value2 = strings.ToLower(value2)

					if reg1.MatchString(value2) {
						stock.Title = ""
						break

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

			dayAgoTime1 := time.Now().AddDate(0, 0, -3).Format("2006-01-02T15:04:05")

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

			reg1 := regexp.MustCompile("(.+(contes).+|.+(challeng).+|.+(concou).+|.+(compet).+)")

			reg2 := regexp.MustCompile("(.+@)")

			if reg1.MatchString(value[0]) {

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

func collectUserNames(urls []string) []string {
	var usernames []string
	for _, url := range urls {

		reg1 := regexp.MustCompile("@[a-z0-9-.]+")

		user := reg1.FindString(url)

		usernames = append(usernames, user)

	}

	return noduplicate(usernames)

}

func filterLinksTounique(urls []string) []string {

	var uniqueUrls []string

	//var collectedUrls []string

	data := make(map[string][]string)

	usernames := collectUserNames(urls)

	for i, username := range usernames {

		if i < len(usernames)+1 {

			for _, url := range urls {

				if strings.Contains(url, username) {

					data[username] = append(data[username], url)

				}

			}

			for key, experience := range data {
				if key == username {
					experience = data[key]

					uniqueUrls = append(uniqueUrls, fmt.Sprint(experience[len(experience)-1]))

				}

			}

			i++

		}

	}

	return noduplicate(uniqueUrls)

}

func Initialized() {

	var blockUrls []string

	blockUrls1 := readfile("/home/kakashinaruto/go/src/github.com/kakaw2016/goscrape/ArtCurationHive/Artconnectlinkscrape.txt")

	blockUrls2 := readfile("/home/kakashinaruto/go/src/github.com/kakaw2016/goscrape/ArtCurationHive/Artposts")

	blockUrls = append(blockUrls, blockUrls1...)
	blockUrls = append(blockUrls, blockUrls2...)
	blockUrls = noduplicate(blockUrls)

	blockUrls = filterLinksTounique(blockUrls)

	fileStoredata, err := os.OpenFile("/home/kakashinaruto/go/src/github.com/kakaw2016/goscrape/ArtCurationHive/Artlivescrape.txt", os.O_CREATE|os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fileStoredata.Close()

	fmt.Println("Total URL", len(blockUrls))

	const headtext string = `
<div class="text-justify">

**Dear friends,**

**Peace,**

</div>

<div class="text-justify">


How are you doing on this beautiful morning?

The art contest article and participants' entries identification has never been easier.

I tried to collect the maximum of publications within the last five days on the Hive ecosystem.

All contest posts missed in the review are considered a contribution and welcome in the comment section.

Enjoy your venture into this wonderful universe of creativity on Hive.


<div class="text-center">

[![freeimageSource](https://files.peakd.com/file/peakd-hive/ybf/242NfAmkdjKdxEJdWxWvdrtz3K955ydTjqep8ZSjZVKdKUtwjM5MvJayGg8fzkfqdDajS.gif)](https://www.canva.com/design/DAFhY98LjxA/bk-QSDT3l92E-VfJcYZynw/edit?utm_content=DAFhY98LjxA&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton)

</div>

--

The compilation of articles searched and filtered through hundreds of contests, and publications related to art topics on Hive.

I hope that this list will grow as the community adds the articles that I have missed in this weekly project. 


--

<div class="text-justify">

So I thought I should post and maintain here an updated directory of several contests, call to action, or raffles related to arts.

The community members are welcome to participate in updating this list. By doing so this project can achieve the main goal as a reference to all bloggers.


--


<div class="text-center">


--

</div>

<div class="text-center">

Enjoy...

</div>

<div class="text-center">

..*..*..*..*..*..*
</div>

<div class="text-justify">
`
	const endtext string = `
</div>

<div class="text-center">

..*..*..*..*..*..*..*

</div>




<div class="text-justify"> 


The entries here in our review list can motivate you all the engage with the authors and even join the event by checking the community hosting the event. 

Please check your favorite contest main page to place your entry before the deadline.
</div>


<div class="text-center">


[![freeimageSource](https://images.hive.blog/500x0/https://files.peakd.com/file/peakd-hive/ybf/23xeva1tAhfSdq9cFuta72ybuBVe7s7WbZg4jgzjYdkjppVySw6JE5DZpXbVfUZNvhA8g.jpg)](https://www.canva.com/design/DAFhY98LjxA/bk-QSDT3l92E-VfJcYZynw/edit?utm_content=DAFhY98LjxA&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton)


</div>

Feel free to read more about previous art-reviewed articles on my Hive publication page. Please visit the following blog page.


<div class="text-center"> 

++[Blog Profile](https://ecency.com/@ybf/posts)++

</div>


<div class="text-center"> 


## Warm regards
</div>

Weekly Art - Contest + Entries Recap: 30+ CreativeWork // Issue n.50

art photofeed creativecoin photography arte contest challenge streetart graffiti digitalart paintings illustration aliens cervantes neoxian ocd waivio alive proofofbrain meme waiv tribes archon vyb ctp trliste palnet

Description

How are you doing Art Fan on Hive? Here are some Challenges for the current week's contests. The links were gathered from multiple communities on Hive.
`

	w := bufio.NewWriter(fileStoredata)

	dataToStore1 := fmt.Sprintf("\n%s\n", headtext)

	_, _ = w.WriteString(dataToStore1)
	w.Flush()

	var collinfo product

	for _, artPosts := range blockUrls {

		postCodeSource := scrapesource(artPosts)
		collinfo.collectedata(postCodeSource)

		if collinfo.Title != "" && collinfo.Images != "" && collinfo.Url != "" && collinfo.Voters != "" && collinfo.Authors != "" && len(collinfo.PostAge) != 0 {

			dataToStore2 := fmt.Sprintf("\n (•◡•) /\n%s\n\n<div class=\"text-center\">\n\n[![](https://images.hive.blog/450x0/%s)](%s)\n</div>\nPosted Since: %s\n\nVoted By %s Hive Bloggers\nArticle Creator: %s\n", collinfo.Title, collinfo.Images, collinfo.Url, collinfo.PostAge[:], collinfo.Voters, collinfo.Authors)

			_, _ = w.WriteString(dataToStore2)
			w.Flush()
		}
	}

	dataToStore3 := fmt.Sprintf("\n%s\n", endtext)

	_, _ = w.WriteString(dataToStore3)
	w.Flush()

	/*for _, artPosts := range blockUrls[:1] {
		fmt.Println(artPosts)
		z := scrapesource(artPosts)
		for k1, z1 := range z {
			fmt.Println(k1, "======>", z1)

		}

	}*/

}

func Cronjob() {

	ch := gocron.Start()

	gocron.Every(30).Minutes().Do(Initialized)

	//go test(ch)

	_, time := gocron.NextRun()
	fmt.Println("Next Schedule In: ", time)

	<-ch

}

/*
Filtration

Date to topic

\n.*•◡•\).*\n.*\n.*\n.*\n.*\n.*((contes)|(challenge)).*\n.*\n.*\n.*\n.*\n.*

*/
