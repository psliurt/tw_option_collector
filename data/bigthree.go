package data

import (
	"time"
	"tw_option_collector/objects"

	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func InsertBigThreeSummary(data *objects.BigThreeSummary) error {
	helper := NewMongoDbHelper(GetDBCol().BigThreeSummary)
	defer helper.Close()
	err := helper.InsertOne(data, GetDBCol().BigThreeSummary)
	if err != nil {
		zap.L().Error("add big three summary error", zap.Error(err))
		return err
	}
	return nil
}

func FindBigThreeSummaryByDate(date time.Time) (*objects.BigThreeSummary, error) {
	helper := NewMongoDbHelper(GetDBCol().BigThreeSummary)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	var obj objects.BigThreeSummary
	err := helper.FetchOne(filter, GetDBCol().BigThreeSummary).Decode(&obj)
	if err != nil {
		zap.L().Error("find big three summary error", zap.Time("date", date), zap.Error(err))
		return nil, err
	}
	return &obj, nil
}

func FindBigThreeSummaryCountByDate(date time.Time) (int64, error) {
	helper := NewMongoDbHelper(GetDBCol().BigThreeSummary)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	count, err := helper.FetchDocCount(filter, GetDBCol().BigThreeSummary)
	if err != nil {
		zap.L().Error("find big three summary error", zap.Time("date", date), zap.Error(err))
		return 0, err
	}
	return count, nil
}

func InsertBigThreeFutureOption(data *objects.BigThreeFutureOption) error {
	helper := NewMongoDbHelper(GetDBCol().BigThreeFutureOption)
	defer helper.Close()
	err := helper.InsertOne(data, GetDBCol().BigThreeFutureOption)
	if err != nil {
		zap.L().Error("add big three future option error", zap.Error(err))
		return err
	}
	return nil
}

func FindBigThreeFutureOptionByDate(date time.Time) (*objects.BigThreeFutureOption, error) {
	helper := NewMongoDbHelper(GetDBCol().BigThreeFutureOption)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	var obj objects.BigThreeFutureOption
	err := helper.FetchOne(filter, GetDBCol().BigThreeFutureOption).Decode(&obj)
	if err != nil {
		zap.L().Error("find big three future option error", zap.Time("date", date), zap.Error(err))
		return nil, err
	}
	return &obj, nil
}

func FindBigThreeFutureOptionCountByDate(date time.Time) (int64, error) {
	helper := NewMongoDbHelper(GetDBCol().BigThreeFutureOption)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	count, err := helper.FetchDocCount(filter, GetDBCol().BigThreeFutureOption)
	if err != nil {
		zap.L().Error("find big three future option error", zap.Time("date", date), zap.Error(err))
		return 0, err
	}
	return count, nil
}

func InsertBigThreeCallPut(data *objects.BigThreeCallPut) error {
	helper := NewMongoDbHelper(GetDBCol().BigThreeCallPut)
	defer helper.Close()
	err := helper.InsertOne(data, GetDBCol().BigThreeCallPut)
	if err != nil {
		zap.L().Error("add big three call put error", zap.Error(err))
		return err
	}
	return nil
}

func FindBigThreeCallPutByDate(date time.Time) (*objects.BigThreeCallPut, error) {
	helper := NewMongoDbHelper(GetDBCol().BigThreeCallPut)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	var obj objects.BigThreeCallPut
	err := helper.FetchOne(filter, GetDBCol().BigThreeCallPut).Decode(&obj)
	if err != nil {
		zap.L().Error("find big three call put error", zap.Time("date", date), zap.Error(err))
		return nil, err
	}
	return &obj, nil
}

func FindBigThreeCallPutCountByDate(date time.Time) (int64, error) {
	helper := NewMongoDbHelper(GetDBCol().BigThreeCallPut)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	count, err := helper.FetchDocCount(filter, GetDBCol().BigThreeCallPut)
	if err != nil {
		zap.L().Error("find big three call put error", zap.Time("date", date), zap.Error(err))
		return 0, err
	}
	return count, nil
}

func InsertBigThreeOption(data *objects.BigThreeOption) error {
	helper := NewMongoDbHelper(GetDBCol().BigThreeOption)
	defer helper.Close()
	err := helper.InsertOne(data, GetDBCol().BigThreeOption)
	if err != nil {
		zap.L().Error("add big three option error", zap.Error(err))
		return err
	}
	return nil
}

func FindBigThreeOptionByDate(date time.Time) (*objects.BigThreeOption, error) {
	helper := NewMongoDbHelper(GetDBCol().BigThreeOption)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	var obj objects.BigThreeOption
	err := helper.FetchOne(filter, GetDBCol().BigThreeOption).Decode(&obj)
	if err != nil {
		zap.L().Error("find big three option error", zap.Time("date", date), zap.Error(err))
		return nil, err
	}
	return &obj, nil
}

func FindBigThreeOptionCountByDate(date time.Time) (int64, error) {
	helper := NewMongoDbHelper(GetDBCol().BigThreeOption)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	count, err := helper.FetchDocCount(filter, GetDBCol().BigThreeOption)
	if err != nil {
		zap.L().Error("find big three option error", zap.Time("date", date), zap.Error(err))
		return 0, err
	}
	return count, nil
}

func InsertBigThreeFuture(data *objects.BigThreeFuture) error {
	helper := NewMongoDbHelper(GetDBCol().BigThreeFuture)
	defer helper.Close()
	err := helper.InsertOne(data, GetDBCol().BigThreeFuture)
	if err != nil {
		zap.L().Error("add big three future error", zap.Error(err))
		return err
	}
	return nil
}

func FindBigThreeFutureByDate(date time.Time) (*objects.BigThreeFuture, error) {
	helper := NewMongoDbHelper(GetDBCol().BigThreeFuture)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	var obj objects.BigThreeFuture
	err := helper.FetchOne(filter, GetDBCol().BigThreeFuture).Decode(&obj)
	if err != nil {
		zap.L().Error("find big three future error", zap.Time("date", date), zap.Error(err))
		return nil, err
	}
	return &obj, nil
}

func FindBigThreeFutureCountByDate(date time.Time) (int64, error) {
	helper := NewMongoDbHelper(GetDBCol().BigThreeFuture)
	defer helper.Close()
	filter := bson.D{{"datadate", date}}
	count, err := helper.FetchDocCount(filter, GetDBCol().BigThreeFuture)
	if err != nil {
		zap.L().Error("find big three future error", zap.Time("date", date), zap.Error(err))
		return 0, err
	}
	return count, nil
}
