package config

import (
	"github.com/akar016012/golang-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

  func Connect()  {
	db, err := gorm.Open(postgres.Open("postgres://postgres:secret@localhost:5432/postgres"),&gorm.Config{})

	if err != nil{
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
  }