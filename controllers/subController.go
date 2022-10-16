package controllers

import (
	"github.com/jinzhu/gorm"
	"poligon/models"
)

func AddSub(addr string, as models.As) (sub models.Subnet) {

	dbUri := models.GetDbUri()
	db, err := gorm.Open("postgres", dbUri)
	db.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}
	var subnet = models.Subnet{}

	db.FirstOrCreate(&subnet, models.Subnet{Address: addr, As: as})
	db.Save(&subnet)
	db.Close()
	return subnet
}
