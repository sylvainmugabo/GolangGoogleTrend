package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	res, err := makeRequest()
	if err != nil {
		fmt.Println(err)
		return
	}

	data := readTrendsResponse(res)
	var r RSS
	err = xml.Unmarshal(data, &r)
	if err != nil {
		fmt.Println("Error occur in parsing")
		os.Exit(1)
	}

	fmt.Println("\n Below are all the Google Search Trends For Today !")
	fmt.Println("-------------------------------------------------------")

	for i := range r.Channel.Items {
		rank := i + 1
		fmt.Println("#", rank)
		fmt.Println("Search Term:", r.Channel.Items[i].Title)
		fmt.Println("Link to the Trend:", r.Channel.Items[i].Link)
		fmt.Println("Headline:", r.Channel.Items[i].NewsItems[0].Title)
		fmt.Println("Link to article:", r.Channel.Items[i].NewsItems[0].Title)
		fmt.Println("-------------------------------------------------------")
	}
}

func makeRequest() (*http.Response, error) {
	res, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")

	return res, err
}

func readTrendsResponse(res *http.Response) []byte {

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}

type RSS struct {
	XmlElement xml.Name `xml:"rss"`
	Channel    Channel  `xml:"channel"`
}
type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Traffic     string `xml:"approx_traffic"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	NewsItems   []News `xml:"news_item"`
}

type News struct {
	Title   string `xml:"news_item_title"`
	Snippet string `xml:"news_item_snippet"`
	Url     string `xml:"news_item_url"`
	Source  string `xml:"news_item_source"`
}
