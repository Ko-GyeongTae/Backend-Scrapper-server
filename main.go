package main

import (
	"strings"

	"github.com/Ko-Gyeongtae/learngo/scrapper"
	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	//fmt.Println(term)
	scrapper.Scrape(term)
	return c.Attachment("jobs.csv", "jobs.csv")
}

func main(){
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":2000"))
	scrapper.Scrape("term")
}