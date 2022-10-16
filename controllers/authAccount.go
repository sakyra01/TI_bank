package controllers

//
//import (
//	"encoding/json"
//	"net/http"
//	"IBank/models"
//	u "IBank/utils"
//)
//
//var CreateUser = func(w http.ResponseWriter, r *http.Request) {
//
//	User := &models.User{}
//	err := json.NewDecoder(r.Body).Decode(User) //decode the request body into struct and failed if any error occur
//	if err != nil {
//		u.Respond(w, u.Message(false, "Invalid request"))
//		return
//	}
//
//	resp := User.Create() //Create User
//	u.Respond(w, resp)
//}
//
//var Authenticate = func(w http.ResponseWriter, r *http.Request) {
//
//	User := &models.User{}
//	err := json.NewDecoder(r.Body).Decode(User) //decode the request body into struct and failed if any error occur
//	if err != nil {
//		u.Respond(w, u.Message(false, "Invalid request"))
//		return
//	}
//
//	resp := models.Login(User.Login, User.Password)
//	u.Respond(w, resp)
//}
