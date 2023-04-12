package database

import (
	"fmt"
	"github.com/runan13/echo-rest-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(config *config.Config) *gorm.DB {
	postgresDsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul", config.DB.Host, config.DB.User, config.DB.Password, config.DB.Name, config.DB.Port)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  postgresDsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
