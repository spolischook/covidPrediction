package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/DzananGanic/numericalgo"
	"github.com/DzananGanic/numericalgo/fit/linear"
	"github.com/gin-gonic/gin"
	"github.com/spolischook/covidPrediction/stats"
)

const statFile = "/home/spolischook/go/study/covid/owid-covid-data.csv"

func main() {
	router := gin.Default()
	router.GET("/covid-predict/:countryCode", handle)
	router.Run(":8080")
}

func handle(c *gin.Context) {
	iter := stats.NewIterator(statFile)
	defer iter.Close()

	countryCode := c.Param("countryCode")
	watchingDateStr := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
	expectedCases := predict(iter, countryCode, watchingDateStr)
	c.String(http.StatusOK, "Hello %d", int(expectedCases))
}

func predict(iter *stats.CovidStatsIter, country string, watchingDateStr string) float64 {
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
		"Expected new cases at %s(%d day): %f\n",
		watchingDateStr,
		int(judgmentDay),
		expectedNewCases)

	return expectedNewCases
}
