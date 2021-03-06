package main

import (
	"net/http"

	"github.com/Ko-Gyeongtae/learngo/scrapper"
	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}
func main(){
	e := echo.New()
	e.GET("/", handleHome)
	e.Logger.Fatal(e.Start(":2000"))
	scrapper.Scrape("term")
}