package module

import (
	"context"
	"fmt"
	"time"
	"tw_option_collector/data"
	"tw_option_collector/objects"

	"github.com/chromedp/chromedp"
)

// 抓取三大法人區分選擇權與期貨二類
func CrawlBigThreeFutureOption(ctx context.Context) {
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.taifex.com.tw/cht/3/futAndOptDate"))
	handleError(err)
	startDate := time.Date(2024, time.January, 3, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2021, time.January, 4, 0, 0, 0, 0, time.UTC)
	for {
		fmt.Println("[FutureOption] date=", startDate)
		c, err := data.FindBigThreeFutureOptionCountByDate(startDate)
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
		table2Cells := make([]string, 36)
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
			chromedp.EvaluateAsDevTools(`document.querySelector("#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4)") !== null`, &isTableExist))
		handleError(err)

		if isTableExist {
			tasks := chromedp.Tasks{
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(5) > td:nth-child(2)`, &table1Cells[0], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(5) > td:nth-child(3)`, &table1Cells[1], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(5) > td:nth-child(4)`, &table1Cells[2], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(5) > td:nth-child(5)`, &table1Cells[3], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(5) > td:nth-child(6)`, &table1Cells[4], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(5) > td:nth-child(7)`, &table1Cells[5], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(5) > td:nth-child(8)`, &table1Cells[6], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(5) > td:nth-child(9)`, &table1Cells[7], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(5) > td:nth-child(10)`, &table1Cells[8], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(5) > td:nth-child(11)`, &table1Cells[9], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(5) > td:nth-child(12)`, &table1Cells[10], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(5) > td:nth-child(13)`, &table1Cells[11], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(6) > td:nth-child(2)`, &table1Cells[12], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(6) > td:nth-child(3)`, &table1Cells[13], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(6) > td:nth-child(4)`, &table1Cells[14], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(6) > td:nth-child(5)`, &table1Cells[15], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(6) > td:nth-child(6)`, &table1Cells[16], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(6) > td:nth-child(7)`, &table1Cells[17], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(6) > td:nth-child(8)`, &table1Cells[18], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(6) > td:nth-child(9)`, &table1Cells[19], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(6) > td:nth-child(10)`, &table1Cells[20], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(6) > td:nth-child(11)`, &table1Cells[21], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(6) > td:nth-child(12)`, &table1Cells[22], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(6) > td:nth-child(13)`, &table1Cells[23], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(7) > td:nth-child(2)`, &table1Cells[24], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(7) > td:nth-child(3)`, &table1Cells[25], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(7) > td:nth-child(4)`, &table1Cells[26], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(7) > td:nth-child(5)`, &table1Cells[27], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(7) > td:nth-child(6)`, &table1Cells[28], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(7) > td:nth-child(7)`, &table1Cells[29], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(7) > td:nth-child(8)`, &table1Cells[30], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(7) > td:nth-child(9)`, &table1Cells[31], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(7) > td:nth-child(10)`, &table1Cells[32], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(7) > td:nth-child(11)`, &table1Cells[33], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(7) > td:nth-child(12)`, &table1Cells[34], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(7) > td:nth-child(13)`, &table1Cells[35], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(5) > td:nth-child(2)`, &table2Cells[0], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(5) > td:nth-child(3)`, &table2Cells[1], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(5) > td:nth-child(4)`, &table2Cells[2], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(5) > td:nth-child(5)`, &table2Cells[3], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(5) > td:nth-child(6)`, &table2Cells[4], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(5) > td:nth-child(7)`, &table2Cells[5], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(5) > td:nth-child(8)`, &table2Cells[6], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(5) > td:nth-child(9)`, &table2Cells[7], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(5) > td:nth-child(10)`, &table2Cells[8], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(5) > td:nth-child(11)`, &table2Cells[9], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(5) > td:nth-child(12)`, &table2Cells[10], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(5) > td:nth-child(13)`, &table2Cells[11], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(6) > td:nth-child(2)`, &table2Cells[12], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(6) > td:nth-child(3)`, &table2Cells[13], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(6) > td:nth-child(4)`, &table2Cells[14], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(6) > td:nth-child(5)`, &table2Cells[15], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(6) > td:nth-child(6)`, &table2Cells[16], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(6) > td:nth-child(7)`, &table2Cells[17], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(6) > td:nth-child(8)`, &table2Cells[18], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(6) > td:nth-child(9)`, &table2Cells[19], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(6) > td:nth-child(10)`, &table2Cells[20], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(6) > td:nth-child(11)`, &table2Cells[21], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(6) > td:nth-child(12)`, &table2Cells[22], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(6) > td:nth-child(13)`, &table2Cells[23], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(7) > td:nth-child(2)`, &table2Cells[24], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(7) > td:nth-child(3)`, &table2Cells[25], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(7) > td:nth-child(4)`, &table2Cells[26], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(7) > td:nth-child(5)`, &table2Cells[27], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(7) > td:nth-child(6)`, &table2Cells[28], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(7) > td:nth-child(7)`, &table2Cells[29], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(7) > td:nth-child(8)`, &table2Cells[30], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(7) > td:nth-child(9)`, &table2Cells[31], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(7) > td:nth-child(10)`, &table2Cells[32], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(7) > td:nth-child(11)`, &table2Cells[33], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(7) > td:nth-child(12)`, &table2Cells[34], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table:nth-child(4) > tbody > tr:nth-child(7) > td:nth-child(13)`, &table2Cells[35], chromedp.ByQuery)}
			err = chromedp.Run(ctx, tasks)
			handleError(err)
			chromedp.Sleep(1 * time.Second)
			fmt.Println(table1Cells)
			fmt.Println(table2Cells)

			data.InsertBigThreeFutureOption(&objects.BigThreeFutureOption{
				DataDate:                               startDate,
				DealerFutureLongLotCount:               strToInt(table1Cells[0]),
				DealerFutureLongMoney:                  strToDecimal(table1Cells[2]),
				DealerFutureShortLotCount:              strToInt(table1Cells[4]),
				DealerFutureShortMoney:                 strToDecimal(table1Cells[6]),
				DealerFutureLotNetCount:                strToInt(table1Cells[8]),
				DealerFutureNetMoney:                   strToDecimal(table1Cells[10]),
				DealerOptionLongLotCount:               strToInt(table1Cells[1]),
				DealerOptionLongMoney:                  strToDecimal(table1Cells[3]),
				DealerOptionShortLotCount:              strToInt(table1Cells[5]),
				DealerOptionShortMoney:                 strToDecimal(table1Cells[7]),
				DealerOptionLotNetCount:                strToInt(table1Cells[9]),
				DealerOptionNetMoney:                   strToDecimal(table1Cells[11]),
				TrustFutureLongLotCount:                strToInt(table1Cells[12]),
				TrustFutureLongMoney:                   strToDecimal(table1Cells[14]),
				TrustFutureShortLotCount:               strToInt(table1Cells[16]),
				TrustFutureShortMoney:                  strToDecimal(table1Cells[18]),
				TrustFutureLotNetCount:                 strToInt(table1Cells[20]),
				TrustFutureNetMoney:                    strToDecimal(table1Cells[22]),
				TrustOptionLongLotCount:                strToInt(table1Cells[13]),
				TrustOptionLongMoney:                   strToDecimal(table1Cells[15]),
				TrustOptionShortLotCount:               strToInt(table1Cells[17]),
				TrustOptionShortMoney:                  strToDecimal(table1Cells[19]),
				TrustOptionLotNetCount:                 strToInt(table1Cells[21]),
				TrustOptionNetMoney:                    strToDecimal(table1Cells[23]),
				ForeignFutureLongLotCount:              strToInt(table1Cells[24]),
				ForeignFutureLongMoney:                 strToDecimal(table1Cells[26]),
				ForeignFutureShortLotCount:             strToInt(table1Cells[28]),
				ForeignFutureShortMoney:                strToDecimal(table1Cells[30]),
				ForeignFutureLotNetCount:               strToInt(table1Cells[32]),
				ForeignFutureNetMoney:                  strToDecimal(table1Cells[34]),
				ForeignOptionLongLotCount:              strToInt(table1Cells[25]),
				ForeignOptionLongMoney:                 strToDecimal(table1Cells[27]),
				ForeignOptionShortLotCount:             strToInt(table1Cells[29]),
				ForeignOptionShortMoney:                strToDecimal(table1Cells[31]),
				ForeignOptionLotNetCount:               strToInt(table1Cells[33]),
				ForeignOptionNetMoney:                  strToDecimal(table1Cells[35]),
				OpenInterestDealerFutureLongLotCount:   strToInt(table2Cells[0]),
				OpenInterestDealerFutureLongMoney:      strToDecimal(table2Cells[2]),
				OpenInterestDealerFutureShortLotCount:  strToInt(table2Cells[4]),
				OpenInterestDealerFutureShortMoney:     strToDecimal(table2Cells[6]),
				OpenInterestDealerFutureLotNetCount:    strToInt(table2Cells[8]),
				OpenInterestDealerFutureNetMoney:       strToDecimal(table2Cells[10]),
				OpenInterestDealerOptionLongLotCount:   strToInt(table2Cells[1]),
				OpenInterestDealerOptionLongMoney:      strToDecimal(table2Cells[3]),
				OpenInterestDealerOptionShortLotCount:  strToInt(table2Cells[5]),
				OpenInterestDealerOptionShortMoney:     strToDecimal(table2Cells[7]),
				OpenInterestDealerOptionLotNetCount:    strToInt(table2Cells[9]),
				OpenInterestDealerOptionNetMoney:       strToDecimal(table2Cells[11]),
				OpenInterestTrustFutureLongLotCount:    strToInt(table2Cells[12]),
				OpenInterestTrustFutureLongMoney:       strToDecimal(table2Cells[14]),
				OpenInterestTrustFutureShortLotCount:   strToInt(table2Cells[16]),
				OpenInterestTrustFutureShortMoney:      strToDecimal(table2Cells[18]),
				OpenInterestTrustFutureLotNetCount:     strToInt(table2Cells[20]),
				OpenInterestTrustFutureNetMoney:        strToDecimal(table2Cells[22]),
				OpenInterestTrustOptionLongLotCount:    strToInt(table2Cells[13]),
				OpenInterestTrustOptionLongMoney:       strToDecimal(table2Cells[15]),
				OpenInterestTrustOptionShortLotCount:   strToInt(table2Cells[17]),
				OpenInterestTrustOptionShortMoney:      strToDecimal(table2Cells[19]),
				OpenInterestTrustOptionLotNetCount:     strToInt(table2Cells[21]),
				OpenInterestTrustOptionNetMoney:        strToDecimal(table2Cells[23]),
				OpenInterestForeignFutureLongLotCount:  strToInt(table2Cells[24]),
				OpenInterestForeignFutureLongMoney:     strToDecimal(table2Cells[26]),
				OpenInterestForeignFutureShortLotCount: strToInt(table2Cells[28]),
				OpenInterestForeignFutureShortMoney:    strToDecimal(table2Cells[30]),
				OpenInterestForeignFutureLotNetCount:   strToInt(table2Cells[32]),
				OpenInterestForeignFutureNetMoney:      strToDecimal(table2Cells[34]),
				OpenInterestForeignOptionLongLotCount:  strToInt(table2Cells[25]),
				OpenInterestForeignOptionLongMoney:     strToDecimal(table2Cells[27]),
				OpenInterestForeignOptionShortLotCount: strToInt(table2Cells[29]),
				OpenInterestForeignOptionShortMoney:    strToDecimal(table2Cells[31]),
				OpenInterestForeignOptionLotNetCount:   strToInt(table2Cells[33]),
				OpenInterestForeignOptionNetMoney:      strToDecimal(table2Cells[35]),
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
