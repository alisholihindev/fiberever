package lib

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func DBEstablish() *gorm.DB {
	dsn := GetConfig("DB_USER") + ":" + GetConfig("DB_PASS") + "@tcp(" + GetConfig("DB_HOST") + ":" + GetConfig("DB_PORT") + ")/" + GetConfig("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		CommonLogger().Error(err)
		panic("failed to connect database")
	}

	return db
}

func DBConnection() *gorm.DB {
	return DBConn
}
