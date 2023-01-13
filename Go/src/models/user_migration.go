package models

import (
	"github.com/jinzhu/gorm"
	"github.com/kanikim/src/entity"
	"log"
)

func db_migrate() (*gorm.DB, error) {
	conn, err := connect_db()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	conn.AutoMigrate(entity.User{})
	log.Println("Migration has been processed")

	return conn, nil
}
