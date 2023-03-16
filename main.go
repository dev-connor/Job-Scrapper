package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	location string
	title    string
	summary  string
}

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?keydownAccess=&searchType=search&searchword=python&panel_type=&search_optional_item=y&search_done=y&panel_count=y&preview=y&recruitSort=relation&recruitPageCount=40&inner_com_type=&company_cd=0%2C1%2C2%2C3%2C4%2C5%2C6%2C7%2C9%2C10&show_applied=&quick_apply=&except_read=&ai_head_hunting="

func main() {
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJob := <-c
		jobs = append(jobs, extractedJob...)

	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))

}

func getPage(page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page+1)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)

	}

	mainC <- jobs
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()

	})

	return pages
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title", "Location", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&recommend_ids=eJxFj8kVw0AIQ6vJHQFiOacQ999FZuyMOf4nISEnGll6FfDJrxulVbiwN3qwGRtjI8EGeqHcqAEzuUr7UWNl9SCq3E8yVcwsTzLhjnx7iYKnzW11ZRzVq1LbxuzO8vnKYrlHbUXL9BaNMapFxpgl7Fn0NwfT8Pa2aPqejx9jukA7&view_type=search&searchword=python&searchType=search&gz=1&t_ref_content=generic&t_ref=search&paid_fl=n&rec_idx=" + job.id, job.title, job.location, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}

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

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	title := cleanString(card.Find(".job_tit").Text())
	location := cleanString(card.Find(".job_condition").Text())
	summary := cleanString(card.Find(".job_sector").Text())

	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		summary:  summary,
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")

}
