package module

import (
	"context"
	"fmt"
	"time"
	"tw_option_collector/data"
	"tw_option_collector/objects"

	"github.com/chromedp/chromedp"
)

// 抓取三大法人區分各期貨契約
func CrawlBigThreeFuture(ctx context.Context) {
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.taifex.com.tw/cht/3/futContractsDate"))
	handleError(err)
	startDate := time.Date(2024, time.January, 3, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2021, time.January, 4, 0, 0, 0, 0, time.UTC)
	for {
		fmt.Println("[Future] date=", startDate)
		c, err := data.FindBigThreeFutureCountByDate(startDate)
		handleError(err)
		if c != 0 {
			fmt.Println("data exist")
			startDate = startDate.AddDate(0, 0, -1)
			if startDate.Before(toDate) {
				break
			}
			continue
		}
		table1Cells := make([]string, 36)
		//先送出日期做查詢
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
			chromedp.EvaluateAsDevTools(`document.querySelector("#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody") !== null`, &isTableExist))
		handleError(err)
		fmt.Println(isTableExist)
		if isTableExist {
			tasks := chromedp.Tasks{
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(4)`, &table1Cells[0], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(5)`, &table1Cells[1], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(6)`, &table1Cells[2], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(7)`, &table1Cells[3], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(8)`, &table1Cells[4], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(9)`, &table1Cells[5], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(10)`, &table1Cells[6], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(11)`, &table1Cells[7], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(12)`, &table1Cells[8], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(13)`, &table1Cells[9], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(14)`, &table1Cells[10], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(15)`, &table1Cells[11], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(2)`, &table1Cells[12], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(3)`, &table1Cells[13], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(4)`, &table1Cells[14], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(5)`, &table1Cells[15], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(6)`, &table1Cells[16], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(7)`, &table1Cells[17], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(8)`, &table1Cells[18], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(9)`, &table1Cells[19], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(10)`, &table1Cells[20], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(11)`, &table1Cells[21], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(12)`, &table1Cells[22], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(13)`, &table1Cells[23], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(2)`, &table1Cells[24], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(3)`, &table1Cells[25], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(4)`, &table1Cells[26], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(5)`, &table1Cells[27], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(6)`, &table1Cells[28], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(7)`, &table1Cells[29], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(8)`, &table1Cells[30], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(9)`, &table1Cells[31], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(10)`, &table1Cells[32], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(11)`, &table1Cells[33], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(12)`, &table1Cells[34], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(13)`, &table1Cells[35], chromedp.ByQuery)}
			err = chromedp.Run(ctx, tasks)
			handleError(err)
			chromedp.Sleep(1 * time.Second)
			fmt.Println(table1Cells)
			//fmt.Println(table1Cells)

			data.InsertBigThreeFuture(&objects.BigThreeFuture{
				DataDate:                         startDate,
				DealerLongLotCount:               strToInt(table1Cells[0]),
				DealerLongMoney:                  strToDecimal(table1Cells[1]),
				DealerShortLotCount:              strToInt(table1Cells[2]),
				DealerShortMoney:                 strToDecimal(table1Cells[3]),
				DealerLotNetCount:                strToInt(table1Cells[4]),
				DealerNetMoney:                   strToDecimal(table1Cells[5]),
				OpenInterestDealerLongLotCount:   strToInt(table1Cells[6]),
				OpenInterestDealerLongMoney:      strToDecimal(table1Cells[7]),
				OpenInterestDealerShortLotCount:  strToInt(table1Cells[8]),
				OpenInterestDealerShortMoney:     strToDecimal(table1Cells[9]),
				OpenInterestDealerLotNetCount:    strToInt(table1Cells[10]),
				OpenInterestDealerNetMoney:       strToDecimal(table1Cells[11]),
				TrustLongLotCount:                strToInt(table1Cells[12]),
				TrustLongMoney:                   strToDecimal(table1Cells[13]),
				TrustShortLotCount:               strToInt(table1Cells[14]),
				TrustShortMoney:                  strToDecimal(table1Cells[15]),
				TrustLotNetCount:                 strToInt(table1Cells[16]),
				TrustNetMoney:                    strToDecimal(table1Cells[17]),
				OpenInterestTrustLongLotCount:    strToInt(table1Cells[18]),
				OpenInterestTrustLongMoney:       strToDecimal(table1Cells[19]),
				OpenInterestTrustShortLotCount:   strToInt(table1Cells[20]),
				OpenInterestTrustShortMoney:      strToDecimal(table1Cells[21]),
				OpenInterestTrustLotNetCount:     strToInt(table1Cells[22]),
				OpenInterestTrustNetMoney:        strToDecimal(table1Cells[23]),
				ForeignLongLotCount:              strToInt(table1Cells[24]),
				ForeignLongMoney:                 strToDecimal(table1Cells[25]),
				ForeignShortLotCount:             strToInt(table1Cells[26]),
				ForeignShortMoney:                strToDecimal(table1Cells[27]),
				ForeignLotNetCount:               strToInt(table1Cells[28]),
				ForeignNetMoney:                  strToDecimal(table1Cells[29]),
				OpenInterestForeignLongLotCount:  strToInt(table1Cells[30]),
				OpenInterestForeignLongMoney:     strToDecimal(table1Cells[31]),
				OpenInterestForeignShortLotCount: strToInt(table1Cells[32]),
				OpenInterestForeignShortMoney:    strToDecimal(table1Cells[33]),
				OpenInterestForeignLotNetCount:   strToInt(table1Cells[34]),
				OpenInterestForeignNetMoney:      strToDecimal(table1Cells[35]),
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
