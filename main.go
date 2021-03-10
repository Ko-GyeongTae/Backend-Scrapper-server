package main

import (
	"os"
	"strings"

	"github.com/Ko-Gyeongtae/learngo/scrapper"
	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error { //홈화면 html 호출
	return c.File("home.html")
}

func handleScrape(c echo.Context) error { //검색후 파일 다운로드
	defer os.Remove("jobs.csv") //파일 삭제 
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term) // 스크래핑
	return c.Attachment("jobs.csv", "jobs.csv") //파일 다운로드
}

func main(){
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":2000"))
	scrapper.Scrape("term")
}