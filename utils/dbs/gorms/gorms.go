package gorms

import (
	"gin-boilerplate/utils/dbs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// InitSimpleClient generates a simple gorm client
func InitSimpleClient(cfg *gorm.Config) *gorm.DB {
	//dbCfg := configs.GetGlobalDbConfig()
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", dbCfg.Usr, dbCfg.Pwd, dbCfg.Host, dbCfg.Port, dbCfg.DbName)
	dsn := "root:pawword@tcp(127.0.0.1:3306)/test?charset=utf8mb4"
	gormDB, err := gorm.Open(mysql.Open(dsn), cfg)

	if err != nil {
		//logs.Log.Error("gorm open error: ", err)
	}

	return gormDB
}

// InitGormClient generates a gorm client.
// If existingDb is true, it will use the existing db connection. Or it will create a new db connection.
func InitGormClient(existingDb bool, customCfg bool) *gorm.DB {
	if existingDb {
		gormClient := InitSimpleClient(getGormConfig(customCfg))
		return gormClient
	} else {
		gormClient, err := gorm.Open(mysql.New(mysql.Config{
			Conn: dbs.NativeClient,
		}), getGormConfig(customCfg))
		if err != nil {
			//logs.Log.Panicln("Init gormClient failed!")
		}

		return gormClient
	}
}

// GetGormConfig returns custom config if customCfg is true, or returns default config.
func getGormConfig(customCfg bool) *gorm.Config {
	if customCfg {
		return &gorm.Config{
			NowFunc: func() time.Time {
				//return time.Now().Local()
				return time.Now().UTC() // using UTC time
			},
			//Logger: logger.New(log.New(logs.Log.Out, "\r\n", log.LUTC), logger.Config{
			//	LogLevel: logger.Error,
			//}),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			SkipDefaultTransaction: false,
		}
	} else {
		return &gorm.Config{}
	}
}

//func InitClient(opt gorm.Option) *gorm.DB {
//}
