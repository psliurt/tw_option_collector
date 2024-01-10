package objects

import (
	"time"

	"github.com/shopspring/decimal"
)

// 三大法人 總表
type BigThreeSummary struct {
	DataDate                         time.Time       `bson:"datadate"`                         //資料日期
	DealerLongLotCount               int             `bson:"dealerlonglotcount"`               //自營商 多方 交易口數
	DealerLongMoney                  decimal.Decimal `bson:"dealerlongmoney"`                  //自營商 多方 契約金額
	DealerShortLotCount              int             `bson:"dealershortlotcount"`              //自營商 空方 交易口數
	DealerShortMoney                 decimal.Decimal `bson:"dealershortmoney"`                 //自營商 空方 契約金額
	DealerLotNetCount                int             `bson:"dealerlotnetcount"`                //自營商 淨額 交易口數
	DealerNetMoney                   decimal.Decimal `bson:"dealernetmoney"`                   //自營商 淨額 契約金額
	TrustLongLotCount                int             `bson:"trustlonglotcount"`                //投信 多方 交易口數
	TrustLongMoney                   decimal.Decimal `bson:"trustlongmoney"`                   //投信 多方 契約金額
	TrustShortLotCount               int             `bson:"trustshortlotcount"`               //投信 空方 交易口數
	TrustShortMoney                  decimal.Decimal `bson:"trustshortmoney"`                  //投信 空方 契約金額
	TrustLotNetCount                 int             `bson:"trustlotnetcount"`                 //投信 淨額 交易口數
	TrustNetMoney                    decimal.Decimal `bson:"trustnetmoney"`                    //投信 淨額 契約金額
	ForeignLongLotCount              int             `bson:"foreignlonglotcount"`              //外資 多方 交易口數
	ForeignLongMoney                 decimal.Decimal `bson:"foreignlongmoney"`                 //外資 多方 契約金額
	ForeignShortLotCount             int             `bson:"foreignshortlotcount"`             //外資 空方 交易口數
	ForeignShortMoney                decimal.Decimal `bson:"foreignshortmoney"`                //外資 空方 契約金額
	ForeignLotNetCount               int             `bson:"foreignlotnetcount"`               //外資 淨額 交易口數
	ForeignNetMoney                  decimal.Decimal `bson:"foreignnetmoney"`                  //外資 淨額 契約金額
	OpenInterestDealerLongLotCount   int             `bson:"openinterestdealerlonglotcount"`   //自營商 多方 未平倉口數
	OpenInterestDealerLongMoney      decimal.Decimal `bson:"openinterestdealerlongmoney"`      //自營商 多方 契約金額
	OpenInterestDealerShortLotCount  int             `bson:"openinterestdealershortlotcount"`  //自營商 空方 未平倉口數
	OpenInterestDealerShortMoney     decimal.Decimal `bson:"openinterestdealershortmoney"`     //自營商 空方 契約金額
	OpenInterestDealerLotNetCount    int             `bson:"openinterestdealerlotnetcount"`    //自營商 淨額 未平倉口數
	OpenInterestDealerNetMoney       decimal.Decimal `bson:"openinterestdealernetmoney"`       //自營商 淨額 契約金額
	OpenInterestTrustLongLotCount    int             `bson:"openinteresttrustlonglotcount"`    //投信 多方 未平倉口數
	OpenInterestTrustLongMoney       decimal.Decimal `bson:"openinteresttrustlongmoney"`       //投信 多方 契約金額
	OpenInterestTrustShortLotCount   int             `bson:"openinteresttrustshortlotcount"`   //投信 空方 未平倉口數
	OpenInterestTrustShortMoney      decimal.Decimal `bson:"openinteresttrustshortmoney"`      //投信 空方 契約金額
	OpenInterestTrustLotNetCount     int             `bson:"openinteresttrustlotnetcount"`     //投信 淨額 未平倉口數
	OpenInterestTrustNetMoney        decimal.Decimal `bson:"openinteresttrustnetmoney"`        //投信 淨額 契約金額
	OpenInterestForeignLongLotCount  int             `bson:"openinterestforeignlonglotcount"`  //外資 多方 未平倉口數
	OpenInterestForeignLongMoney     decimal.Decimal `bson:"openinterestforeignlongmoney"`     //外資 多方 契約金額
	OpenInterestForeignShortLotCount int             `bson:"openinterestforeignshortlotcount"` //外資 空方 未平倉口數
	OpenInterestForeignShortMoney    decimal.Decimal `bson:"openinterestforeignshortmoney"`    //外資 空方 契約金額
	OpenInterestForeignLotNetCount   int             `bson:"openinterestforeignlotnetcount"`   //外資 淨額 未平倉口數
	OpenInterestForeignNetMoney      decimal.Decimal `bson:"openinterestforeignnetmoney"`      //外資 淨額 契約金額
}

