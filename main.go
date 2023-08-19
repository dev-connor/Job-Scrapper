package main

import (
	"Job-Scrapper/app"
	"Job-Scrapper/pkg"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)

	startTime := time.Now()
	term := strings.ToLower(app.CleanString(c.FormValue("term")))
	app.Scrape(term)

	endTime := time.Now()
	duration := pkg.TimeToKo(endTime.Sub(startTime))
	fmt.Printf("CSV 파일을 다운로드하는데 %s 걸렸습니다.\n", duration)

	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()

	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)

	e.Logger.Fatal(e.Start(":1323"))

}
