package main

import (
	"fmt"
	"github.com/DzananGanic/numericalgo"
	"github.com/DzananGanic/numericalgo/fit/linear"
	"strconv"
	"time"

	"github.com/spolischook/covidPrediction/stats"
)

const statFile = "/home/spolischook/go/study/covid/owid-covid-data.csv"

func main() {
	iter := stats.NewIterator(statFile)
	defer iter.Close()

	predict(iter, "UKR", "2021-09-01")
}

func predict(iter *stats.CovidStatsIter, country string, watchingDateStr string) {
	var startDate *time.Time
	var x numericalgo.Vector
	var y numericalgo.Vector

	for stat := iter.Next(); stat != nil; stat = iter.Next() {
		if country != stat.CountryIso {
			continue
		}
		if startDate == nil {
			startDate = &stat.Date
		}

		day := stat.Date.Sub(*startDate).Hours() / 24
		x = append(x, day)
		y = append(y, float64(stat.NewCases))

		fmt.Println(stat.CountryIso + " " + strconv.Itoa(int(day)) + " " + strconv.Itoa(stat.NewCases))
	}

	watchingDate, _ := time.Parse("2006-01-02", watchingDateStr)
	judgmentDay := watchingDate.Sub(*startDate).Hours() / 24
	lf := linear.New()
	_ = lf.Fit(x, y)
	expectedNewCases := lf.Predict(judgmentDay)

	fmt.Printf(
		"Expected new cases at %s(%d day): %d",
		watchingDateStr,
		int(judgmentDay),
		int(expectedNewCases))
}
