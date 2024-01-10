package module

import (
	"fmt"
	"time"
	"tw_option_collector/data"
	"tw_option_collector/objects"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/mongo"
)

func AnalyzeSummary(date time.Time) {

	dateArray := make([]string, 0)
	for i := 31; i >= 0; i-- {
		d := date.AddDate(0, 0, i*-1)
		rawSummary, err := data.FindBigThreeSummaryByDate(d)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				continue
			} else {
				panic(err)
			}
		}

		fmt.Println("Date: ", d.Format("2006-01-02"))
		fmt.Println("外資交易口數淨額: ", rawSummary.ForeignLotNetCount, " 外資未平倉口數淨額: ", rawSummary.OpenInterestForeignLotNetCount)
		if rawSummary.ForeignLotNetCount >= 0 && rawSummary.OpenInterestForeignLotNetCount < 0 {
			fmt.Println("外資盤中交易[ 轉多 ]")
			dateArray = append(dateArray, d.Format("2006-01-02"))
		}
		if rawSummary.ForeignLotNetCount < 0 && rawSummary.OpenInterestForeignLotNetCount >= 0 {
			fmt.Println("外資盤中交易[ 轉空 ]")
			dateArray = append(dateArray, d.Format("2006-01-02"))
		}

		fmt.Println("============================================================================")

	}
	for _, v := range dateArray {
		fmt.Println(v)
	}

}

func AnalyzePutCallRatio() {
	ratioList := make([]*objects.MarketPutCallRatio, 0)
	today := time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC)
	for i := 0; i <= 100; i++ {
		ratio, err := data.FindPutCallRatioByDate(today)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				//fmt.Println("--today=", today)
				today = today.AddDate(0, 0, -1)
				continue
			}
			handleError(err)
		}
		ratioList = append(ratioList, ratio)
		today = today.AddDate(0, 0, -1)
		//fmt.Println("today=", today)
	}
	for _, v := range ratioList {
		if v.PutCallOIRatio.LessThan(decimal.NewFromFloat(90)) {
			fmt.Println("date: ", v.DataDate, " tradeRatio: ", v.PutCallTradeRatio, " oiratio: ", v.PutCallOIRatio)
		}
	}
}
