package db

import (
	"log"

	"github.com/14jasimmtp/order-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct{
	DB *gorm.DB
}

func Connection(DB string) Handler{
	db,err:=gorm.Open(postgres.Open(DB),&gorm.Config{})
	if err != nil{
		log.Fatal("error in connecting db",err)
	}

	db.AutoMigrate(&models.Order{})

	return Handler{db}
}