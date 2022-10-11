package main

import (
	"bloodBank/model"
	"bloodBank/service"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var ser = service.Connection{}
var finalResponse model.Response

func init() {
	ser.Server = "mongodb://localhost:27017"
	ser.Database = "BloodBank"
	ser.Collection1 = "User"
	ser.Collection2 = "Donor"
	ser.Collection3 = "AvailableBlood"
	ser.Collection4 = "Patient"

	ser.Connect()
}

func saveUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var dataBody model.User
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.SaveUserDetails(dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func searchUsersById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		respondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
	}

	if result, err := ser.SearchUsersDetailsById(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func updateUserById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	var dataBody model.User
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.UpdateUserDetailsById(dataBody, id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func deleteUserById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "DELETE" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	if result, err := ser.DeleteUserDetailsById(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func saveDonor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var donorData model.Donor
	if err := json.NewDecoder(r.Body).Decode(&donorData); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.SaveDonorData(donorData); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func searchDonorById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		respondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
	}

	if result, err := ser.SearchDonorDetailsById(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func updateDonorById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	var dataBody model.Donor
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.UpdateDonorDetailsById(dataBody, id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func deleteDonorById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "DELETE" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	if result, err := ser.DeleteDonorDetailsById(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func bloodRequestPatient(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var dataBody model.Patient
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.ApplyBloodPatientDetails(dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func bloodProvidedPatient(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		respondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
	}

	if result, err := ser.GivenBloodPatientDetailsById(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func searchFilterBloodDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var bloodDetailsRequest model.AvailableBlood
	if err := json.NewDecoder(r.Body).Decode(&bloodDetailsRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	if result, err := ser.SearchFilterBloodDetails(bloodDetailsRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg}, "error")
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}, err string) {
	if err == "error" {
		finalResponse.Success = "false"
	} else {
		finalResponse.Success = "true"
	}
	finalResponse.SucessCode = fmt.Sprintf("%v", code)
	finalResponse.Response = payload
	response, _ := json.Marshal(finalResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	http.Handle("/save-user-details", ValidateJWT(saveUser))
	http.HandleFunc("/search-user-details-id/", searchUsersById)
	http.Handle("/update-user-details-id/", ValidateJWT(updateUserById))
	http.Handle("/delete-user-details-id/", ValidateJWT(deleteUserById))
	http.Handle("/save-donor-details", ValidateJWT(saveDonor))
	http.HandleFunc("/search-donor-details-id/", searchDonorById)
	http.Handle("/update-donor-details-id/", ValidateJWT(updateDonorById))
	http.Handle("/delete-donor-details-id/", ValidateJWT(deleteDonorById))
	http.Handle("/blood-request-patient-details", ValidateJWT(bloodRequestPatient))
	http.HandleFunc("/blood-provided-patient-details-id/", bloodProvidedPatient)
	http.HandleFunc("/search-filter-blood-details/", searchFilterBloodDetails)
	http.HandleFunc("/jwt", GetJwt)
	// http.Handle("/api", ValidateJWT(saveDonor))
	log.Println("Server started at 8080")
	http.ListenAndServe(":8080", nil)
}

func GetJwt(w http.ResponseWriter, r *http.Request) {
	// user := r.Header["User"]
	password := r.Header["Password"][0]
	userId := r.Header["Userid"][0]

	err := ser.AuthenticateUser(password, userId)

	if err == nil {

		token, err := CreateJWT(userId, password)
		if err != nil {
			return
		}
		fmt.Fprint(w, token)
	}
}

var SECRET = []byte("super-secret-auth-key")
var api_key = "1234"

type JWTClaim struct {
	Password string `json:"username"`
	UserId   string `json:"usedId"`
	jwt.StandardClaims
}

func CreateJWT(usedId string, password string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	token := jwt.New(jwt.SigningMethodHS256)

	// claims := token.Claims.(jwt.MapClaims)

	// claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims := &JWTClaim{
		UserId:   usedId,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.ParseWithClaims(r.Header["Token"][0], &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}
				return SECRET, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error()))
			}
			claims, ok := token.Claims.(*JWTClaim)
			if !ok {
				err = errors.New("couldn't parse claims")
				return
			}
			if claims.ExpiresAt < time.Now().Local().Unix() {
				err = errors.New("token expired")
				return
			}
			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
		}
	})
}
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "super secret area")
}
