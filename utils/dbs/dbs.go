package dbs

import (
	"context"
	"database/sql"
	"fmt"
	"gin-boilerplate/utils/configs"
	"gin-boilerplate/utils/dbs/gorms"
	"gorm.io/gorm"
)

var nativeClient *sql.DB
var gormClient *gorm.DB

func InitNativeDBClient() {
	nativeClient = initNativeDBClient()
}

func initNativeDBClient() *sql.DB {
	dbCfg := configs.GetGlobalDbConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", dbCfg.Usr, dbCfg.Pwd, dbCfg.Host, dbCfg.Port, dbCfg.DbName)
	dbClient, err := sql.Open(dbCfg.Driver, dsn)
	if err != nil {
		panic(err)
	}

	return dbClient
}

func InitGormClient() {
	usingExistDb := configs.GetBoolByKey("gorm.using-exist-db")
	usingCustomGormCfg := configs.GetBoolByKey("gorm.custom-gorm-cfg")
	if usingExistDb {
		nativeClient = initNativeDBClient()
		gormClient = gorms.InitGormClient(usingCustomGormCfg, nativeClient)
	} else {
		gormClient = gorms.InitSimpleClient(usingCustomGormCfg)
	}
}

// WithCustomConnectionPool sets custom connection pool config
func WithCustomConnectionPool(db *sql.DB) {
	db.SetConnMaxLifetime(0)
	db.SetMaxOpenConns(1000)
	db.SetConnMaxIdleTime(10)
	//db.SetConnMaxIdleTime()
}

func GormWithContext(ctx context.Context) *gorm.DB {
	return gormClient.WithContext(ctx)
}

// InitGlobalDBClient initializes database client, it will switch to use different database client according to the config
func InitGlobalDBClient() {
	switch configs.GetGlobalDbConfig().Client {
	case "native":
		InitNativeDBClient()
	case "gorm":
		InitGormClient()
	}
}

func GormClient() *gorm.DB {
	return gormClient
}

func NativeClient() *sql.DB {
	return nativeClient
}
