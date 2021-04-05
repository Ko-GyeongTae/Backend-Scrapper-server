/*
package main

import (
	"fmt"
	"time"
)

func main(){
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i< len(people); i++ {
		fmt.Println("Waiting for", i)
		fmt.Println("Received this message:", <-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is sexy"
}


func Count(person string) {
	for i:=0;i<10;i++{
		fmt.Println(person, "is good", i)
		time.Sleep(time.Second)
	}
}
*/

package scrapper

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
	salary   string
	summary  string
}

var baseURL string

// Scrape Indeed bby a term
func Scrape(term string) { //스크랩
	baseURL = "https://kr.indeed.com/jobs?q=" + term + "&limit=50" //검색 50제한
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages() //총 페이지 수

	for i := 0; i < totalPages; i++ {
		go getPage(i, c) // 페이지 정보 스크랩
	}

	for i := 0; i < totalPages; i++ {
		extractedJob := <-c // 페이지 정보 저장
		jobs = append(jobs, extractedJob...)
	}

	writeJobs(jobs) // 파일 작성
	fmt.Println("Done, extracted", len(jobs))
}

func getPage(page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // 함수 종료후 응답 body close

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) { // 엑셀 파일 작성
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title>a").Text())
	location := CleanString(card.Find(".sjcl").Text())
	salary := CleanString(card.Find(".salartText").Text())
	summary := CleanString(card.Find(".summary").Text())
	c <- extractedJob{ // 정보 객체
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func CleanString(str string) string { //스페이스, 탭 제거
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages() int { //페이지 측정
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

func writeJobs(jobs []extractedJob) { // 파일 작성시작
	file, err := os.Create("jobs.csv")
	checkErr(err)

	utf8bom := []byte{0xEF, 0xBB, 0xBF}
	file.Write(utf8bom)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		writeFile(jobSlice, w)
	}
}

func writeFile(jobSlice []string, w *csv.Writer) { //파일 작성
	jwErr := w.Write(jobSlice)
	checkErr(jwErr)
}

func checkErr(err error) { //에러 처리
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) { // 에러 출력
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
