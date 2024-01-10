package module

import (
	"context"
	"fmt"
	"time"
	"tw_option_collector/data"
	"tw_option_collector/objects"

	"github.com/chromedp/chromedp"
)

// 抓取三大法人買賣權分計
func CrawlBigThreeCallPut(ctx context.Context) {
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.taifex.com.tw/cht/3/callsAndPutsDate"))
	handleError(err)
	startDate := time.Date(2024, time.January, 3, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2021, time.January, 4, 0, 0, 0, 0, time.UTC)
	for {
		fmt.Println("[CallPut] date=", startDate)
		c, err := data.FindBigThreeCallPutCountByDate(startDate)
		handleError(err)
		if c != 0 {
			fmt.Println("data exist")
			startDate = startDate.AddDate(0, 0, -1)
			if startDate.Before(toDate) {
				break
			}
			continue
		}
		table1Cells := make([]string, 72)
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

		if isTableExist {
			tasks := chromedp.Tasks{
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(5)`, &table1Cells[0], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(6)`, &table1Cells[1], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(7)`, &table1Cells[2], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(8)`, &table1Cells[3], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(9)`, &table1Cells[4], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(10)`, &table1Cells[5], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(11)`, &table1Cells[6], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(12)`, &table1Cells[7], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(13)`, &table1Cells[8], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(14)`, &table1Cells[9], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(15)`, &table1Cells[10], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(16)`, &table1Cells[11], chromedp.ByQuery),
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
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(13)`, &table1Cells[35], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(3)`, &table1Cells[36], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(4)`, &table1Cells[37], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(5)`, &table1Cells[38], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(6)`, &table1Cells[39], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(7)`, &table1Cells[40], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(8)`, &table1Cells[41], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(9)`, &table1Cells[42], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(10)`, &table1Cells[43], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(11)`, &table1Cells[44], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(12)`, &table1Cells[45], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(13)`, &table1Cells[46], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(7) > td:nth-child(14)`, &table1Cells[47], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(2)`, &table1Cells[48], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(3)`, &table1Cells[49], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(4)`, &table1Cells[50], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(5)`, &table1Cells[51], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(6)`, &table1Cells[52], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(7)`, &table1Cells[53], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(8)`, &table1Cells[54], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(9)`, &table1Cells[55], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(10)`, &table1Cells[56], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(11)`, &table1Cells[57], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(12)`, &table1Cells[58], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(8) > td:nth-child(13)`, &table1Cells[59], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(2)`, &table1Cells[60], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(3)`, &table1Cells[61], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(4)`, &table1Cells[62], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(5)`, &table1Cells[63], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(6)`, &table1Cells[64], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(7)`, &table1Cells[65], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(8)`, &table1Cells[66], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(9)`, &table1Cells[67], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(10)`, &table1Cells[68], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(11)`, &table1Cells[69], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(12)`, &table1Cells[70], chromedp.ByQuery),
				chromedp.Text(`#printhere > div:nth-child(4) > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(9) > td:nth-child(13)`, &table1Cells[71], chromedp.ByQuery)}
			err = chromedp.Run(ctx, tasks)
			handleError(err)
			chromedp.Sleep(1 * time.Second)
			fmt.Println(table1Cells)
			//fmt.Println(table1Cells)

			data.InsertBigThreeCallPut(&objects.BigThreeCallPut{
				DataDate:                  startDate,
				CallDealerBuyLotCount:     strToInt(table1Cells[0]),
				CallDealerBuyMoney:        strToDecimal(table1Cells[1]),
				CallDealerSellLotCount:    strToInt(table1Cells[2]),
				CallDealerSellMoney:       strToDecimal(table1Cells[3]),
				CallDealerDiffLotCount:    strToInt(table1Cells[4]),
				CallDealerDiffMoney:       strToDecimal(table1Cells[5]),
				CallDealerOIBuyLotCount:   strToInt(table1Cells[6]),
				CallDealerOIBuyMoney:      strToDecimal(table1Cells[7]),
				CallDealerOISellLotCount:  strToInt(table1Cells[8]),
				CallDealerOISellMoney:     strToDecimal(table1Cells[9]),
				CallDealerOIDiffLotCount:  strToInt(table1Cells[10]),
				CallDealerOIDiffMoney:     strToDecimal(table1Cells[11]),
				CallTrustBuyLotCount:      strToInt(table1Cells[12]),
				CallTrustBuyMoney:         strToDecimal(table1Cells[13]),
				CallTrustSellLotCount:     strToInt(table1Cells[14]),
				CallTrustSellMoney:        strToDecimal(table1Cells[15]),
				CallTrustDiffLotCount:     strToInt(table1Cells[16]),
				CallTrustDiffMoney:        strToDecimal(table1Cells[17]),
				CallTrustOIBuyLotCount:    strToInt(table1Cells[18]),
				CallTrustOIBuyMoney:       strToDecimal(table1Cells[19]),
				CallTrustOISellLotCount:   strToInt(table1Cells[20]),
				CallTrustOISellMoney:      strToDecimal(table1Cells[21]),
				CallTrustOIDiffLotCount:   strToInt(table1Cells[22]),
				CallTrustOIDiffMoney:      strToDecimal(table1Cells[23]),
				CallForeignBuyLotCount:    strToInt(table1Cells[24]),
				CallForeignBuyMoney:       strToDecimal(table1Cells[25]),
				CallForeignSellLotCount:   strToInt(table1Cells[26]),
				CallForeignSellMoney:      strToDecimal(table1Cells[27]),
				CallForeignDiffLotCount:   strToInt(table1Cells[28]),
				CallForeignDiffMoney:      strToDecimal(table1Cells[29]),
				CallForeignOIBuyLotCount:  strToInt(table1Cells[30]),
				CallForeignOIBuyMoney:     strToDecimal(table1Cells[31]),
				CallForeignOISellLotCount: strToInt(table1Cells[32]),
				CallForeignOISellMoney:    strToDecimal(table1Cells[33]),
				CallForeignOIDiffLotCount: strToInt(table1Cells[34]),
				CallForeignOIDiffMoney:    strToDecimal(table1Cells[35]),
				PutDealerBuyLotCount:      strToInt(table1Cells[36]),
				PutDealerBuyMoney:         strToDecimal(table1Cells[37]),
				PutDealerSellLotCount:     strToInt(table1Cells[38]),
				PutDealerSellMoney:        strToDecimal(table1Cells[39]),
				PutDealerDiffLotCount:     strToInt(table1Cells[40]),
				PutDealerDiffMoney:        strToDecimal(table1Cells[41]),
				PutDealerOIBuyLotCount:    strToInt(table1Cells[42]),
				PutDealerOIBuyMoney:       strToDecimal(table1Cells[43]),
				PutDealerOISellLotCount:   strToInt(table1Cells[44]),
				PutDealerOISellMoney:      strToDecimal(table1Cells[45]),
				PutDealerOIDiffLotCount:   strToInt(table1Cells[46]),
				PutDealerOIDiffMoney:      strToDecimal(table1Cells[47]),
				PutTrustBuyLotCount:       strToInt(table1Cells[48]),
				PutTrustBuyMoney:          strToDecimal(table1Cells[49]),
				PutTrustSellLotCount:      strToInt(table1Cells[50]),
				PutTrustSellMoney:         strToDecimal(table1Cells[51]),
				PutTrustDiffLotCount:      strToInt(table1Cells[52]),
				PutTrustDiffMoney:         strToDecimal(table1Cells[53]),
				PutTrustOIBuyLotCount:     strToInt(table1Cells[54]),
				PutTrustOIBuyMoney:        strToDecimal(table1Cells[55]),
				PutTrustOISellLotCount:    strToInt(table1Cells[56]),
				PutTrustOISellMoney:       strToDecimal(table1Cells[57]),
				PutTrustOIDiffLotCount:    strToInt(table1Cells[58]),
				PutTrustOIDiffMoney:       strToDecimal(table1Cells[59]),
				PutForeignBuyLotCount:     strToInt(table1Cells[60]),
				PutForeignBuyMoney:        strToDecimal(table1Cells[61]),
				PutForeignSellLotCount:    strToInt(table1Cells[62]),
				PutForeignSellMoney:       strToDecimal(table1Cells[63]),
				PutForeignDiffLotCount:    strToInt(table1Cells[64]),
				PutForeignDiffMoney:       strToDecimal(table1Cells[65]),
				PutForeignOIBuyLotCount:   strToInt(table1Cells[66]),
				PutForeignOIBuyMoney:      strToDecimal(table1Cells[67]),
				PutForeignOISellLotCount:  strToInt(table1Cells[68]),
				PutForeignOISellMoney:     strToDecimal(table1Cells[69]),
				PutForeignOIDiffLotCount:  strToInt(table1Cells[70]),
				PutForeignOIDiffMoney:     strToDecimal(table1Cells[71]),
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
