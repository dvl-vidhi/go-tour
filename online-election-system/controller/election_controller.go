package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-election-system/config"
	"online-election-system/dao"
	"online-election-system/helper"
	"online-election-system/model"
	"strings"
)

var ead = dao.ElectionDAO{}

func init() {
	ead.Server = "mongodb://localhost:27017/"
	ead.Database = config.APP_CONFIG.Database
	if config.APP_CONFIG.Environment == "test" {
		ead.Database = config.APP_CONFIG.TestDatabase
	}
	ead.Collection = "Election"

	ead.ElectionConnect()
}

func AddElection(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	var dataBody model.Election
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ead.ElectionInsert(dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record inserted successfully", true, result)
	}
}

func AddCandidate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	var dataBody model.CandidatesRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ead.AddCandidate(dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record inserted successfully", true, result)
	}
}

func VerifyCandidate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	// token := r.Header.Get("tokenid")
	// mail, role, err := validateToken(token)
	// if err != nil {
	// 	respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err), "Error Occurred")
	// 	return
	// }
	// if role != "Admin" {
	// 	respondWithError(w, http.StatusBadRequest, "Token is invalid as it's role is different", "Invalid")
	// 	return
	// }

	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	var dataBody model.CandidatesRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ead.VerifyCandidate(dataBody, id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Candidate verified successfully", true, result)
	}
}

func SearchOneElection(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
	}

	var dataBody model.Election
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	user, err := ead.ElectionFindById(id)

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record found successfully", true, user)
	}
}

func SearchMultipleElection(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	var dataBody model.Election
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	election, err := ead.FilterOnElectionDetails(dataBody)

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record found successfully", true, election)
	}
}

func DeactivateElection(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	// token := r.Header.Get("tokenid")
	// mail, role, err := validateToken(token)
	// if err != nil {
	// 	respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err), "Error Occurred")
	// 	return
	// }
	// if role != "Admin" {
	// 	respondWithError(w, http.StatusBadRequest, "Token is invalid as it's role is different", "Invalid")
	// 	return
	// }

	var dataBody model.Election
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	if result, err := ead.Update(dataBody, id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Election Status changed successfully", true, result)
	}
}

func UpdateElection(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	var dataBody model.Election
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if result, err := ead.Update(dataBody, id); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Election Updated successfully", true, result)
	}
}
