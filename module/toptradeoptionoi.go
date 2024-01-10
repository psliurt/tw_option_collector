package module

import (
	"context"
	"fmt"
	"time"
	"tw_option_collector/data"
	"tw_option_collector/objects"

	"github.com/chromedp/chromedp"
)

// 抓取大額交易人選擇權未沖銷 TODO:
func CrawlTopTradeOptionOI(ctx context.Context) {
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.taifex.com.tw/cht/3/largeTraderOptQry"))
	handleError(err)
	startDate := time.Date(2021, time.February, 3, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2021, time.January, 5, 0, 0, 0, 0, time.UTC)
	for {
		fmt.Println("[TopOption] date=", startDate)
		c, err := data.FindTopTradeOIOptionCountByDate(startDate)
		handleError(err)
		if c != 0 {
			fmt.Println("data exist")
			startDate = startDate.AddDate(0, 0, -1)
			if startDate.Before(toDate) {
				break
			}
			continue
		}
		table1Cells := make([]string, 56)
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
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(3)`, &table1Cells[24], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(4)`, &table1Cells[25], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(5)`, &table1Cells[26], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(6)`, &table1Cells[27], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(7)`, &table1Cells[28], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(8)`, &table1Cells[29], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(9)`, &table1Cells[30], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(10)`, &table1Cells[31], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(2)`, &table1Cells[32], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(3)`, &table1Cells[33], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(4)`, &table1Cells[34], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(5)`, &table1Cells[35], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(6)`, &table1Cells[36], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(7)`, &table1Cells[37], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(8)`, &table1Cells[38], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(9)`, &table1Cells[39], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(2)`, &table1Cells[40], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(3)`, &table1Cells[41], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(4)`, &table1Cells[42], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(5)`, &table1Cells[43], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(6)`, &table1Cells[44], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(7)`, &table1Cells[45], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(8)`, &table1Cells[46], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(9)`, &table1Cells[47], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(11)`, &table1Cells[48], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(10)`, &table1Cells[49], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(10)`, &table1Cells[50], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(5) > td:nth-child(1)`, &table1Cells[51], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(11)`, &table1Cells[52], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(10)`, &table1Cells[53], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(10)`, &table1Cells[54], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(3) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(1)`, &table1Cells[55], chromedp.ByQuery)}
			err = chromedp.Run(ctx, tasks)
			handleError(err)
			chromedp.Sleep(1 * time.Second)
			fmt.Println(table1Cells)
			//fmt.Println(table1Cells)

			data.InsertTopTradeOIOption(&objects.TopOptionOI{
				DataDate:                      startDate,
				CallWeekTop5BuyLot:            strToIntData(table1Cells[0], false),
				CallWeekTop5BuyPercent:        strToDecimalData(table1Cells[1], false),
				CallWeekTop5SpecBuyLot:        strToIntData(table1Cells[0], true),
				CallWeekTop5SpecBuyPercent:    strToDecimalData(table1Cells[1], true),
				CallWeekTop5SellLot:           strToIntData(table1Cells[4], false),
				CallWeekTop5SellPercent:       strToDecimalData(table1Cells[5], false),
				CallWeekTop5SpecSellLot:       strToIntData(table1Cells[4], true),
				CallWeekTop5SpecSellPercent:   strToDecimalData(table1Cells[5], true),
				CallMonthTop5BuyLot:           strToIntData(table1Cells[8], false),
				CallMonthTop5BuyPercent:       strToDecimalData(table1Cells[9], false),
				CallMonthTop5SpecBuyLot:       strToIntData(table1Cells[8], true),
				CallMonthTop5SpecBuyPercent:   strToDecimalData(table1Cells[9], true),
				CallMonthTop5SellLot:          strToIntData(table1Cells[12], false),
				CallMonthTop5SellPercent:      strToDecimalData(table1Cells[13], false),
				CallMonthTop5SpecSellLot:      strToIntData(table1Cells[12], true),
				CallMonthTop5SpecSellPercent:  strToDecimalData(table1Cells[13], true),
				CallAllTop5BuyLot:             strToIntData(table1Cells[16], false),
				CallAllTop5BuyPercent:         strToDecimalData(table1Cells[17], false),
				CallAllTop5SpecBuyLot:         strToIntData(table1Cells[16], true),
				CallAllTop5SpecBuyPercent:     strToDecimalData(table1Cells[17], true),
				CallAllTop5SellLot:            strToIntData(table1Cells[20], false),
				CallAllTop5SellPercent:        strToDecimalData(table1Cells[21], false),
				CallAllTop5SpecSellLot:        strToIntData(table1Cells[20], true),
				CallAllTop5SpecSellPercent:    strToDecimalData(table1Cells[21], true),
				CallWeekTop10BuyLot:           strToIntData(table1Cells[2], false),
				CallWeekTop10BuyPercent:       strToDecimalData(table1Cells[3], false),
				CallWeekTop10SpecBuyLot:       strToIntData(table1Cells[2], true),
				CallWeekTop10SpecBuyPercent:   strToDecimalData(table1Cells[3], true),
				CallWeekTop10SellLot:          strToIntData(table1Cells[6], false),
				CallWeekTop10SellPercent:      strToDecimalData(table1Cells[7], false),
				CallWeekTop10SpecSellLot:      strToIntData(table1Cells[6], true),
				CallWeekTop10SpecSellPercent:  strToDecimalData(table1Cells[7], true),
				CallMonthTop10BuyLot:          strToIntData(table1Cells[10], false),
				CallMonthTop10BuyPercent:      strToDecimalData(table1Cells[11], false),
				CallMonthTop10SpecBuyLot:      strToIntData(table1Cells[10], true),
				CallMonthTop10SpecBuyPercent:  strToDecimalData(table1Cells[11], true),
				CallMonthTop10SellLot:         strToIntData(table1Cells[14], false),
				CallMonthTop10SellPercent:     strToDecimalData(table1Cells[15], false),
				CallMonthTop10SpecSellLot:     strToIntData(table1Cells[14], true),
				CallMonthTop10SpecSellPercent: strToDecimalData(table1Cells[15], true),
				CallAllTop10BuyLot:            strToIntData(table1Cells[18], false),
				CallAllTop10BuyPercent:        strToDecimalData(table1Cells[19], false),
				CallAllTop10SpecBuyLot:        strToIntData(table1Cells[18], true),
				CallAllTop10SpecBuyPercent:    strToDecimalData(table1Cells[19], true),
				CallAllTop10SellLot:           strToIntData(table1Cells[22], false),
				CallAllTop10SellPercent:       strToDecimalData(table1Cells[23], false),
				CallAllTop10SpecSellLot:       strToIntData(table1Cells[22], true),
				CallAllTop10SpecSellPercent:   strToDecimalData(table1Cells[23], true),
				PutWeekTop5BuyLot:             strToIntData(table1Cells[24], false),
				PutWeekTop5BuyPercent:         strToDecimalData(table1Cells[25], false),
				PutWeekTop5SpecBuyLot:         strToIntData(table1Cells[24], true),
				PutWeekTop5SpecBuyPercent:     strToDecimalData(table1Cells[25], true),
				PutWeekTop5SellLot:            strToIntData(table1Cells[28], false),
				PutWeekTop5SellPercent:        strToDecimalData(table1Cells[29], false),
				PutWeekTop5SpecSellLot:        strToIntData(table1Cells[28], true),
				PutWeekTop5SpecSellPercent:    strToDecimalData(table1Cells[29], true),
				PutMonthTop5BuyLot:            strToIntData(table1Cells[32], false),
				PutMonthTop5BuyPercent:        strToDecimalData(table1Cells[33], false),
				PutMonthTop5SpecBuyLot:        strToIntData(table1Cells[32], true),
				PutMonthTop5SpecBuyPercent:    strToDecimalData(table1Cells[33], true),
				PutMonthTop5SellLot:           strToIntData(table1Cells[36], false),
				PutMonthTop5SellPercent:       strToDecimalData(table1Cells[37], false),
				PutMonthTop5SpecSellLot:       strToIntData(table1Cells[36], true),
				PutMonthTop5SpecSellPercent:   strToDecimalData(table1Cells[37], true),
				PutAllTop5BuyLot:              strToIntData(table1Cells[40], false),
				PutAllTop5BuyPercent:          strToDecimalData(table1Cells[41], false),
				PutAllTop5SpecBuyLot:          strToIntData(table1Cells[40], true),
				PutAllTop5SpecBuyPercent:      strToDecimalData(table1Cells[41], true),
				PutAllTop5SellLot:             strToIntData(table1Cells[44], false),
				PutAllTop5SellPercent:         strToDecimalData(table1Cells[45], false),
				PutAllTop5SpecSellLot:         strToIntData(table1Cells[44], true),
				PutAllTop5SpecSellPercent:     strToDecimalData(table1Cells[45], true),
				PutWeekTop10BuyLot:            strToIntData(table1Cells[26], false),
				PutWeekTop10BuyPercent:        strToDecimalData(table1Cells[27], false),
				PutWeekTop10SpecBuyLot:        strToIntData(table1Cells[26], true),
				PutWeekTop10SpecBuyPercent:    strToDecimalData(table1Cells[27], true),
				PutWeekTop10SellLot:           strToIntData(table1Cells[30], false),
				PutWeekTop10SellPercent:       strToDecimalData(table1Cells[31], false),
				PutWeekTop10SpecSellLot:       strToIntData(table1Cells[30], true),
				PutWeekTop10SpecSellPercent:   strToDecimalData(table1Cells[31], true),
				PutMonthTop10BuyLot:           strToIntData(table1Cells[34], false),
				PutMonthTop10BuyPercent:       strToDecimalData(table1Cells[35], false),
				PutMonthTop10SpecBuyLot:       strToIntData(table1Cells[34], true),
				PutMonthTop10SpecBuyPercent:   strToDecimalData(table1Cells[35], true),
				PutMonthTop10SellLot:          strToIntData(table1Cells[38], false),
				PutMonthTop10SellPercent:      strToDecimalData(table1Cells[39], false),
				PutMonthTop10SpecSellLot:      strToIntData(table1Cells[38], true),
				PutMonthTop10SpecSellPercent:  strToDecimalData(table1Cells[39], true),
				PutAllTop10BuyLot:             strToIntData(table1Cells[42], false),
				PutAllTop10BuyPercent:         strToDecimalData(table1Cells[43], false),
				PutAllTop10SpecBuyLot:         strToIntData(table1Cells[42], true),
				PutAllTop10SpecBuyPercent:     strToDecimalData(table1Cells[43], true),
				PutAllTop10SellLot:            strToIntData(table1Cells[46], false),
				PutAllTop10SellPercent:        strToDecimalData(table1Cells[47], false),
				PutAllTop10SpecSellLot:        strToIntData(table1Cells[46], true),
				PutAllTop10SpecSellPercent:    strToDecimalData(table1Cells[47], true),
				CallWeekTotalOI:               strToInt(table1Cells[48]),
				CallMonthTotalOI:              strToInt(table1Cells[49]),
				CallAllTotalOI:                strToInt(table1Cells[50]),
				CallMonthContract:             deleteSpace(table1Cells[51]),
				PutWeekTotalOI:                strToInt(table1Cells[52]),
				PutMonthTotalOI:               strToInt(table1Cells[53]),
				PutAllTotalOI:                 strToInt(table1Cells[54]),
				PutMonthContract:              deleteSpace(table1Cells[55]),
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