// 三大法人 區分期貨與選擇權二類
type BigThreeFutureOption struct {
	DataDate                               time.Time       `bson:"datadate"`
	DealerFutureLongLotCount               int             `bson:"dealerfuturelonglotcount"`
	DealerFutureLongMoney                  decimal.Decimal `bson:"dealerfuturelongmoney"`
	DealerFutureShortLotCount              int             `bson:"dealerfutureshortlotcount"`
	DealerFutureShortMoney                 decimal.Decimal `bson:"dealerfutureshortmoney"`
	DealerFutureLotNetCount                int             `bson:"dealerfuturelotnetcount"`
	DealerFutureNetMoney                   decimal.Decimal `bson:"dealerfuturenetmoney"`
	DealerOptionLongLotCount               int             `bson:"dealeroptionlonglotcount"`
	DealerOptionLongMoney                  decimal.Decimal `bson:"dealeroptionlongmoney"`
	DealerOptionShortLotCount              int             `bson:"dealeroptionshortlotcount"`
	DealerOptionShortMoney                 decimal.Decimal `bson:"dealeroptionshortmoney"`
	DealerOptionLotNetCount                int             `bson:"dealeroptionlotnetcount"`
	DealerOptionNetMoney                   decimal.Decimal `bson:"dealeroptionnetmoney"`
	TrustFutureLongLotCount                int             `bson:"trustfuturelonglotcount"`
	TrustFutureLongMoney                   decimal.Decimal `bson:"trustfuturelongmoney"`
	TrustFutureShortLotCount               int             `bson:"trustfutureshortlotcount"`
	TrustFutureShortMoney                  decimal.Decimal `bson:"trustfutureshortmoney"`
	TrustFutureLotNetCount                 int             `bson:"trustfuturelotnetcount"`
	TrustFutureNetMoney                    decimal.Decimal `bson:"trustfuturenetmoney"`
	TrustOptionLongLotCount                int             `bson:"trustoptionlonglotcount"`
	TrustOptionLongMoney                   decimal.Decimal `bson:"trustoptionlongmoney"`
	TrustOptionShortLotCount               int             `bson:"trustoptionshortlotcount"`
	TrustOptionShortMoney                  decimal.Decimal `bson:"trustoptionshortmoney"`
	TrustOptionLotNetCount                 int             `bson:"trustoptionlotnetcount"`
	TrustOptionNetMoney                    decimal.Decimal `bson:"trustoptionnetmoney"`
	ForeignFutureLongLotCount              int             `bson:"foreignfuturelonglotcount"`
	ForeignFutureLongMoney                 decimal.Decimal `bson:"foreignfuturelongmoney"`
	ForeignFutureShortLotCount             int             `bson:"foreignfutureshortlotcount"`
	ForeignFutureShortMoney                decimal.Decimal `bson:"foreignfutureshortmoney"`
	ForeignFutureLotNetCount               int             `bson:"foreignfuturelotnetcount"`
	ForeignFutureNetMoney                  decimal.Decimal `bson:"foreignfuturenetmoney"`
	ForeignOptionLongLotCount              int             `bson:"foreignoptionlonglotcount"`
	ForeignOptionLongMoney                 decimal.Decimal `bson:"foreignoptionlongmoney"`
	ForeignOptionShortLotCount             int             `bson:"foreignoptionshortlotcount"`
	ForeignOptionShortMoney                decimal.Decimal `bson:"foreignoptionshortmoney"`
	ForeignOptionLotNetCount               int             `bson:"foreignoptionlotnetcount"`
	ForeignOptionNetMoney                  decimal.Decimal `bson:"foreignoptionnetmoney"`
	OpenInterestDealerFutureLongLotCount   int             `bson:"openinterestdealerfuturelonglotcount"`
	OpenInterestDealerFutureLongMoney      decimal.Decimal `bson:"openinterestdealerfuturelongmoney"`
	OpenInterestDealerFutureShortLotCount  int             `bson:"openinterestdealerfutureshortlotcount"`
	OpenInterestDealerFutureShortMoney     decimal.Decimal `bson:"openinterestdealerfutureshortmoney"`
	OpenInterestDealerFutureLotNetCount    int             `bson:"openinterestdealerfuturelotnetcount"`
	OpenInterestDealerFutureNetMoney       decimal.Decimal `bson:"openinterestdealerfuturenetmoney"`
	OpenInterestDealerOptionLongLotCount   int             `bson:"openinterestdealeroptionlonglotcount"`
	OpenInterestDealerOptionLongMoney      decimal.Decimal `bson:"openinterestdealeroptionlongmoney"`
	OpenInterestDealerOptionShortLotCount  int             `bson:"openinterestdealeroptionshortlotcount"`
	OpenInterestDealerOptionShortMoney     decimal.Decimal `bson:"openinterestdealeroptionshortmoney"`
	OpenInterestDealerOptionLotNetCount    int             `bson:"openinterestdealeroptionlotnetcount"`
	OpenInterestDealerOptionNetMoney       decimal.Decimal `bson:"openinterestdealeroptionnetmoney"`
	OpenInterestTrustFutureLongLotCount    int             `bson:"openinteresttrustfuturelonglotcount"`
	OpenInterestTrustFutureLongMoney       decimal.Decimal `bson:"openinteresttrustfuturelongmoney"`
	OpenInterestTrustFutureShortLotCount   int             `bson:"openinteresttrustfutureshortlotcount"`
	OpenInterestTrustFutureShortMoney      decimal.Decimal `bson:"openinteresttrustfutureshortmoney"`
	OpenInterestTrustFutureLotNetCount     int             `bson:"openinteresttrustfuturelotnetcount"`
	OpenInterestTrustFutureNetMoney        decimal.Decimal `bson:"openinteresttrustfuturenetmoney"`
	OpenInterestTrustOptionLongLotCount    int             `bson:"openinteresttrustoptionlonglotcount"`
	OpenInterestTrustOptionLongMoney       decimal.Decimal `bson:"openinteresttrustoptionlongmoney"`
	OpenInterestTrustOptionShortLotCount   int             `bson:"openinteresttrustoptionshortlotcount"`
	OpenInterestTrustOptionShortMoney      decimal.Decimal `bson:"openinteresttrustoptionshortmoney"`
	OpenInterestTrustOptionLotNetCount     int             `bson:"openinteresttrustoptionlotnetcount"`
	OpenInterestTrustOptionNetMoney        decimal.Decimal `bson:"openinteresttrustoptionnetmoney"`
	OpenInterestForeignFutureLongLotCount  int             `bson:"openinterestforeignfuturelonglotcount"`
	OpenInterestForeignFutureLongMoney     decimal.Decimal `bson:"openinterestforeignfuturelongmoney"`
	OpenInterestForeignFutureShortLotCount int             `bson:"openinterestforeignfutureshortlotcount"`
	OpenInterestForeignFutureShortMoney    decimal.Decimal `bson:"openinterestforeignfutureshortmoney"`
	OpenInterestForeignFutureLotNetCount   int             `bson:"openinterestforeignfuturelotnetcount"`
	OpenInterestForeignFutureNetMoney      decimal.Decimal `bson:"openinterestforeignfuturenetmoney"`
	OpenInterestForeignOptionLongLotCount  int             `bson:"openinterestforeignoptionlonglotcount"`
	OpenInterestForeignOptionLongMoney     decimal.Decimal `bson:"openinterestforeignoptionlongmoney"`
	OpenInterestForeignOptionShortLotCount int             `bson:"openinterestforeignoptionshortlotcount"`
	OpenInterestForeignOptionShortMoney    decimal.Decimal `bson:"openinterestforeignoptionshortmoney"`
	OpenInterestForeignOptionLotNetCount   int             `bson:"openinterestforeignoptionlotnetcount"`
	OpenInterestForeignOptionNetMoney      decimal.Decimal `bson:"openinterestforeignoptionnetmoney"`
}

