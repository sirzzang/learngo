package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	pages := getPages()
	fmt.Println(pages)
}

func getPages() int {
	resp, err := http.Get(baseURL)
	checkErr(err)
	checkCode(resp)

	defer resp.Body.Close() // prevent memory leak

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	fmt.Println(doc)

	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalln("Request failed with status:", resp.StatusCode)
	}
}
