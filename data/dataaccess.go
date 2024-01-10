package data

import (
	"context"
	"tw_option_collector/env"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoDbHelper struct {
	ctx     context.Context
	session mongo.Session
	colsMap map[string]*mongo.Collection
	curs    []*mongo.Cursor
}

func GetDBCol() *env.MongoDbCollections {
	return env.Instance().MgoCollections
}

func NewMongoDbHelper(cols ...string) *MongoDbHelper {
	environment := env.Instance()
	ses, err := environment.MgoDb.Client.StartSession(options.Session())
	if err != nil {
		zap.L().Info("start session error", zap.Any("collections", cols), zap.Error(err))
		panic(err)
	}

	m := make(map[string]*mongo.Collection)
	for _, v := range cols {
		m[v] = ses.Client().Database(environment.MgoDb.DbName).Collection(v)
	}

	return &MongoDbHelper{
		ctx:     context.Background(),
		session: ses,
		colsMap: m,
		curs:    make([]*mongo.Cursor, 0),
	}
}

func (mh *MongoDbHelper) getCollection(colName string) *mongo.Collection {
	col, ok := mh.colsMap[colName]
	if !ok {
		col = mh.session.Client().Database(env.Instance().MgoDb.DbName).Collection(colName)
		mh.colsMap[colName] = col
	}
	return col
}

func (mh *MongoDbHelper) InsertOne(newObject interface{}, colName string) error {

	col := mh.getCollection(colName)
	rst, err := col.InsertOne(mh.ctx, newObject)
	if err != nil {
		zap.L().Error("add data to collection error", zap.Any("obj", newObject), zap.String("collection", colName), zap.Error(err))
		return err
	}

	zap.L().Info("add data to collection success",
		zap.Any("obj", newObject), zap.String("collection", colName), zap.Any("insertid", rst.InsertedID))

	return nil
}

func (mh *MongoDbHelper) UpdateOne(filter interface{}, obj interface{}, colName string) error {
	col := mh.getCollection(colName)
	_, err := col.UpdateOne(mh.ctx, filter, bson.D{{"$set", obj}})

	if err != nil {
		zap.L().Error("update data to collection error", zap.Any("filter", filter), zap.Any("obj", obj), zap.String("collection", colName), zap.Error(err))
		return err
	}

	return nil
}

func (mh *MongoDbHelper) UpdateInsertOne(filter interface{}, obj interface{}, colName string) error {
	col := mh.getCollection(colName)

	opts := options.Update().SetUpsert(true)
	_, err := col.UpdateOne(mh.ctx, filter, bson.D{{"$set", obj}}, opts)

	if err != nil {
		zap.L().Error("upsert data to collection error", zap.Any("filter", filter), zap.Any("obj", obj), zap.String("collection", colName), zap.Error(err))
		return err
	}

	return nil
}

func (mh *MongoDbHelper) UpdateManySpec(filter interface{}, update interface{}, colName string) error {
	col := mh.getCollection(colName)
	_, err := col.UpdateMany(mh.ctx, filter, update)

	if err != nil {
		zap.L().Error("upsert many data to collection error",
			zap.Any("filter", filter), zap.Any("update", update), zap.String("collection", colName), zap.Error(err))
		return err
	}

	return nil
}

func (mh *MongoDbHelper) DeleteOne(filter interface{}, colName string) error {
	col := mh.getCollection(colName)
	_, err := col.DeleteOne(mh.ctx, filter)

	if err != nil {
		zap.L().Error("delete data from collection error", zap.Any("filter", filter), zap.String("collection", colName), zap.Error(err))
		return err
	}

	return nil
}

func (mh *MongoDbHelper) CheckExistData(filter interface{}, colName string) (bool, error) {
	col := mh.getCollection(colName)
	opts := options.Count().SetLimit(1)
	c, err := col.CountDocuments(mh.ctx, filter, opts)
	if err != nil {
		zap.L().Error("get collection documents exist error ", zap.Any("filter", filter), zap.String("collection", colName), zap.Error(err))
		return false, err
	}
	if c > 0 {
		return true, nil
	}
	return false, nil
}

func (mh *MongoDbHelper) Fetch(filter interface{}, sort interface{}, colName string) (*mongo.Cursor, error) {
	col := mh.getCollection(colName)
	opts := options.Find().SetSort(sort)
	cur, err := col.Find(mh.ctx, filter, opts)
	if err != nil {
		zap.L().Error("fetch all cursor error", zap.Any("filter", filter), zap.Any("sort", sort), zap.String("collection", colName), zap.Error(err))
		return nil, err
	}

	mh.curs = append(mh.curs, cur)

	return cur, nil
}

func (mh *MongoDbHelper) FetchPagedData(filter interface{}, sort interface{}, skip int64, count int64, colName string) (*mongo.Cursor, error) {
	col := mh.getCollection(colName)
	opts := options.Find().SetSort(sort).SetSkip(skip).SetLimit(count)
	cur, err := col.Find(mh.ctx, filter, opts)
	if err != nil {
		zap.L().Error("fetch paged data cursor error",
			zap.Any("filter", filter), zap.Any("sort", sort), zap.String("collection", colName),
			zap.Int64("skip", skip), zap.Int64("count", count), zap.Error(err))
		return nil, err
	}
	mh.curs = append(mh.curs, cur)
	return cur, nil
}

func (mh *MongoDbHelper) FetchOne(filter interface{}, colName string) *mongo.SingleResult {
	col := mh.getCollection(colName)

	return col.FindOne(mh.ctx, filter)
}

func (mh *MongoDbHelper) FetchDocCount(filter interface{}, colName string) (int64, error) {
	col := mh.getCollection(colName)

	return col.CountDocuments(mh.ctx, filter)
}

func (mh *MongoDbHelper) FindOneAndUpdate(filter interface{}, update interface{}, colName string) *mongo.SingleResult {
	col := mh.getCollection(colName)
	return col.FindOneAndUpdate(mh.ctx, filter, update)
}

func (mh *MongoDbHelper) Close() {
	mh.session.EndSession(mh.ctx)
	mh.colsMap = make(map[string]*mongo.Collection)
	var err error
	for i, _ := range mh.curs {
		err = mh.curs[i].Close(mh.ctx)
		if err != nil {
			zap.L().Error("close mongo cursor error", zap.Error(err))
		}
	}
}
