package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
	"tw_option_collector/configure"
	"tw_option_collector/env"
	"tw_option_collector/module"

	"github.com/chromedp/chromedp"
)

func main() {
	//load configuration
	configure.LoadConfiguration()

	// command
	// env init
	env.Initialize()

	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false), chromedp.Flag("start-maximized", true))
	newCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	//defer cancel()

	_, cancel := chromedp.NewContext(newCtx, chromedp.WithLogf(log.Printf))
	exitChan := make(chan int, 1)
	go func() {
		popCLI()
		exitChan <- 1
	}()
	//module.CrawlDayPutCallRatio(ctx)
	//module.CrawlTopTradeFutureOI(ctx)
	//module.CrawlTopTradeOptionOI(ctx)

	//module.CrawlBigThreeFuture(ctx)
	//module.CrawlBigThreeOption(ctx)
	//module.CrawlBigThreeCallPut(ctx)
	//module.CrawlBigThreeFutureOption(ctx)

	//module.CrawlBigThreeSummary(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	<-exitChan
	fmt.Println("program exit")
}

func popCLI() {
	for {
		var cmdInput string
		fmt.Println("Input Command:[A]/[B]/[Exit]")
		fmt.Scan(&cmdInput)
		if strings.ToLower(cmdInput) == "exit" {
			break
		}
		switch strings.ToLower(cmdInput) {
		case "a", "A":
			doAnalyzeCommand()
		case "b", "B":
			doAnalyzePutCallRatio()
		default:
			fmt.Println("invalid command!")
		}
	}
	fmt.Println("exit cli !")
}

func doAnalyzeCommand() {
	var dateStr string
	fmt.Println("Input Date:yyyy/MM/dd")
	fmt.Scan(&dateStr)
	dt, err := time.Parse("2006/01/02", dateStr)
	handleError(err)
	module.AnalyzeSummary(dt)

}

func doAnalyzePutCallRatio() {
	module.AnalyzePutCallRatio()
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
