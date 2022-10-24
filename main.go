package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"poligon/controllers"
	"poligon/models"
)

func main() {
	router := mux.NewRouter()
	models.Init() //open for new db release

	router.HandleFunc("/api/v1/as", controllers.AllAs).Methods("GET")
	router.HandleFunc("/api/v1/as/{slug}", controllers.GetSlug).Methods("GET")
	router.HandleFunc("/api/v1/as/{slug}/targets", controllers.StoreHostsForAs).Methods("POST")
	router.HandleFunc("/api/v1/as/{slug}/targets", controllers.GetHostsForAs).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("frontend/dist")))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //localhost
	}

	fmt.Println(port)
	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8080/api
	if err != nil {
		fmt.Print(err)
	}
}
