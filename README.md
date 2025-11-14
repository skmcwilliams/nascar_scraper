# nascar_scraper
## NASCAR SCRAPER is not affiliated with NASCAR in any way,
#### This code returns end-of-race driver data based on desired seasons into a JSON file
#### Single season entried ought to use the same season in both parameters
#### Searching for an out-of-bounds season will return null for that season, if a season is out of bounds, data will be included for desired seasons that are in bounds
##### see nascar900_901.json and nascar2024_2026.json as examples

# nascar_scraper

This package scrapes Driver Averages (driveraverages.com) race tables and writes results to JSON files named using the pattern:


```
nascar<start>_<end>.json
```

Example: `nascar2018_2020.json`

## Generate JSON files

nascar_scraper exposes `GetRaceData(start_season int, end_season int)` which produces the JSON file `nascar<start>_<end>.json`.

Create a small `main.go` in this repo (replace `your/module/path` with the module path from your `go.mod`):

```go
package main

import (
    "log"

    "your/module/path/nascar_scraper"
)

func main() {
    // Generates nascar2018_2020.json in the current working directory
    nascar_scraper.GetRaceData(2018, 2020)
    log.Println("Done")
}
```

Run on macOS:

```bash
# from /Users/skm/CodeTemp/nascar_scraper
go run main.go
```

This will create `nascar2018_2020.json` (or corresponding filename for the seasons you request).

## Inspect and query the produced JSON

The JSON is an array of objects with fields such as Season, Race, Finish, Start, Number, Driver, Make, Pts, Laps, Led, Status, Team, Stage1, Stage2, Stage3, Rating.

Quick CLI inspection with jq:

```bash
# list all entries for driver "Logano"
jq '.[] | select(.Driver | test("Logano"; "i"))' nascar2018_2020.json

# show unique drivers
jq -r '.[].Driver' nascar2018_2020.json | sort | uniq
```

## Use the JSON from Go

Example snippet to load and use the generated JSON in Go (place in another small program in the repo):

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type TableData struct {
    Season int    `json:"Season"`
    Race   string `json:"Race"`
    Driver string `json:"Driver"`
    Finish string `json:"Finish"`
    // ... other fields omitted for brevity
}

func main() {
    b, err := os.ReadFile("nascar2018_2020.json")
    if err != nil {
        panic(err)
    }

    var rows []TableData
    if err := json.Unmarshal(b, &rows); err != nil {
        panic(err)
    }

    // Example: print first 10 rows
    for i := 0; i < 10 && i < len(rows); i++ {
        fmt.Printf("%d: %s - %s (Finish: %s)\n", rows[i].Season, rows[i].Race, rows[i].Driver, rows[i].Finish)
    }
}
```

## Index

- [Constants](#constants)
- [Variables](#variables)
- [Functions](#functions)
- [Types](#types)

## Constants

There are no exported constants in this package.

## Variables

There are no exported variables in this package.

## Functions

### GetRaceData

```go
func GetRaceData(start_season int, end_season int)
```

GetRaceData scrapes driver race data from driveraverages.com for a specified range of seasons and writes the results to a JSON file named `nascar<start>_<end>.json`.

**Parameters:**
- `start_season` (int): The starting season year for the data scrape
- `end_season` (int): The ending season year for the data scrape

**Returns:** None

**Note:** Single season queries should use the same value for both parameters. Out-of-bounds seasons will be excluded from results, with data returned for valid seasons only.

## Types

### tableData

```go
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
```

tableData represents a single driver's race result record containing:
- **Season**: The NASCAR season year
- **Race**: The race name/identifier
- **Finish**: Finishing position
- **Start**: Starting position
- **Number**: Driver's car number
- **Driver**: Driver's name
- **Make**: Vehicle manufacturer
- **Pts**: Points earned
- **Laps**: Laps completed
- **Led**: Laps led
- **Status**: Final status (e.g., finished, DNF)
- **Team**: Team name
- **Stage1**, **Stage2**, **Stage3**: Stage results
- **Rating**: Driver rating for the race
```

