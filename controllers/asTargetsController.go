package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"poligon/models"
)

var GetHostsForAs = func(w http.ResponseWriter, r *http.Request) {

	dbUri := models.GetDbUri()
	db, err := gorm.Open("postgres", dbUri)
	db.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}

	params := mux.Vars(r)
	slug := params["slug"]

	var as = models.As{}
	var subnets []models.Subnet
	var targets []models.Target

	db.Model(&as).Where("slug = ?", slug).First(&as)
	db.Model(&subnets).Where("as_id = ?", &as.ID).Find(&subnets)
	for _, sub := range subnets {
		db.Model(&targets).Where("subnet_id = ?", sub.ID).Find(&targets)

		res, _ := json.Marshal(targets)
		fmt.Println(string(res))
		fmt.Fprintf(w, string(res))
	}
	db.Close()
}