type BigThreeCallPut struct {
	DataDate                  time.Time       `bson:"datadate"`
	CallDealerBuyLotCount     int             `bson:"calldealerbuylotcount"`
	CallDealerBuyMoney        decimal.Decimal `bson:"calldealerbuymoney"`
	CallDealerSellLotCount    int             `bson:"calldealerselllotcount"`
	CallDealerSellMoney       decimal.Decimal `bson:"calldealersellmoney"`
	CallDealerDiffLotCount    int             `bson:"calldealerdifflotcount"`
	CallDealerDiffMoney       decimal.Decimal `bson:"calldealerdiffmoney"`
	CallDealerOIBuyLotCount   int             `bson:"calldealeroibuylotcount"`
	CallDealerOIBuyMoney      decimal.Decimal `bson:"calldealeroibuymoney"`
	CallDealerOISellLotCount  int             `bson:"calldealeroiselllotcount"`
	CallDealerOISellMoney     decimal.Decimal `bson:"calldealeroisellmoney"`
	CallDealerOIDiffLotCount  int             `bson:"calldealeroidifflotcount"`
	CallDealerOIDiffMoney     decimal.Decimal `bson:"calldealeroidiffmoney"`
	CallTrustBuyLotCount      int             `bson:"calltrustbuylotcount"`
	CallTrustBuyMoney         decimal.Decimal `bson:"calltrustbuymoney"`
	CallTrustSellLotCount     int             `bson:"calltrustselllotcount"`
	CallTrustSellMoney        decimal.Decimal `bson:"calltrustsellmoney"`
	CallTrustDiffLotCount     int             `bson:"calltrustdifflotcount"`
	CallTrustDiffMoney        decimal.Decimal `bson:"calltrustdiffmoney"`
	CallTrustOIBuyLotCount    int             `bson:"calltrustoibuylotcount"`
	CallTrustOIBuyMoney       decimal.Decimal `bson:"calltrustoibuymoney"`
	CallTrustOISellLotCount   int             `bson:"calltrustoiselllotcount"`
	CallTrustOISellMoney      decimal.Decimal `bson:"calltrustoisellmoney"`
	CallTrustOIDiffLotCount   int             `bson:"calltrustoidifflotcount"`
	CallTrustOIDiffMoney      decimal.Decimal `bson:"calltrustoidiffmoney"`
	CallForeignBuyLotCount    int             `bson:"callforeignbuylotcount"`
	CallForeignBuyMoney       decimal.Decimal `bson:"callforeignbuymoney"`
	CallForeignSellLotCount   int             `bson:"callforeignselllotcount"`
	CallForeignSellMoney      decimal.Decimal `bson:"callforeignsellmoney"`
	CallForeignDiffLotCount   int             `bson:"callforeigndifflotcount"`
	CallForeignDiffMoney      decimal.Decimal `bson:"callforeigndiffmoney"`
	CallForeignOIBuyLotCount  int             `bson:"callforeignoibuylotcount"`
	CallForeignOIBuyMoney     decimal.Decimal `bson:"callforeignoibuymoney"`
	CallForeignOISellLotCount int             `bson:"callforeignoiselllotcount"`
	CallForeignOISellMoney    decimal.Decimal `bson:"callforeignoisellmoney"`
	CallForeignOIDiffLotCount int             `bson:"callforeignoidifflotcount"`
	CallForeignOIDiffMoney    decimal.Decimal `bson:"callforeignoidiffmoney"`
	PutDealerBuyLotCount      int             `bson:"putdealerbuylotcount"`
	PutDealerBuyMoney         decimal.Decimal `bson:"putdealerbuymoney"`
	PutDealerSellLotCount     int             `bson:"putdealerselllotcount"`
	PutDealerSellMoney        decimal.Decimal `bson:"putdealersellmoney"`
	PutDealerDiffLotCount     int             `bson:"putdealerdifflotcount"`
	PutDealerDiffMoney        decimal.Decimal `bson:"putdealerdiffmoney"`
	PutDealerOIBuyLotCount    int             `bson:"putdealeroibuylotcount"`
	PutDealerOIBuyMoney       decimal.Decimal `bson:"putdealeroibuymoney"`
	PutDealerOISellLotCount   int             `bson:"putdealeroiselllotcount"`
	PutDealerOISellMoney      decimal.Decimal `bson:"putdealeroisellmoney"`
	PutDealerOIDiffLotCount   int             `bson:"putdealeroidifflotcount"`
	PutDealerOIDiffMoney      decimal.Decimal `bson:"putdealeroidiffmoney"`
	PutTrustBuyLotCount       int             `bson:"puttrustbuylotcount"`
	PutTrustBuyMoney          decimal.Decimal `bson:"puttrustbuymoney"`
	PutTrustSellLotCount      int             `bson:"puttrustselllotcount"`
	PutTrustSellMoney         decimal.Decimal `bson:"puttrustsellmoney"`
	PutTrustDiffLotCount      int             `bson:"puttrustdifflotcount"`
	PutTrustDiffMoney         decimal.Decimal `bson:"puttrustdiffmoney"`
	PutTrustOIBuyLotCount     int             `bson:"puttrustoibuylotcount"`
	PutTrustOIBuyMoney        decimal.Decimal `bson:"puttrustoibuymoney"`
	PutTrustOISellLotCount    int             `bson:"puttrustoiselllotcount"`
	PutTrustOISellMoney       decimal.Decimal `bson:"puttrustoisellmoney"`
	PutTrustOIDiffLotCount    int             `bson:"puttrustoidifflotcount"`
	PutTrustOIDiffMoney       decimal.Decimal `bson:"puttrustoidiffmoney"`
	PutForeignBuyLotCount     int             `bson:"putforeignbuylotcount"`
	PutForeignBuyMoney        decimal.Decimal `bson:"putforeignbuymoney"`
	PutForeignSellLotCount    int             `bson:"putforeignselllotcount"`
	PutForeignSellMoney       decimal.Decimal `bson:"putforeignsellmoney"`
	PutForeignDiffLotCount    int             `bson:"putforeigndifflotcount"`
	PutForeignDiffMoney       decimal.Decimal `bson:"putforeigndiffmoney"`
	PutForeignOIBuyLotCount   int             `bson:"putforeignoibuylotcount"`
	PutForeignOIBuyMoney      decimal.Decimal `bson:"putforeignoibuymoney"`
	PutForeignOISellLotCount  int             `bson:"putforeignoiselllotcount"`
	PutForeignOISellMoney     decimal.Decimal `bson:"putforeignoisellmoney"`
	PutForeignOIDiffLotCount  int             `bson:"putforeignoidifflotcount"`
	PutForeignOIDiffMoney     decimal.Decimal `bson:"putforeignoidiffmoney"`
}

