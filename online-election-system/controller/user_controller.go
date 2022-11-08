package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-election-system/auth"
	"online-election-system/config"
	"online-election-system/dao"
	"online-election-system/helper"
	"online-election-system/model"
	"strings"
)

// var Database = []byte(config.APP_CONFIG.Database)

var uad = dao.UserDAO{}

func init() {
	uad.Server = "mongodb://localhost:27017/"
	uad.Database = config.APP_CONFIG.Database
	if config.APP_CONFIG.Environment == "test" {
		uad.Database = config.APP_CONFIG.TestDatabase
	}
	uad.Collection = "User"

	uad.UserConnect()
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var dataBody model.User
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if data, err := uad.UserInsert(dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid request")
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record inserted successfully", true, data)
	}
}

func VerifyUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	var dataBody model.User
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := uad.VerifyUser(dataBody, id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid request")
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "User verified successfully", true, result)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	var dataBody model.User
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := uad.Update(dataBody, id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "User Updated successfully", true, result)
	}
}

func SearchOneUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
	}

	// var dataBody model.User
	// if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
	// 	helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
	// 	return
	// }

	user, err := uad.UserFindById(id)

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid request")
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record found successfully", true, user)
	}
}

func SearchMultipleUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var dataBody model.UserFilter
	// if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
	// 	helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
	// 	return
	// }

	user, err := uad.FilterOnUsersDetails(dataBody)

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid request")
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record found successfully", true, user)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "DELETE" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	if result, err := uad.UserDelete(id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, result, true, nil)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var dataBody model.User
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	users, err := uad.FindByEmailAndPassword(dataBody.Email, dataBody.Password)

	if err != nil {
		http.Error(w, "No user found", http.StatusNotFound)
		return
	}

	user := users[0]

	fmt.Println(user)

	claim := auth.User_Claims{Username: user.Email, Authorized: true, Role: user.Role}

	jwt, err := auth.GenerateJWT(claim)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Token", jwt)

	helper.RespondWithJson(w, http.StatusOK, "Logged in successfully", true, nil)
}
