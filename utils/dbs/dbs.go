package dbs

import (
	"database/sql"
)

var NativeDB *sql.DB

func InitNativeDBClient() *sql.DB {
	//dbCfg := configs.GetGlobalDbConfig()
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", dbCfg.Usr, dbCfg.Pwd, dbCfg.Host, dbCfg.Port, dbCfg.DbName)
	//db, err := sql.Open(dbCfg.Driver, dsn)
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4")
	if err != nil {
		panic(err)
	}

	return db
}
