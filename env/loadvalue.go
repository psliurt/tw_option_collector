package env

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"time"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"

	b64 "encoding/base64"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func setUpZap() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/pinaryplayer.log",
		MaxSize:    30,  // megabytes
		MaxBackups: 300, // backup files
		MaxAge:     31,  // days
	})
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.CallerKey = "caller"
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		w,
		zap.InfoLevel,
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
	zap.L().Info("zap started")

}

func loadServicePort() int {
	return viper.GetInt("serviceport")
}

func loadPinarySecret() string {
	val := viper.GetString("pinarysecret")
	return val
}
func loadPinaryClientID() string {
	val := viper.GetString("pinaryclientid")
	return val
}

func loadPinaryAuthUrl() string {
	val := viper.GetString("pinaryloginurl")
	return val
}
func loadPinaryDepositUrl() string {
	val := viper.GetString("pinarydepositurl")
	return val
}
func loadPinaryWithdrawUrl() string {
	val := viper.GetString("pinarywithdrawurl")
	return val
}

type DecimalCodec struct{}

func (dc *DecimalCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	// Use reflection to convert the reflect.Value to decimal.Decimal.
	dec, ok := val.Interface().(decimal.Decimal)
	if !ok {
		return fmt.Errorf("value %v to encode is not of type decimal.Decimal", val)
	}

	// Convert decimal.Decimal to primitive.Decimal128.
	primDec, err := primitive.ParseDecimal128(dec.String())
	if err != nil {
		return fmt.Errorf("error converting decimal.Decimal %v to primitive.Decimal128: %v", dec, err)
	}
	return vw.WriteDecimal128(primDec)
}

func (dc *DecimalCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	// Read primitive.Decimal128 from the ValueReader.
	primDec, err := vr.ReadDecimal128()
	if err != nil {
		return fmt.Errorf("error reading primitive.Decimal128 from ValueReader: %v", err)
	}

	// Convert primitive.Decimal128 to decimal.Decimal.
	dec, err := decimal.NewFromString(primDec.String())
	if err != nil {
		return fmt.Errorf("error converting primitive.Decimal128 %v to decimal.Decimal: %v", primDec, err)
	}

	// Set val to the decimal.Decimal value contained in dec.
	val.Set(reflect.ValueOf(dec))
	return nil
}

func loadMongoSetting() *Mgo {
	objFields := viper.GetStringMap("mongoconfig")
	hosts := viper.GetStringSlice("mongohosts")

	obj := &Mgo{
		Host:        hosts,
		AdminDbName: objFields["admindbname"].(string),
		MinPoolSize: int(objFields["minpoolsize"].(float64)),
		PoolLimit:   int(objFields["poollimit"].(float64)),
		UserName:    objFields["username"].(string),
		Password:    objFields["password"].(string),
		DbName:      objFields["dbname"].(string),
	}
	credential := options.Credential{
		Username:   obj.UserName,
		Password:   obj.Password,
		AuthSource: obj.AdminDbName,
	}
	opts := options.Client().
		SetHosts(obj.Host).
		SetAuth(credential).
		SetConnectTimeout(10 * time.Second).
		SetMinPoolSize(uint64(obj.MinPoolSize)).
		SetMaxPoolSize(uint64(obj.PoolLimit)).
		SetMaxConnIdleTime(time.Duration(3000) * time.Millisecond).
		SetRegistry(bson.NewRegistryBuilder().RegisterCodec(
			reflect.TypeOf(decimal.Decimal{}),
			&DecimalCodec{},
		).Build())

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		zap.L().Error("Ping the Mongodb error", zap.Error(err))
		panic(err)
	}
	obj.Client = client

	return obj
}

func loadMongoDbCollections() *MongoDbCollections {
	return &MongoDbCollections{
		BigThreeSummary:      "BigThreeSummary",
		BigThreeFutureOption: "BigThreeFutureOption",
		BigThreeFuture:       "BigThreeFuture",
		BigThreeOption:       "BigThreeOption",
		BigThreeCallPut:      "BigThreeCallPut",
		TopTradeFutureOI:     "TopTradeFutureOI",
		TopTradeOptionOI:     "TopTradeOptionOI",
		PutCallRatio:         "PutCallRatio",
	}
}

func loadListenSymbols() []string {
	allSymbols := make([]string, 0)
	allSymbols = append(allSymbols, viper.GetStringSlice("forexsymbols")...)
	allSymbols = append(allSymbols, viper.GetStringSlice("indexsymbols")...)
	//allSymbols = append(allSymbols, viper.GetStringSlice("futuresymbols")...)
	allSymbols = append(allSymbols, viper.GetStringSlice("stocksymbols")...)
	allSymbols = append(allSymbols, viper.GetStringSlice("coinsymbols")...)
	return allSymbols
}

