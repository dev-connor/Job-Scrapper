package main

import (
	"Job-Scrapper/scrapper"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}
func handleScrape(c echo.Context) error {
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	fmt.Println(term)

	return nil
}

func main() {
	e := echo.New()
	// e.GET()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))

}
