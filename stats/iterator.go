package stats

import (
	"encoding/csv"
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

type CovidStatsIter struct {
	*csv.Reader
	file *os.File
}

func NewIterator(filePath string) *CovidStatsIter {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(csvFile)
	iter := CovidStatsIter{reader, csvFile}

	return &iter
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

	return nil
}

func (i CovidStatsIter) Close() error {
	return i.file.Close()
}
