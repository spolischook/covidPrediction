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
	statFile = "/home/spolischook/go/study/covid/owid-covid-data.csv"
	columnIsoCode = 0
	columnDate = 3
	columnNewCases = 5
)

type CovidStats struct {
	CountryIso string
	Date       time.Time
	NewCases   int
}
type CovidStatsIter struct {
	*csv.Reader
	file *os.File
}
func (i CovidStatsIter) Next() *CovidStats {
	for line, _ := i.Read(); line != nil; line, _ = i.Read() {
		date, err := time.Parse("2006-01-02", line[columnDate])
		if err != nil { continue }
		count, err := strconv.Atoi(line[columnNewCases])
		if err != nil { continue }

		stat := CovidStats{
			CountryIso: line[columnIsoCode],
			Date:       date,
			NewCases:   count,
		}

		return &stat
	}

	i.file.Close()

	return nil
}

func main() {
	csvFile, err := os.Open(statFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	iter := CovidStatsIter{reader, csvFile}

	for stat := iter.Next(); stat != nil; stat = iter.Next() {
		fmt.Println(stat.CountryIso + " " + stat.Date.String() + " " + strconv.Itoa(stat.NewCases))
	}
}
