package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nk-code-lab/EasyMoney/pkg/models"
	"github.com/nk-code-lab/EasyMoney/pkg/utils"
)

var NewBook models.User

func GetAllUser(w http.ResponseWriter, r *http.Request){
	newUsers, err := models.GetAllUser()
    if err != nil {
        http.Error(w, "Failed to retrieve user data", http.StatusInternalServerError)
        return
    }

    res, err := json.Marshal(newUsers)
    if err != nil {
        http.Error(w, "Failed to marshal user data", http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserByID(w http.ResponseWriter, r *http.Request){
	req := mux.Vars(r)
	userId := req["user_id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Printf("Error: error while parsing => %s", err)
	}
	userDetails, _ := models.GetUserByID(ID)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	createUser := &models.User{}
	utils.ParseBody(r, createUser)
	u := createUser.CreateUser()
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	req := mux.Vars((r))
	userId := req["user_id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Printf("Error: error while parsing => %s", err)
	}
	user := models.DeleteUser(ID)
	res, _ :=json.Marshal(user)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	req := mux.Vars(r)
	userId := req["user_id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if  err != nil {
		fmt.Printf("Error: error in parsing => %s", err)
	}
	userDetails, db := models.GetUserByID(ID)

	//TODO: Retrieve the user details from the database
	// if userDetails == nil {
	// 	http.Error(w, "User not found", http.StatusNotFound)
	// 	return
	// }

	// Update user details if provided in the request
	if updateUser.Name != ""{
		userDetails.Name = updateUser.Name
	}

	if updateUser.Email != ""{
		userDetails.Name = updateUser.Email
	}

	if updateUser.Mobile != 0{
		userDetails.Name = updateUser.Name
	}
	// Save to the database
	db.Save(&userDetails)
	
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}