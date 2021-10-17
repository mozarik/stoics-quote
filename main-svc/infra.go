package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_Host     string
	DB_User     string
	DB_Password string
	DB_Name     string
	DB_Port     string
	DB_SSLMode  string
	DB_TimeZone string
}

func (config *ConfigDB) InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		config.DB_Host,
		config.DB_User,
		config.DB_Password,
		config.DB_Name,
		config.DB_Port,
		config.DB_SSLMode,
		config.DB_TimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
