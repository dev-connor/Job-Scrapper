package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"
var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"

func main() {
	getPages()

}

func getPages() int {
	// req, rErr := http.NewRequest("GET", baseURL, nil)
	// checkErr(rErr)

	// 프록시로 호출
	// purl, err := url.Parse(baseURL)
	// client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(purl)}}

	// res, err := client.Do(req)
	// checkErr(err)
	// checkCode(res)

	// defer res.Body.Close()

	// doc, err := goquery.NewDocumentFromReader(res.Body)
	// checkErr(err)

	// fmt.Println(doc)

	// doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
	// 	fmt.Println(s.Html())

	// })

	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination")

	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}

}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
