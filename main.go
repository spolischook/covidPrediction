package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	columnIsoCode = 0
	columnDate = 3
	columnNewCases = 5
)

type CovidStats struct {
	CountryIso string
	Date       time.Time
	NewCases   int
}
func main() {
	csvFile, err := os.Open("/home/spolischook/go/study/covid/owid-covid-data.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	for line, _ := reader.Read(); line != nil; line, _ = reader.Read() {
		date, err := time.Parse("2006-01-02", line[columnDate])
		if err != nil { continue }
		count, err := strconv.Atoi(line[columnNewCases])
		if err != nil { continue }

		stat := CovidStats{
			CountryIso: line[columnIsoCode],
			Date:       date,
			NewCases:   count,
		}
		fmt.Println(stat.CountryIso + " " + stat.Date.String() + " " + strconv.Itoa(stat.NewCases))
	}
}