type BigThreeOption struct {
	DataDate                         time.Time       `bson:"datadate"`
	DealerLongLotCount               int             `bson:"dealerlonglotcount"`
	DealerLongMoney                  decimal.Decimal `bson:"dealerlongmoney"`
	DealerShortLotCount              int             `bson:"dealershortlotcount"`
	DealerShortMoney                 decimal.Decimal `bson:"dealershortmoney"`
	DealerLotNetCount                int             `bson:"dealerlotnetcount"`
	DealerNetMoney                   decimal.Decimal `bson:"dealernetmoney"`
	TrustLongLotCount                int             `bson:"trustlonglotcount"`
	TrustLongMoney                   decimal.Decimal `bson:"trustlongmoney"`
	TrustShortLotCount               int             `bson:"trustshortlotcount"`
	TrustShortMoney                  decimal.Decimal `bson:"trustshortmoney"`
	TrustLotNetCount                 int             `bson:"trustlotnetcount"`
	TrustNetMoney                    decimal.Decimal `bson:"trustnetmoney"`
	ForeignLongLotCount              int             `bson:"foreignlonglotcount"`
	ForeignLongMoney                 decimal.Decimal `bson:"foreignlongmoney"`
	ForeignShortLotCount             int             `bson:"foreignshortlotcount"`
	ForeignShortMoney                decimal.Decimal `bson:"foreignshortmoney"`
	ForeignLotNetCount               int             `bson:"foreignlotnetcount"`
	ForeignNetMoney                  decimal.Decimal `bson:"foreignnetmoney"`
	OpenInterestDealerLongLotCount   int             `bson:"openinterestdealerlonglotcount"`
	OpenInterestDealerLongMoney      decimal.Decimal `bson:"openinterestdealerlongmoney"`
	OpenInterestDealerShortLotCount  int             `bson:"openinterestdealershortlotcount"`
	OpenInterestDealerShortMoney     decimal.Decimal `bson:"openinterestdealershortmoney"`
	OpenInterestDealerLotNetCount    int             `bson:"openinterestdealerlotnetcount"`
	OpenInterestDealerNetMoney       decimal.Decimal `bson:"openinterestdealernetmoney"`
	OpenInterestTrustLongLotCount    int             `bson:"openinteresttrustlonglotcount"`
	OpenInterestTrustLongMoney       decimal.Decimal `bson:"openinteresttrustlongmoney"`
	OpenInterestTrustShortLotCount   int             `bson:"openinteresttrustshortlotcount"`
	OpenInterestTrustShortMoney      decimal.Decimal `bson:"openinteresttrustshortmoney"`
	OpenInterestTrustLotNetCount     int             `bson:"openinteresttrustlotnetcount"`
	OpenInterestTrustNetMoney        decimal.Decimal `bson:"openinteresttrustnetmoney"`
	OpenInterestForeignLongLotCount  int             `bson:"openinterestforeignlonglotcount"`
	OpenInterestForeignLongMoney     decimal.Decimal `bson:"openinterestforeignlongmoney"`
	OpenInterestForeignShortLotCount int             `bson:"openinterestforeignshortlotcount"`
	OpenInterestForeignShortMoney    decimal.Decimal `bson:"openinterestforeignshortmoney"`
	OpenInterestForeignLotNetCount   int             `bson:"openinterestforeignlotnetcount"`
	OpenInterestForeignNetMoney      decimal.Decimal `bson:"openinterestforeignnetmoney"`
}