func loadStockSymbols() []string {
	return viper.GetStringSlice("stocksymbols")
}

func loadForexSymbols() []string {
	return viper.GetStringSlice("forexsymbols")
}

func loadIndexSymbols() []string {
	return viper.GetStringSlice("indexsymbols")
}

func loadCryptoSymbols() []string {
	return viper.GetStringSlice("coinsymbols")
}

func loadQuoteServerConfig() *QuoteServerConfig {
	objFields := viper.GetStringMap("quoteserver")
	cfg := &QuoteServerConfig{
		Protocol: objFields["protocol"].(string),
		Port:     int(objFields["port"].(float64)),
		Name:     objFields["name"].(string),
		Address:  objFields["address"].(string),
	}
	//fmt.Println(cfg)
	return cfg
}

func loadBinanceQuoteServerConfig() *QuoteServerConfig {
	objFields := viper.GetStringMap("binancebridgeserver")
	cfg := &QuoteServerConfig{
		Protocol: objFields["protocol"].(string),
		Port:     int(objFields["port"].(float64)),
		Name:     objFields["name"].(string),
		Address:  objFields["address"].(string),
	}
	//fmt.Println(cfg)
	return cfg
}

func loadHMACKey() string {
	mackey := viper.GetString("hmackey")
	return mackey
}

func loadRedisConfig() *redis.Client {

	redisHost := viper.GetString("redishosts")
	redisDb := viper.GetInt("redisdb")
	return redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",      // no password set
		DB:       redisDb, // use default DB
	})

}

func loadEnableBinance() bool {
	return viper.GetBool("enablebinance")
}

func loadLandingAuthUrl() string {
	mackey := viper.GetString("landingauthurl")
	return mackey
}

func loadConsoleRedisConfig() *redis.Client {

	redisHost := viper.GetString("redishosts")
	redisDb := viper.GetInt("consoleredisdb")
	return redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",      // no password set
		DB:       redisDb, // use default DB
	})

}

func loadLineNotifyToken() string {
	token := viper.GetString("linenotifytoken")
	return token
}

func loadLineNotifyUrl() string {
	url := viper.GetString("linenotifyurl")
	return url
}

func loadLineNotifyEnable() bool {
	return viper.GetBool("enablelinenotify")
}

func encryptData(plainText string, key []byte, iv []byte) string {
	blk, err := aes.NewCipher(key)
	if err != nil {
		zap.L().Error("encrypt data error", zap.Error(err))
		return ""
	}

	blkMode := cipher.NewCBCEncrypter(blk, iv)
	plainTextBytes := []byte(plainText)

	inputPadding, err := pkcs7Pad(plainTextBytes, blk.BlockSize())

	encryptOutput := make([]byte, len(inputPadding))
	if err != nil {
		zap.L().Error("pad error", zap.Error(err))
		return ""
	}
	blkMode.CryptBlocks(encryptOutput, inputPadding)

	encryptedData := b64.RawURLEncoding.EncodeToString(encryptOutput)
	return encryptedData
}

func decryptData(encryptedb64String string, key []byte, iv []byte) string {
	encryptedBytes, err := b64.RawURLEncoding.DecodeString(encryptedb64String)
	if err != nil {
		zap.L().Error("decode encrypted base64 string error", zap.String("encstr", encryptedb64String), zap.Error(err))
		return ""
	}
	deBlk, err := aes.NewCipher(key)
	if err != nil {
		zap.L().Error("new aes256 cipher error", zap.Error(err))
		return ""
	}

	deMode := cipher.NewCBCDecrypter(deBlk, iv)

	decryptedBytes := make([]byte, len(encryptedBytes))

	deMode.CryptBlocks(decryptedBytes, encryptedBytes)
	//加密後的資料長度 - 最後一個Byte的數字就是 加密前的文字的長度
	decryptedBytesRealLength := len(encryptedBytes) - int(decryptedBytes[len(decryptedBytes)-1])

	plainText := string(decryptedBytes[:decryptedBytesRealLength])

	return plainText
}

func pkcs7Pad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		zap.L().Error("the block size is zero")
		return nil, errors.New("block size is zero")
	}
	if b == nil || len(b) == 0 {
		zap.L().Error("plain text is nil or null, or empty string")
		return nil, errors.New("input data (plain text) is nil or length is zero")
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}

func loadAuthShortTokenRegex() *regexp.Regexp {
	reg, err := regexp.Compile("^[0-9a-f]{22,32}$")
	if err != nil {
		panic(err)
	}
	return reg
}
