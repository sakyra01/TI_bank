package controllers

import (
	"github.com/jinzhu/gorm"
	"poligon/models"
)

func AddTargets(targets []string, sub models.Subnet) {

	dbUri := models.GetDbUri()
	db, err := gorm.Open("postgres", dbUri)
	db.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}

	for _, host := range targets {
		var target = models.Target{}
		db.FirstOrCreate(&target, models.Target{Addr: host, Subnet: sub})
		db.Save(&target)
	}
	db.Close()

}