type BigThreeFuture struct {
	DataDate                         time.Time       `bson:"datadate"`
	DealerLongLotCount               int             `bson:"dealerlonglotcount"`
	DealerLongMoney                  decimal.Decimal `bson:"dealerlongmoney"`
	DealerShortLotCount              int             `bson:"dealershortlotcount"`
	DealerShortMoney                 decimal.Decimal `bson:"dealershortmoney"`
	DealerLotNetCount                int             `bson:"dealerlotnetcount"`
	DealerNetMoney                   decimal.Decimal `bson:"dealernetmoney"`
	TrustLongLotCount                int             `bson:"trustlonglotcount"`
	TrustLongMoney                   decimal.Decimal `bson:"trustlongmoney"`
	TrustShortLotCount               int             `bson:"trustshortlotcount"`
	TrustShortMoney                  decimal.Decimal `bson:"trustshortmoney"`
	TrustLotNetCount                 int             `bson:"trustlotnetcount"`
	TrustNetMoney                    decimal.Decimal `bson:"trustnetmoney"`
	ForeignLongLotCount              int             `bson:"foreignlonglotcount"`
	ForeignLongMoney                 decimal.Decimal `bson:"foreignlongmoney"`
	ForeignShortLotCount             int             `bson:"foreignshortlotcount"`
	ForeignShortMoney                decimal.Decimal `bson:"foreignshortmoney"`
	ForeignLotNetCount               int             `bson:"foreignlotnetcount"`
	ForeignNetMoney                  decimal.Decimal `bson:"foreignnetmoney"`
	OpenInterestDealerLongLotCount   int             `bson:"openinterestdealerlonglotcount"`
	OpenInterestDealerLongMoney      decimal.Decimal `bson:"openinterestdealerlongmoney"`
	OpenInterestDealerShortLotCount  int             `bson:"openinterestdealershortlotcount"`
	OpenInterestDealerShortMoney     decimal.Decimal `bson:"openinterestdealershortmoney"`
	OpenInterestDealerLotNetCount    int             `bson:"openinterestdealerlotnetcount"`
	OpenInterestDealerNetMoney       decimal.Decimal `bson:"openinterestdealernetmoney"`
	OpenInterestTrustLongLotCount    int             `bson:"openinteresttrustlonglotcount"`
	OpenInterestTrustLongMoney       decimal.Decimal `bson:"openinteresttrustlongmoney"`
	OpenInterestTrustShortLotCount   int             `bson:"openinteresttrustshortlotcount"`
	OpenInterestTrustShortMoney      decimal.Decimal `bson:"openinteresttrustshortmoney"`
	OpenInterestTrustLotNetCount     int             `bson:"openinteresttrustlotnetcount"`
	OpenInterestTrustNetMoney        decimal.Decimal `bson:"openinteresttrustnetmoney"`
	OpenInterestForeignLongLotCount  int             `bson:"openinterestforeignlonglotcount"`
	OpenInterestForeignLongMoney     decimal.Decimal `bson:"openinterestforeignlongmoney"`
	OpenInterestForeignShortLotCount int             `bson:"openinterestforeignshortlotcount"`
	OpenInterestForeignShortMoney    decimal.Decimal `bson:"openinterestforeignshortmoney"`
	OpenInterestForeignLotNetCount   int             `bson:"openinterestforeignlotnetcount"`
	OpenInterestForeignNetMoney      decimal.Decimal `bson:"openinterestforeignnetmoney"`
}

