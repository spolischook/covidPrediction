package main

import (
	"fmt"
	"strconv"

	"github.com/spolischook/covidPrediction/stats"
)

const statFile = "/home/spolischook/go/study/covid/owid-covid-data.csv"

func main() {
	iter := stats.NewIterator(statFile)
	defer iter.Close()

	for stat := iter.Next(); stat != nil; stat = iter.Next() {
		fmt.Println(stat.CountryIso + " " + stat.Date.String() + " " + strconv.Itoa(stat.NewCases))
	}
}
