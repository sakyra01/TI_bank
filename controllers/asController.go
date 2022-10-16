package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"poligon/decoder"
	"poligon/models"
)

// GetAs - API GET as by ID
var GetAs = func(w http.ResponseWriter, r *http.Request) {

	dbUri := models.GetDbUri()
	db, err := gorm.Open("postgres", dbUri)

	if err != nil {
		panic("failed to connect database")
	}

	params := mux.Vars(r)
	id := params["id"]

	var result models.As
	var as = db.First(&result, id)

	db.Close()

	res, err := json.Marshal(as)
	fmt.Println(string(res))
	fmt.Fprintf(w, string(res))
}

// AllAs - API GET all as
var AllAs = func(w http.ResponseWriter, r *http.Request) {

	dbUri := models.GetDbUri()
	db, err := gorm.Open("postgres", dbUri)
	db.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}

	var ass []models.As
	result := db.Find(&ass)

	db.Close()
	res, err := json.Marshal(result)
	fmt.Fprintf(w, string(res))
}

// GetSlug - API GET as by slug
var GetSlug = func(w http.ResponseWriter, r *http.Request) {
	dbUri := models.GetDbUri()
	db, err := gorm.Open("postgres", dbUri)
	db.LogMode(true)

	if err != nil {
		panic("failed to connect database")
	}

	params := mux.Vars(r)
	slug := params["slug"]

	var as = models.As{}
	model := db.Where("slug = ?", slug).First(&as)
	db.Close()

	res, err := json.Marshal(model)
	fmt.Println(string(res))
	fmt.Fprintf(w, string(res))

}

// StoreHostsForAs - store data from worker
var StoreHostsForAs = func(w http.ResponseWriter, r *http.Request) {

	/*-------Store and decode json------*/
	err := r.ParseForm()
	if err != nil {
		panic("err")
	}
	content := r.Form
	var targets []string
	var addr string

	for _, val := range content {
		addr = ""
		targets = []string{}
		for _, nal := range val {
			sub, hosts := decoder.Worker(nal)
			addr = sub
			targets = hosts
		}
	}

	dbUri := models.GetDbUri()
	db, err := gorm.Open("postgres", dbUri)
	db.LogMode(true)

	if err != nil {
		panic("failed to connect database")
	}

	params := mux.Vars(r)
	slug := params["slug"]

	var as = models.As{}
	db.Where("slug = ?", slug).First(&as)
	db.Save(&as)
	db.Close()

	//res, err := json.Marshal(as)
	//fmt.Println(string(res))

	sub := AddSub(addr, as)
	AddTargets(targets, sub)

}