type TopFutureOI struct {
	DataDate                  time.Time       `bson:"datadate"`
	WeekTop5BuyLot            int             `bson:"weektop5buylot"`
	WeekTop5BuyPercent        decimal.Decimal `bson:"weektop5buypercent"`
	WeekTop5SpecBuyLot        int             `bson:"weektop5specbuylot"`
	WeekTop5SpecBuyPercent    decimal.Decimal `bson:"weektop5specbuypercent"`
	WeekTop5SellLot           int             `bson:"weektop5selllot"`
	WeekTop5SellPercent       decimal.Decimal `bson:"weektop5sellpercent"`
	WeekTop5SpecSellLot       int             `bson:"weektop5specselllot"`
	WeekTop5SpecSellPercent   decimal.Decimal `bson:"weektop5specsellpercent"`
	MonthTop5BuyLot           int             `bson:"monthtop5buylot"`
	MonthTop5BuyPercent       decimal.Decimal `bson:"monthtop5buypercent"`
	MonthTop5SpecBuyLot       int             `bson:"monthtop5specbuylot"`
	MonthTop5SpecBuyPercent   decimal.Decimal `bson:"monthtop5specbuypercent"`
	MonthTop5SellLot          int             `bson:"monthtop5selllot"`
	MonthTop5SellPercent      decimal.Decimal `bson:"monthtop5sellpercent"`
	MonthTop5SpecSellLot      int             `bson:"monthtop5specselllot"`
	MonthTop5SpecSellPercent  decimal.Decimal `bson:"monthtop5specsellpercent"`
	AllTop5BuyLot             int             `bson:"alltop5buylot"`
	AllTop5BuyPercent         decimal.Decimal `bson:"alltop5buypercent"`
	AllTop5SpecBuyLot         int             `bson:"alltop5specbuylot"`
	AllTop5SpecBuyPercent     decimal.Decimal `bson:"alltop5specbuypercent"`
	AllTop5SellLot            int             `bson:"alltop5selllot"`
	AllTop5SellPercent        decimal.Decimal `bson:"alltop5sellpercent"`
	AllTop5SpecSellLot        int             `bson:"alltop5specselllot"`
	AllTop5SpecSellPercent    decimal.Decimal `bson:"alltop5specsellpercent"`
	WeekTop10BuyLot           int             `bson:"weektop10buylot"`
	WeekTop10BuyPercent       decimal.Decimal `bson:"weektop10buypercent"`
	WeekTop10SpecBuyLot       int             `bson:"weektop10specbuylot"`
	WeekTop10SpecBuyPercent   decimal.Decimal `bson:"weektop10specbuypercent"`
	WeekTop10SellLot          int             `bson:"weektop10selllot"`
	WeekTop10SellPercent      decimal.Decimal `bson:"weektop10sellpercent"`
	WeekTop10SpecSellLot      int             `bson:"weektop10specselllot"`
	WeekTop10SpecSellPercent  decimal.Decimal `bson:"weektop10specsellpercent"`
	MonthTop10BuyLot          int             `bson:"monthtop10buylot"`
	MonthTop10BuyPercent      decimal.Decimal `bson:"monthtop10buypercent"`
	MonthTop10SpecBuyLot      int             `bson:"monthtop10specbuylot"`
	MonthTop10SpecBuyPercent  decimal.Decimal `bson:"monthtop10specbuypercent"`
	MonthTop10SellLot         int             `bson:"monthtop10selllot"`
	MonthTop10SellPercent     decimal.Decimal `bson:"monthtop10sellpercent"`
	MonthTop10SpecSellLot     int             `bson:"monthtop10specselllot"`
	MonthTop10SpecSellPercent decimal.Decimal `bson:"monthtop10specsellpercent"`
	AllTop10BuyLot            int             `bson:"alltop10buylot"`
	AllTop10BuyPercent        decimal.Decimal `bson:"alltop10buypercent"`
	AllTop10SpecBuyLot        int             `bson:"alltop10specbuylot"`
	AllTop10SpecBuyPercent    decimal.Decimal `bson:"alltop10specbuypercent"`
	AllTop10SellLot           int             `bson:"alltop10selllot"`
	AllTop10SellPercent       decimal.Decimal `bson:"alltop10sellpercent"`
	AllTop10SpecSellLot       int             `bson:"alltop10specselllot"`
	AllTop10SpecSellPercent   decimal.Decimal `bson:"alltop10specsellpercent"`
	WeekTotalOI               int             `bson:"weektotaloi"`
	MonthTotalOI              int             `bson:"monthtotaloi"`
	AllTotalOI                int             `bson:"alltotaloi"`
	MonthContract             string          `bson:"monthcontract"`
}

