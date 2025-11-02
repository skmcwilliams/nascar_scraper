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

## Use the JSON from Python (pandas)

```python
import pandas as pd

df = pd.read_json("nascar2018_2020.json")
print(df.head())

# Example: group by driver and count races
counts = df.groupby("Driver").size().sort_values(ascending=False)
print(counts.head(20))
```

## Notes

- Replace `your/module/path` with the module path specified in your `go.mod`.
- The scraper follows the site's HTML structure; if the site changes, the code may need updates.
- Output filenames follow the `nascar<start>_<end>.json` pattern and are written to the current working directory.