## Use the JSON from Python (pandas)

```python
import pandas as pd

df = pd.read_json("nascar2018_2020.json")
print(df.head())

# Example: group by driver and count races
counts = df.groupby("Driver").size().sort_values(ascending=False)
print(counts.head(20))
```
## Analyze the JSON data in Go

Example snippet to load and analyze the generated JSON in Go with more advanced queries:

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

type TableData struct {
    Season int    `json:"Season"`
    Race   string `json:"Race"`
    Driver string `json:"Driver"`
    Finish string `json:"Finish"`
    Start  string `json:"Start"`
    Number string `json:"Number"`
    Make   string `json:"Make"`
    Pts    string `json:"Pts"`
    Laps   string `json:"Laps"`
    Led    string `json:"Led"`
    Status string `json:"Status"`
    Team   string `json:"Team"`
    Stage1 string `json:"Stage1"`
    Stage2 string `json:"Stage2"`
    Stage3 string `json:"Stage3"`
    Rating string `json:"Rating"`
}

func main() {
    b, err := os.ReadFile("nascar2018_2020.json")
    if err != nil {
        panic(err)
    }

    var rows []TableData
    if err := json.Unmarshal(b, &rows); err != nil {
        panic(err)
    }

    // Example 1: Find all races for a specific driver
    driver := "Logano"
    fmt.Printf("\n=== All races for driver '%s' ===\n", driver)
    for _, row := range rows {
        if strings.Contains(row.Driver, driver) {
            fmt.Printf("Season %d - %s: Finished %s (Started %s)\n", 
                row.Season, row.Race, row.Finish, row.Start)
        }
    }

    // Example 2: Get unique drivers
    fmt.Println("\n=== Total unique drivers ===")
    driverSet := make(map[string]bool)
    for _, row := range rows {
        driverSet[row.Driver] = true
    }
    fmt.Printf("Found %d unique drivers\n", len(driverSet))

    // Example 3: Count races by season
    fmt.Println("\n=== Races per season ===")
    seasonRaces := make(map[int]int)
    for _, row := range rows {
        seasonRaces[row.Season]++
    }
    
    seasons := make([]int, 0, len(seasonRaces))
    for season := range seasonRaces {
        seasons = append(seasons, season)
    }
    sort.Ints(seasons)
    
    for _, season := range seasons {
        fmt.Printf("Season %d: %d races\n", season, seasonRaces[season])
    }

    // Example 4: Average finishing position per driver
    fmt.Println("\n=== Average finish position (top 5 drivers) ===")
    driverStats := make(map[string]struct{ total, count int })
    
    for _, row := range rows {
        finish, _ := strconv.Atoi(strings.TrimSpace(row.Finish))
        if finish > 0 {
            stats := driverStats[row.Driver]
            stats.total += finish
            stats.count++
            driverStats[row.Driver] = stats
        }
    }

    type driverAvg struct {
        driver string
        avg    float64
    }
    
    var avgs []driverAvg
    for driver, stats := range driverStats {
        if stats.count > 0 {
            avgs = append(avgs, driverAvg{driver, float64(stats.total) / float64(stats.count)})
        }
    }
    
    sort.Slice(avgs, func(i, j int) bool {
        return avgs[i].avg < avgs[j].avg
    })
    
    for i := 0; i < 5 && i < len(avgs); i++ {
        fmt.Printf("%s: %.2f avg finish\n", avgs[i].driver, avgs[i].avg)
    }
}
```

## Notes

- Replace `your/module/path` with the module path specified in your `go.mod`.
- The scraper follows the site's HTML structure; if the site changes, the code may need updates.
- Output filenames follow the `nascar<start>_<end>.json` pattern and are written to the current working directory.



