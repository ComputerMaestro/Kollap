package database

import (
	"Collap/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	GetDb() *gorm.DB
}

func InitializeDB(conf *config.Db) *gorm.DB {
	dsn := "host=" + conf.Host +
		" user=" + conf.User +
		" password=" + conf.Password +
		" dbname=" + conf.DBName +
		" port=" + conf.Port +
		" sslmode=" + conf.SSLMode +
		" TimeZone=" + conf.TimeZone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to establish connection to db")
	}
	return db
}
