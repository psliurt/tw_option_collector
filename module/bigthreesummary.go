package module

import (
	"context"
	"fmt"
	"time"
	"tw_option_collector/data"
	"tw_option_collector/objects"

	"github.com/chromedp/chromedp"
)

// 抓取三大法人總表
func CrawlBigThreeSummary(ctx context.Context) {
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.taifex.com.tw/cht/3/totalTableDate"))
	handleError(err)
	startDate := time.Date(2024, time.January, 3, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2021, time.January, 4, 0, 0, 0, 0, time.UTC)
	for {
		fmt.Println("[Summary] date=", startDate)
		c, err := data.FindBigThreeSummaryCountByDate(startDate)
		handleError(err)
		if c != 0 {
			fmt.Println("data exist")
			startDate = startDate.AddDate(0, 0, -1)
			if startDate.Before(toDate) {
				break
			}
			continue
		}
		table1Cells := make([]string, 18)
		table2Cells := make([]string, 18)

		sendQueryTasks := chromedp.Tasks{chromedp.Sleep(500 * time.Millisecond),
			chromedp.Evaluate(`document.querySelector("#queryDate").value = ""`, nil),
			chromedp.Sleep(100 * time.Millisecond),
			chromedp.SendKeys("#queryDate", startDate.Format("2006/01/02"), chromedp.NodeVisible, chromedp.ByID),
			chromedp.Sleep(100 * time.Millisecond),
			chromedp.Focus(`#uForm > table > tbody > tr:nth-child(3) > td > #button`),
			chromedp.Click(`#uForm > table > tbody > tr:nth-child(3) > td > #button`, chromedp.NodeVisible),
			chromedp.Sleep(4000 * time.Millisecond)}
		err = chromedp.Run(ctx, sendQueryTasks)
		handleError(err)

		isTableExist := false
		err = chromedp.Run(ctx,
			chromedp.EvaluateAsDevTools(`document.querySelector("#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c") !== null`, &isTableExist))
		handleError(err)

		if isTableExist {
			tasks := chromedp.Tasks{
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(4) > td:nth-child(2) > div`, &table1Cells[0], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(4) > td:nth-child(3) > div`, &table1Cells[1], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(4) > td:nth-child(4) > div`, &table1Cells[2], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(4) > td:nth-child(5) > div`, &table1Cells[3], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(4) > td:nth-child(6) > div`, &table1Cells[4], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(4) > td:nth-child(7) > div`, &table1Cells[5], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(5) > td:nth-child(2) > div`, &table1Cells[6], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(5) > td:nth-child(3) > div`, &table1Cells[7], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(5) > td:nth-child(4) > div`, &table1Cells[8], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(5) > td:nth-child(5) > div`, &table1Cells[9], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(5) > td:nth-child(6) > div`, &table1Cells[10], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(5) > td:nth-child(7) > div`, &table1Cells[11], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(6) > td:nth-child(2) > div`, &table1Cells[12], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(6) > td:nth-child(3) > div`, &table1Cells[13], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(6) > td:nth-child(4) > div`, &table1Cells[14], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(6) > td:nth-child(5) > div`, &table1Cells[15], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(6) > td:nth-child(6) > div`, &table1Cells[16], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_f > tbody > tr:nth-child(6) > td:nth-child(7) > div`, &table1Cells[17], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(4) > td:nth-child(2) > div`, &table2Cells[0], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(4) > td:nth-child(3) > div`, &table2Cells[1], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(4) > td:nth-child(4) > div`, &table2Cells[2], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(4) > td:nth-child(5) > div`, &table2Cells[3], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(4) > td:nth-child(6) > div`, &table2Cells[4], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(4) > td:nth-child(7) > div`, &table2Cells[5], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(5) > td:nth-child(2) > div`, &table2Cells[6], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(5) > td:nth-child(3) > div`, &table2Cells[7], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(5) > td:nth-child(4) > div`, &table2Cells[8], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(5) > td:nth-child(5) > div`, &table2Cells[9], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(5) > td:nth-child(6) > div`, &table2Cells[10], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(5) > td:nth-child(7) > div`, &table2Cells[11], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(6) > td:nth-child(2) > div`, &table2Cells[12], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(6) > td:nth-child(3) > div`, &table2Cells[13], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(6) > td:nth-child(4) > div`, &table2Cells[14], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(6) > td:nth-child(5) > div`, &table2Cells[15], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(6) > td:nth-child(6) > div`, &table2Cells[16], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table.table_c > tbody > tr:nth-child(6) > td:nth-child(7) > div`, &table2Cells[17], chromedp.ByQuery)}
			err = chromedp.Run(ctx, tasks)
			handleError(err)
			chromedp.Sleep(1 * time.Second)
			fmt.Println(table1Cells)
			fmt.Println(table2Cells)

			data.InsertBigThreeSummary(&objects.BigThreeSummary{
				DataDate:                         startDate,
				DealerLongLotCount:               strToInt(table1Cells[0]),
				DealerLongMoney:                  strToDecimal(table1Cells[1]),
				DealerShortLotCount:              strToInt(table1Cells[2]),
				DealerShortMoney:                 strToDecimal(table1Cells[3]),
				DealerLotNetCount:                strToInt(table1Cells[4]),
				DealerNetMoney:                   strToDecimal(table1Cells[5]),
				TrustLongLotCount:                strToInt(table1Cells[6]),
				TrustLongMoney:                   strToDecimal(table1Cells[7]),
				TrustShortLotCount:               strToInt(table1Cells[8]),
				TrustShortMoney:                  strToDecimal(table1Cells[9]),
				TrustLotNetCount:                 strToInt(table1Cells[10]),
				TrustNetMoney:                    strToDecimal(table1Cells[11]),
				ForeignLongLotCount:              strToInt(table1Cells[12]),
				ForeignLongMoney:                 strToDecimal(table1Cells[13]),
				ForeignShortLotCount:             strToInt(table1Cells[14]),
				ForeignShortMoney:                strToDecimal(table1Cells[15]),
				ForeignLotNetCount:               strToInt(table1Cells[16]),
				ForeignNetMoney:                  strToDecimal(table1Cells[17]),
				OpenInterestDealerLongLotCount:   strToInt(table2Cells[0]),
				OpenInterestDealerLongMoney:      strToDecimal(table2Cells[1]),
				OpenInterestDealerShortLotCount:  strToInt(table2Cells[2]),
				OpenInterestDealerShortMoney:     strToDecimal(table2Cells[3]),
				OpenInterestDealerLotNetCount:    strToInt(table2Cells[4]),
				OpenInterestDealerNetMoney:       strToDecimal(table2Cells[5]),
				OpenInterestTrustLongLotCount:    strToInt(table2Cells[6]),
				OpenInterestTrustLongMoney:       strToDecimal(table2Cells[7]),
				OpenInterestTrustShortLotCount:   strToInt(table2Cells[8]),
				OpenInterestTrustShortMoney:      strToDecimal(table2Cells[9]),
				OpenInterestTrustLotNetCount:     strToInt(table2Cells[10]),
				OpenInterestTrustNetMoney:        strToDecimal(table2Cells[11]),
				OpenInterestForeignLongLotCount:  strToInt(table2Cells[12]),
				OpenInterestForeignLongMoney:     strToDecimal(table2Cells[13]),
				OpenInterestForeignShortLotCount: strToInt(table2Cells[14]),
				OpenInterestForeignShortMoney:    strToDecimal(table2Cells[15]),
				OpenInterestForeignLotNetCount:   strToInt(table2Cells[16]),
				OpenInterestForeignNetMoney:      strToDecimal(table2Cells[17]),
			})

		} else {
			chromedp.Sleep(2 * time.Second)
		}

		startDate = startDate.AddDate(0, 0, -1)
		if startDate.Before(toDate) {
			break
		}
	}

}
