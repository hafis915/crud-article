package model

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=localhost user=postgres password=12345678 dbname=article port=5432 sslmode=disable TimeZone=Asia/Jakarta")

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true}), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(&Article{})
	if err != nil {
		log.Panicln(err)
	}

	return db, err
}
