package dbs

import (
	"context"
	"database/sql"
	"gin-boilerplate/utils/configs"
	"gin-boilerplate/utils/dbs/gorms"
	"gorm.io/gorm"
)

var NativeClient *sql.DB
var GormClient *gorm.DB

func InitNativeDBClient() {
	NativeClient = initNativeDBClient()
}

func initNativeDBClient() *sql.DB {
	//dbCfg := configs.GetGlobalDbConfig()
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", dbCfg.Usr, dbCfg.Pwd, dbCfg.Host, dbCfg.Port, dbCfg.DbName)
	//db, err := sql.Open(dbCfg.Driver, dsn)
	dbClient, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4")
	if err != nil {
		panic(err)
	}

	return dbClient
}

func InitGormClient() {
	usingExistDb := configs.GetBoolByKey("gorm.using-exist-db")
	usingCustomGormCfg := configs.GetBoolByKey("gorm.custom-gorm-cfg")
	if usingExistDb {
		GormClient = gorms.InitGormClient(usingCustomGormCfg, NativeClient)
	} else {
		GormClient = gorms.InitSimpleClient(usingCustomGormCfg)
	}

}

func GormWithContext(ctx context.Context) *gorm.DB {
	return GormClient.WithContext(ctx)
}
