package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	dsn := "root:root@tcp(localhost:33062)/db_hd_wallet_saas?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		//logs.Log.Error("gorm open error: ", err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "utils/dbs/gorms/gorm-gen/query",
		Mode:              gen.WithoutContext | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})
	g.UseDB(db)

	dataMap := map[string]func(detailType string) (dataType string){
		"int":       func(detailType string) (dataType string) { return "int64" },
		"tinyint":   func(detailType string) (dataType string) { return "int64" },
		"smallint":  func(detailType string) (dataType string) { return "int64" },
		"mediumint": func(detailType string) (dataType string) { return "int64" },
		"bigint":    func(detailType string) (dataType string) { return "int64" },
	}
	g.WithDataTypeMap(dataMap)

	allModel := g.GenerateAllTable()
	g.ApplyBasic(allModel...)

	g.Execute()
}
