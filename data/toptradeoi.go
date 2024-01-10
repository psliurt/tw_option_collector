package data

import (
	"time"
	"tw_option_collector/objects"

	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func InsertTopTradeOIFuture(data *objects.TopFutureOI) error {
	helper := NewMongoDbHelper(GetDBCol().TopTradeFutureOI)
	defer helper.Close()
	err := helper.InsertOne(data, GetDBCol().TopTradeFutureOI)
	if err != nil {
		zap.L().Error("add top trade future oi error", zap.Error(err))
		return err
	}
	return nil
}

func FindTopTradeOIFutureByDate(date time.Time) (*objects.TopFutureOI, error) {
	helper := NewMongoDbHelper(GetDBCol().TopTradeFutureOI)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	var obj objects.TopFutureOI
	err := helper.FetchOne(filter, GetDBCol().TopTradeFutureOI).Decode(&obj)
	if err != nil {
		zap.L().Error("find top trade future oi error", zap.Time("date", date), zap.Error(err))
		return nil, err
	}
	return &obj, nil
}

func FindTopTradeOIFutureCountByDate(date time.Time) (int64, error) {
	helper := NewMongoDbHelper(GetDBCol().TopTradeFutureOI)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	count, err := helper.FetchDocCount(filter, GetDBCol().TopTradeFutureOI)
	if err != nil {
		zap.L().Error("find top trade future oi error", zap.Time("date", date), zap.Error(err))
		return 0, err
	}
	return count, nil
}

func InsertTopTradeOIOption(data *objects.TopOptionOI) error {
	helper := NewMongoDbHelper(GetDBCol().TopTradeOptionOI)
	defer helper.Close()
	err := helper.InsertOne(data, GetDBCol().TopTradeOptionOI)
	if err != nil {
		zap.L().Error("add top trade option oi error", zap.Error(err))
		return err
	}
	return nil
}

func FindTopTradeOIOptionByDate(date time.Time) (*objects.TopOptionOI, error) {
	helper := NewMongoDbHelper(GetDBCol().TopTradeOptionOI)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	var obj objects.TopOptionOI
	err := helper.FetchOne(filter, GetDBCol().TopTradeOptionOI).Decode(&obj)
	if err != nil {
		zap.L().Error("find top trade option oi error", zap.Time("date", date), zap.Error(err))
		return nil, err
	}
	return &obj, nil
}

func FindTopTradeOIOptionCountByDate(date time.Time) (int64, error) {
	helper := NewMongoDbHelper(GetDBCol().TopTradeOptionOI)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	count, err := helper.FetchDocCount(filter, GetDBCol().TopTradeOptionOI)
	if err != nil {
		zap.L().Error("find top trade option oi error", zap.Time("date", date), zap.Error(err))
		return 0, err
	}
	return count, nil
}
