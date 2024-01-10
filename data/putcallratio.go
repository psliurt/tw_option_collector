package data

import (
	"time"
	"tw_option_collector/objects"

	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func InsertPutCallRatio(data *objects.MarketPutCallRatio) error {
	helper := NewMongoDbHelper(GetDBCol().PutCallRatio)
	defer helper.Close()
	err := helper.InsertOne(data, GetDBCol().PutCallRatio)
	if err != nil {
		zap.L().Error("add put call ratio error", zap.Error(err))
		return err
	}
	return nil
}

func FindPutCallRatioByDate(date time.Time) (*objects.MarketPutCallRatio, error) {
	helper := NewMongoDbHelper(GetDBCol().PutCallRatio)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	var obj objects.MarketPutCallRatio
	err := helper.FetchOne(filter, GetDBCol().PutCallRatio).Decode(&obj)
	if err != nil {
		zap.L().Error("find market put call ratio error", zap.Time("date", date), zap.Error(err))
		return nil, err
	}
	return &obj, nil
}

func FindPutCallRatioCountByDate(date time.Time) (int64, error) {
	helper := NewMongoDbHelper(GetDBCol().PutCallRatio)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	count, err := helper.FetchDocCount(filter, GetDBCol().PutCallRatio)
	if err != nil {
		zap.L().Error("find put call ratio error", zap.Time("date", date), zap.Error(err))
		return 0, err
	}
	return count, nil
}
