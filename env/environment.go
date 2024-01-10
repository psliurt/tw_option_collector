package env

import (
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

var envInstance *Environment
var createEnvironmentOnce sync.Once

type Environment struct {
	ServicePort    int
	MgoDb          *Mgo
	MgoCollections *MongoDbCollections
}

func Initialize() {
	createEnvironmentOnce.Do(func() {
		envInstance = createEnvironment()
	})
}

func Instance() *Environment {
	Initialize()
	return envInstance
}

func createEnvironment() *Environment {
	setUpZap()
	return &Environment{
		ServicePort:    loadServicePort(),
		MgoDb:          loadMongoSetting(),
		MgoCollections: loadMongoDbCollections(),
	}
}

type Mgo struct {
	Host        []string
	AdminDbName string
	MinPoolSize int
	PoolLimit   int
	UserName    string
	Password    string
	DbName      string
	Client      *mongo.Client
}

type MongoDbCollections struct {
	BigThreeSummary      string //三大法人 總表
	BigThreeFutureOption string //三大法人 區分期貨與選擇權二類
	BigThreeFuture       string //三大法人 各期貨契約
	BigThreeOption       string //三大法人 各選擇權契約
	BigThreeCallPut      string //三大法人 買賣權分計
	TopTradeFutureOI     string //大額交易人期貨未沖銷
	TopTradeOptionOI     string //大額交易人選擇權未沖銷
	PutCallRatio         string
}

type QuoteServerConfig struct {
	Name     string
	Protocol string
	Address  string
	Port     int
}