type TopOptionOI struct {
	DataDate                      time.Time       `bson:"datadate"`
	CallWeekTop5BuyLot            int             `bson:"callweektop5buylot"`
	CallWeekTop5BuyPercent        decimal.Decimal `bson:"callweektop5buypercent"`
	CallWeekTop5SpecBuyLot        int             `bson:"callweektop5specbuylot"`
	CallWeekTop5SpecBuyPercent    decimal.Decimal `bson:"callweektop5specbuypercent"`
	CallWeekTop5SellLot           int             `bson:"callweektop5selllot"`
	CallWeekTop5SellPercent       decimal.Decimal `bson:"callweektop5sellpercent"`
	CallWeekTop5SpecSellLot       int             `bson:"callweektop5specselllot"`
	CallWeekTop5SpecSellPercent   decimal.Decimal `bson:"callweektop5specsellpercent"`
	CallMonthTop5BuyLot           int             `bson:"callmonthtop5buylot"`
	CallMonthTop5BuyPercent       decimal.Decimal `bson:"callmonthtop5buypercent"`
	CallMonthTop5SpecBuyLot       int             `bson:"callmonthtop5specbuylot"`
	CallMonthTop5SpecBuyPercent   decimal.Decimal `bson:"callmonthtop5specbuypercent"`
	CallMonthTop5SellLot          int             `bson:"callmonthtop5selllot"`
	CallMonthTop5SellPercent      decimal.Decimal `bson:"callmonthtop5sellpercent"`
	CallMonthTop5SpecSellLot      int             `bson:"callmonthtop5specselllot"`
	CallMonthTop5SpecSellPercent  decimal.Decimal `bson:"callmonthtop5specsellpercent"`
	CallAllTop5BuyLot             int             `bson:"callalltop5buylot"`
	CallAllTop5BuyPercent         decimal.Decimal `bson:"callalltop5buypercent"`
	CallAllTop5SpecBuyLot         int             `bson:"callalltop5specbuylot"`
	CallAllTop5SpecBuyPercent     decimal.Decimal `bson:"callalltop5specbuypercent"`
	CallAllTop5SellLot            int             `bson:"callalltop5selllot"`
	CallAllTop5SellPercent        decimal.Decimal `bson:"callalltop5sellpercent"`
	CallAllTop5SpecSellLot        int             `bson:"callalltop5specselllot"`
	CallAllTop5SpecSellPercent    decimal.Decimal `bson:"callalltop5specsellpercent"`
	CallWeekTop10BuyLot           int             `bson:"callweektop10buylot"`
	CallWeekTop10BuyPercent       decimal.Decimal `bson:"callweektop10buypercent"`
	CallWeekTop10SpecBuyLot       int             `bson:"callweektop10specbuylot"`
	CallWeekTop10SpecBuyPercent   decimal.Decimal `bson:"callweektop10specbuypercent"`
	CallWeekTop10SellLot          int             `bson:"callweektop10selllot"`
	CallWeekTop10SellPercent      decimal.Decimal `bson:"callweektop10sellpercent"`
	CallWeekTop10SpecSellLot      int             `bson:"callweektop10specselllot"`
	CallWeekTop10SpecSellPercent  decimal.Decimal `bson:"callweektop10specsellpercent"`
	CallMonthTop10BuyLot          int             `bson:"callmonthtop10buylot"`
	CallMonthTop10BuyPercent      decimal.Decimal `bson:"callmonthtop10buypercent"`
	CallMonthTop10SpecBuyLot      int             `bson:"callmonthtop10specbuylot"`
	CallMonthTop10SpecBuyPercent  decimal.Decimal `bson:"callmonthtop10specbuypercent"`
	CallMonthTop10SellLot         int             `bson:"callmonthtop10selllot"`
	CallMonthTop10SellPercent     decimal.Decimal `bson:"callmonthtop10sellpercent"`
	CallMonthTop10SpecSellLot     int             `bson:"callmonthtop10specselllot"`
	CallMonthTop10SpecSellPercent decimal.Decimal `bson:"callmonthtop10specsellpercent"`
	CallAllTop10BuyLot            int             `bson:"callalltop10buylot"`
	CallAllTop10BuyPercent        decimal.Decimal `bson:"callalltop10buypercent"`
	CallAllTop10SpecBuyLot        int             `bson:"callalltop10specbuylot"`
	CallAllTop10SpecBuyPercent    decimal.Decimal `bson:"callalltop10specbuypercent"`
	CallAllTop10SellLot           int             `bson:"callalltop10selllot"`
	CallAllTop10SellPercent       decimal.Decimal `bson:"callalltop10sellpercent"`
	CallAllTop10SpecSellLot       int             `bson:"callalltop10specselllot"`
	CallAllTop10SpecSellPercent   decimal.Decimal `bson:"callalltop10specsellpercent"`
	CallWeekTotalOI               int             `bson:"callweektotaloi"`
	CallMonthTotalOI              int             `bson:"callmonthtotaloi"`
	CallAllTotalOI                int             `bson:"callalltotaloi"`
	CallMonthContract             string          `bson:"callmonthcontract"`
	PutWeekTop5BuyLot             int             `bson:"putweektop5buylot"`
	PutWeekTop5BuyPercent         decimal.Decimal `bson:"putweektop5buypercent"`
	PutWeekTop5SpecBuyLot         int             `bson:"putweektop5specbuylot"`
	PutWeekTop5SpecBuyPercent     decimal.Decimal `bson:"putweektop5specbuypercent"`
	PutWeekTop5SellLot            int             `bson:"putweektop5selllot"`
	PutWeekTop5SellPercent        decimal.Decimal `bson:"putweektop5sellpercent"`
	PutWeekTop5SpecSellLot        int             `bson:"putweektop5specselllot"`
	PutWeekTop5SpecSellPercent    decimal.Decimal `bson:"putweektop5specsellpercent"`
	PutMonthTop5BuyLot            int             `bson:"putmonthtop5buylot"`
	PutMonthTop5BuyPercent        decimal.Decimal `bson:"putmonthtop5buypercent"`
	PutMonthTop5SpecBuyLot        int             `bson:"putmonthtop5specbuylot"`
	PutMonthTop5SpecBuyPercent    decimal.Decimal `bson:"putmonthtop5specbuypercent"`
	PutMonthTop5SellLot           int             `bson:"putmonthtop5selllot"`
	PutMonthTop5SellPercent       decimal.Decimal `bson:"putmonthtop5sellpercent"`
	PutMonthTop5SpecSellLot       int             `bson:"putmonthtop5specselllot"`
	PutMonthTop5SpecSellPercent   decimal.Decimal `bson:"putmonthtop5specsellpercent"`
	PutAllTop5BuyLot              int             `bson:"putalltop5buylot"`
	PutAllTop5BuyPercent          decimal.Decimal `bson:"putalltop5buypercent"`
	PutAllTop5SpecBuyLot          int             `bson:"putalltop5specbuylot"`
	PutAllTop5SpecBuyPercent      decimal.Decimal `bson:"putalltop5specbuypercent"`
	PutAllTop5SellLot             int             `bson:"putalltop5selllot"`
	PutAllTop5SellPercent         decimal.Decimal `bson:"putalltop5sellpercent"`
	PutAllTop5SpecSellLot         int             `bson:"putalltop5specselllot"`
	PutAllTop5SpecSellPercent     decimal.Decimal `bson:"putalltop5specsellpercent"`
	PutWeekTop10BuyLot            int             `bson:"putweektop10buylot"`
	PutWeekTop10BuyPercent        decimal.Decimal `bson:"putweektop10buypercent"`
	PutWeekTop10SpecBuyLot        int             `bson:"putweektop10specbuylot"`
	PutWeekTop10SpecBuyPercent    decimal.Decimal `bson:"putweektop10specbuypercent"`
	PutWeekTop10SellLot           int             `bson:"putweektop10selllot"`
	PutWeekTop10SellPercent       decimal.Decimal `bson:"putweektop10sellpercent"`
	PutWeekTop10SpecSellLot       int             `bson:"putweektop10specselllot"`
	PutWeekTop10SpecSellPercent   decimal.Decimal `bson:"putweektop10specsellpercent"`
	PutMonthTop10BuyLot           int             `bson:"putmonthtop10buylot"`
	PutMonthTop10BuyPercent       decimal.Decimal `bson:"putmonthtop10buypercent"`
	PutMonthTop10SpecBuyLot       int             `bson:"putmonthtop10specbuylot"`
	PutMonthTop10SpecBuyPercent   decimal.Decimal `bson:"putmonthtop10specbuypercent"`
	PutMonthTop10SellLot          int             `bson:"putmonthtop10selllot"`
	PutMonthTop10SellPercent      decimal.Decimal `bson:"putmonthtop10sellpercent"`
	PutMonthTop10SpecSellLot      int             `bson:"putmonthtop10specselllot"`
	PutMonthTop10SpecSellPercent  decimal.Decimal `bson:"putmonthtop10specsellpercent"`
	PutAllTop10BuyLot             int             `bson:"putalltop10buylot"`
	PutAllTop10BuyPercent         decimal.Decimal `bson:"putalltop10buypercent"`
	PutAllTop10SpecBuyLot         int             `bson:"putalltop10specbuylot"`
	PutAllTop10SpecBuyPercent     decimal.Decimal `bson:"putalltop10specbuypercent"`
	PutAllTop10SellLot            int             `bson:"putalltop10selllot"`
	PutAllTop10SellPercent        decimal.Decimal `bson:"putalltop10sellpercent"`
	PutAllTop10SpecSellLot        int             `bson:"putalltop10specselllot"`
	PutAllTop10SpecSellPercent    decimal.Decimal `bson:"putalltop10specsellpercent"`
	PutWeekTotalOI                int             `bson:"putweektotaloi"`
	PutMonthTotalOI               int             `bson:"putmonthtotaloi"`
	PutAllTotalOI                 int             `bson:"putalltotaloi"`
	PutMonthContract              string          `bson:"putmonthcontract"`
}

type MarketPutCallRatio struct {
	DataDate          time.Time       `bson:"datadate"`
	CallTradeVolume   int             `bson:"calltradevolume"`
	PutTradeVolume    int             `bson:"puttradevolume"`
	PutCallTradeRatio decimal.Decimal `bson:"putcalltraderatio"`
	OICallVolume      int             `bson:"oicallvolume"`
	OIPutVolume       int             `bson:"oiputvolume"`
	PutCallOIRatio    decimal.Decimal `bson:"putcalloiratio"`
}
