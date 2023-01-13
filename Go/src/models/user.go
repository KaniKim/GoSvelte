package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var config = ConfigDB{}

type ConfigDB struct {
	User     string
	Password string
	Host     string
	Port     string
	DB       string
}

func (config *ConfigDB) Read() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	config.Host = os.Getenv("Host")
	config.Password = os.Getenv("Password")
	config.User = os.Getenv("User")
	config.Port = os.Getenv("Port")
	config.DB = os.Getenv("DB")
}

func connect_db() (*gorm.DB, error) {
	config.Read()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.User, config.Password, config.Host, config.Port, config.DB)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
