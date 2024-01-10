package module

import (
	"context"
	"fmt"
	"time"
	"tw_option_collector/data"
	"tw_option_collector/objects"

	"github.com/chromedp/chromedp"
)

// 抓取大額交易人期貨未沖銷 TODO:
func CrawlTopTradeFutureOI(ctx context.Context) {
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.taifex.com.tw/cht/3/largeTraderFutQry"))
	handleError(err)
	startDate := time.Date(2023, time.January, 21, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2021, time.January, 4, 0, 0, 0, 0, time.UTC)
	for {
		fmt.Println("[TopFuture] date=", startDate)
		c, err := data.FindTopTradeOIFutureCountByDate(startDate)
		handleError(err)
		if c != 0 {
			fmt.Println("data exist")
			startDate = startDate.AddDate(0, 0, -1)
			if startDate.Before(toDate) {
				break
			}
			continue
		}
		table1Cells := make([]string, 28)
		//先送出日期做查詢
		sendQueryTasks := chromedp.Tasks{chromedp.Sleep(500 * time.Millisecond),
			chromedp.Evaluate(`document.querySelector("#queryDate").value = ""`, nil),
			chromedp.Sleep(100 * time.Millisecond),
			chromedp.SendKeys("#queryDate", startDate.Format("2006/01/02"), chromedp.NodeVisible, chromedp.ByID),
			chromedp.Sleep(300 * time.Millisecond),
			chromedp.Click(`#uForm > table > tbody > tr:nth-child(1) > td > table > tbody > tr:nth-child(1) > td:nth-child(2) > button`, chromedp.ByQuery),
			chromedp.Sleep(300 * time.Millisecond),
			chromedp.Click(`#uForm > table > tbody > tr:nth-child(1) > td > table > tbody > tr:nth-child(1) > td:nth-child(1) > label`, chromedp.ByQuery),
			chromedp.Sleep(500 * time.Millisecond)}

		err = chromedp.Run(ctx, sendQueryTasks)
		handleError(err)
		symbolTypeLen := 0
		err = chromedp.Run(ctx,
			chromedp.Evaluate(`document.querySelector("#contractId").getElementsByTagName("option").length;`, &symbolTypeLen))
		handleError(err)
		fmt.Println("len=", symbolTypeLen)
		if symbolTypeLen < 2 {
			fmt.Println("today no data, holiday?!")
			startDate = startDate.AddDate(0, 0, -1)
			if startDate.Before(toDate) {
				break
			}
			continue
		}
		clickButtonTasks := chromedp.Tasks{
			chromedp.SetValue(`#contractId`, "all", chromedp.ByQuery),
			chromedp.Focus(`#uForm > table > tbody > tr:nth-child(2) > td > #submitButton`),
			chromedp.Click(`#uForm > table > tbody > tr:nth-child(2) > td > #submitButton`, chromedp.ByID),
			chromedp.Sleep(4000 * time.Millisecond)}
		err = chromedp.Run(ctx, clickButtonTasks)
		handleError(err)

		isTableExist := false
		err = chromedp.Run(ctx,
			chromedp.EvaluateAsDevTools(`document.querySelector("#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody") !== null`, &isTableExist))
		handleError(err)
		fmt.Println(isTableExist)
		if isTableExist {
			tasks := chromedp.Tasks{
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(3)`, &table1Cells[0], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(4)`, &table1Cells[1], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(5)`, &table1Cells[2], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(6)`, &table1Cells[3], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(7)`, &table1Cells[4], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(8)`, &table1Cells[5], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(9)`, &table1Cells[6], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(10)`, &table1Cells[7], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(2)`, &table1Cells[8], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(3)`, &table1Cells[9], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(4)`, &table1Cells[10], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(5)`, &table1Cells[11], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(6)`, &table1Cells[12], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(7)`, &table1Cells[13], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(8)`, &table1Cells[14], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(9)`, &table1Cells[15], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(2)`, &table1Cells[16], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(3)`, &table1Cells[17], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(4)`, &table1Cells[18], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(5)`, &table1Cells[19], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(6)`, &table1Cells[20], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(7)`, &table1Cells[21], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(8)`, &table1Cells[22], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(9)`, &table1Cells[23], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(11)`, &table1Cells[24], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(10)`, &table1Cells[25], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(10)`, &table1Cells[26], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(1)`, &table1Cells[27], chromedp.ByQuery)}
			err = chromedp.Run(ctx, tasks)
			handleError(err)
			chromedp.Sleep(1 * time.Second)
			fmt.Println(table1Cells)
			//fmt.Println(table1Cells)

			data.InsertTopTradeOIFuture(&objects.TopFutureOI{
				DataDate:                  startDate,
				WeekTop5BuyLot:            strToIntData(table1Cells[0], false),
				WeekTop5BuyPercent:        strToDecimalData(table1Cells[1], false),
				WeekTop5SpecBuyLot:        strToIntData(table1Cells[0], true),
				WeekTop5SpecBuyPercent:    strToDecimalData(table1Cells[1], true),
				WeekTop5SellLot:           strToIntData(table1Cells[4], false),
				WeekTop5SellPercent:       strToDecimalData(table1Cells[5], false),
				WeekTop5SpecSellLot:       strToIntData(table1Cells[4], true),
				WeekTop5SpecSellPercent:   strToDecimalData(table1Cells[5], true),
				MonthTop5BuyLot:           strToIntData(table1Cells[8], false),
				MonthTop5BuyPercent:       strToDecimalData(table1Cells[9], false),
				MonthTop5SpecBuyLot:       strToIntData(table1Cells[8], true),
				MonthTop5SpecBuyPercent:   strToDecimalData(table1Cells[9], true),
				MonthTop5SellLot:          strToIntData(table1Cells[12], false),
				MonthTop5SellPercent:      strToDecimalData(table1Cells[13], false),
				MonthTop5SpecSellLot:      strToIntData(table1Cells[12], true),
				MonthTop5SpecSellPercent:  strToDecimalData(table1Cells[13], true),
				AllTop5BuyLot:             strToIntData(table1Cells[16], false),
				AllTop5BuyPercent:         strToDecimalData(table1Cells[17], false),
				AllTop5SpecBuyLot:         strToIntData(table1Cells[16], true),
				AllTop5SpecBuyPercent:     strToDecimalData(table1Cells[17], true),
				AllTop5SellLot:            strToIntData(table1Cells[20], false),
				AllTop5SellPercent:        strToDecimalData(table1Cells[21], false),
				AllTop5SpecSellLot:        strToIntData(table1Cells[20], true),
				AllTop5SpecSellPercent:    strToDecimalData(table1Cells[21], true),
				WeekTop10BuyLot:           strToIntData(table1Cells[2], false),
				WeekTop10BuyPercent:       strToDecimalData(table1Cells[3], false),
				WeekTop10SpecBuyLot:       strToIntData(table1Cells[2], true),
				WeekTop10SpecBuyPercent:   strToDecimalData(table1Cells[3], true),
				WeekTop10SellLot:          strToIntData(table1Cells[6], false),
				WeekTop10SellPercent:      strToDecimalData(table1Cells[7], false),
				WeekTop10SpecSellLot:      strToIntData(table1Cells[6], true),
				WeekTop10SpecSellPercent:  strToDecimalData(table1Cells[7], true),
				MonthTop10BuyLot:          strToIntData(table1Cells[10], false),
				MonthTop10BuyPercent:      strToDecimalData(table1Cells[11], false),
				MonthTop10SpecBuyLot:      strToIntData(table1Cells[10], true),
				MonthTop10SpecBuyPercent:  strToDecimalData(table1Cells[11], true),
				MonthTop10SellLot:         strToIntData(table1Cells[14], false),
				MonthTop10SellPercent:     strToDecimalData(table1Cells[15], false),
				MonthTop10SpecSellLot:     strToIntData(table1Cells[14], true),
				MonthTop10SpecSellPercent: strToDecimalData(table1Cells[15], true),
				AllTop10BuyLot:            strToIntData(table1Cells[18], false),
				AllTop10BuyPercent:        strToDecimalData(table1Cells[19], false),
				AllTop10SpecBuyLot:        strToIntData(table1Cells[18], true),
				AllTop10SpecBuyPercent:    strToDecimalData(table1Cells[19], true),
				AllTop10SellLot:           strToIntData(table1Cells[22], false),
				AllTop10SellPercent:       strToDecimalData(table1Cells[23], false),
				AllTop10SpecSellLot:       strToIntData(table1Cells[22], true),
				AllTop10SpecSellPercent:   strToDecimalData(table1Cells[23], true),
				WeekTotalOI:               strToInt(table1Cells[24]),
				MonthTotalOI:              strToInt(table1Cells[25]),
				AllTotalOI:                strToInt(table1Cells[26]),
				MonthContract:             deleteSpace(table1Cells[27]),
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
