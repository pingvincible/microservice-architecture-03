package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/pingvincible/crud-service/pkg/models"
	"github.com/pingvincible/crud-service/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewUser models.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	userDetails, _ := models.GetUserById(ID)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	userDetails, db := models.GetUserById(ID)
	if updateUser.UserName != "" {
		userDetails.UserName = updateUser.UserName
	}
	if updateUser.FirstName != "" {
		userDetails.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != "" {
		userDetails.LastName = updateUser.LastName
	}
	if updateUser.Email != "" {
		userDetails.Email = updateUser.Email
	}
	if updateUser.Phone != "" {
		userDetails.Phone = updateUser.Phone
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
