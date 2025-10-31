package main

import (
	"encoding/json"

	"fmt"

	"github.com/gocolly/colly"

	"strings"

	"os"
	//"strconv"
)

type tableData struct {
	Season int
	Race   string
	Finish string
	Start  string
	Number string
	Driver string
	Make   string
	Pts    string
	Laps   string
	Led    string
	Status string
	Team   string
	Stage1 string
	Stage2 string
	Stage3 string
	Rating string
}

func race_stats(start int, end int) []tableData {

	var raceData []tableData
	//var seasonData []tableData

	for season := start; season <= end; season++ {
		c := colly.NewCollector()

		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			race := e.Text
			raceString := fmt.Sprintf("race.php?sked_id=%d0", season)
			if strings.Contains(link, raceString) {
				c.Visit(e.Request.AbsoluteURL(link))
				c.OnHTML("table.sortable.tabledata-nascar.table-large", func(h *colly.HTMLElement) {
					h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
						tableData := tableData{
							Season: season,
							Race:   race,
							Finish: el.ChildText("td:nth-child(1)"),
							Start:  el.ChildText("td:nth-child(2)"),
							Number: el.ChildText("td:nth-child(3)"),
							Driver: el.ChildText("td:nth-child(4)"),
							Make:   el.ChildText("td:nth-child(5)"),
							Pts:    el.ChildText("td:nth-child(6)"),
							Laps:   el.ChildText("td:nth-child(7)"),
							Led:    el.ChildText("td:nth-child(8)"),
							Status: el.ChildText("td:nth-child(9)"),
							Team:   el.ChildText("td:nth-child(10)"),
							Stage1: el.ChildText("td:nth-child(11)"),
							Stage2: el.ChildText("td:nth-child(12)"),
							Stage3: el.ChildText("td:nth-child(13)"),
							Rating: el.ChildText("td:nth-child(14)"),
						}
						raceData = append(raceData, tableData)

					})

				})
			}

		})

		c.OnResponse(func(r *colly.Response) {
			fmt.Println("Status: ", r.StatusCode)
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})

		startURL := fmt.Sprintf("https://www.driveraverages.com/nascar/year.php?yr_id=%d", season)
		c.Visit(startURL)
	}

	return raceData

}

func to_json(x []tableData, start int, end int) {
	// convert to JSON file
	content, err := json.Marshal(x)
	if err != nil {
		fmt.Println(err.Error())
	}
	filename := fmt.Sprintf("nascar%d_%d.json", start, end)
	os.WriteFile(filename, content, 0644)

}

func Get_race_data(start_season int, end_season int) {
	df := race_stats(start_season, end_season)
	to_json(df, start_season, end_season)
}
