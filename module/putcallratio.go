package module

import (
	"context"
	"fmt"
	"time"
	"tw_option_collector/data"
	"tw_option_collector/objects"

	"github.com/chromedp/chromedp"
)

// 抓取PutCall ratio
func CrawlDayPutCallRatio(ctx context.Context) {
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.taifex.com.tw/cht/3/pcRatio"))
	handleError(err)
	startDate := time.Date(2024, time.January, 7, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2021, time.January, 4, 0, 0, 0, 0, time.UTC)
	for {
		daysAgo := 0
		switch startDate.Weekday() {
		case time.Monday:
			daysAgo = -1
		case time.Tuesday:
			daysAgo = -2
		case time.Wednesday:
			daysAgo = -3
		case time.Thursday:
			daysAgo = -4
		case time.Friday:
			daysAgo = -5
		case time.Saturday:
			daysAgo = -6
		case time.Sunday:
			daysAgo = -7
		}
		weekStart := startDate.AddDate(0, 0, daysAgo)

		theEndDate := startDate
		fmt.Println("[PutCallRatio] start=", weekStart, " end=", theEndDate)

		sendQueryTasks := chromedp.Tasks{chromedp.Sleep(500 * time.Millisecond),
			chromedp.Evaluate(`document.querySelector("#queryStartDate").value = ""`, nil),
			chromedp.Evaluate(`document.querySelector("#queryEndDate").value = ""`, nil),
			chromedp.Sleep(100 * time.Millisecond),
			chromedp.SendKeys("#queryStartDate", weekStart.Format("2006/01/02"), chromedp.NodeVisible, chromedp.ByID),
			chromedp.SendKeys("#queryEndDate", theEndDate.Format("2006/01/02"), chromedp.NodeVisible, chromedp.ByID),
			chromedp.Sleep(500 * time.Millisecond)}

		err = chromedp.Run(ctx, sendQueryTasks)
		handleError(err)
		clickButtonTasks := chromedp.Tasks{
			chromedp.Focus(`#button4`),
			chromedp.Click(`#button4`, chromedp.ByID),
			chromedp.Sleep(3500 * time.Millisecond)}
		err = chromedp.Run(ctx, clickButtonTasks)
		handleError(err)

		isTableExist := false
		err = chromedp.Run(ctx,
			chromedp.EvaluateAsDevTools(`document.querySelector("#uForm > div > table.table_a > tbody") !== null`, &isTableExist))
		handleError(err)
		fmt.Println(isTableExist)
		if isTableExist {
			row := 2

			for {
				rowExist := false
				err = chromedp.Run(ctx,
					chromedp.EvaluateAsDevTools(fmt.Sprint(`document.querySelector("#uForm > div > table.table_a > tbody > tr:nth-child(`, row, `) > td:nth-child(1)") !== null`), &rowExist))
				handleError(err)
				if !rowExist {
					break
				}
				row += 1
			}
			for r := 2; r < row; r++ {
				rowCells := make([]string, 7)

				tasks := chromedp.Tasks{
					chromedp.Text(fmt.Sprint(`#uForm > div > table.table_a > tbody > tr:nth-child(`, r, `) > td:nth-child(1)`), &rowCells[0], chromedp.ByQuery),
					chromedp.Text(fmt.Sprint(`#uForm > div > table.table_a > tbody > tr:nth-child(`, r, `) > td:nth-child(2)`), &rowCells[1], chromedp.ByQuery),
					chromedp.Text(fmt.Sprint(`#uForm > div > table.table_a > tbody > tr:nth-child(`, r, `) > td:nth-child(3)`), &rowCells[2], chromedp.ByQuery),
					chromedp.Text(fmt.Sprint(`#uForm > div > table.table_a > tbody > tr:nth-child(`, r, `) > td:nth-child(4)`), &rowCells[3], chromedp.ByQuery),
					chromedp.Text(fmt.Sprint(`#uForm > div > table.table_a > tbody > tr:nth-child(`, r, `) > td:nth-child(5)`), &rowCells[4], chromedp.ByQuery),
					chromedp.Text(fmt.Sprint(`#uForm > div > table.table_a > tbody > tr:nth-child(`, r, `) > td:nth-child(6)`), &rowCells[5], chromedp.ByQuery),
					chromedp.Text(fmt.Sprint(`#uForm > div > table.table_a > tbody > tr:nth-child(`, r, `) > td:nth-child(7)`), &rowCells[6], chromedp.ByQuery)}
				err = chromedp.Run(ctx, tasks)
				handleError(err)

				fmt.Println(rowCells)
				theDate, err := parseDateString(rowCells[0])
				handleError(err)
				c, err := data.FindPutCallRatioCountByDate(theDate)
				handleError(err)
				if c != 0 {
					fmt.Println("data exist")
					continue
				}

				data.InsertPutCallRatio(&objects.MarketPutCallRatio{
					DataDate:          theDate,
					PutTradeVolume:    strToInt(rowCells[1]),
					CallTradeVolume:   strToInt(rowCells[2]),
					PutCallTradeRatio: strToDecimal(rowCells[3]),
					OIPutVolume:       strToInt(rowCells[4]),
					OICallVolume:      strToInt(rowCells[5]),
					PutCallOIRatio:    strToDecimal(rowCells[6]),
				})
			}

		} else {
			chromedp.Sleep(2500 * time.Millisecond)
			startDate = weekStart
			if startDate.Before(toDate) {
				break
			}
		}

		chromedp.Sleep(2500 * time.Millisecond)
		startDate = weekStart
		if startDate.Before(toDate) {
			break
		}
	}

}

func parseDateString(str string) (time.Time, error) {
	strLen := len(str)
	switch strLen {
	case 8:
		return time.Parse("2006/1/2", str)
	case 9:
		if string(str[6]) == "/" {
			//2006/1/02
			return time.Parse("2006/1/02", str)
		} else {
			//2006/01/2
			return time.Parse("2006/01/2", str)
		}

	default:
		return time.Parse("2006/01/02", str)
	}
}
