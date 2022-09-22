package main

import (
	mail "email-service/emailModel"
	"email-service/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var con = service.Connection{}

func init() {
	con.Server = "mongodb://localhost:27017"
	con.Database = "EmailService"
	con.Collection = "email_data"

	con.Connect()
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid method")
		return
	}

	var mailDetails mail.Mail

	if err := json.NewDecoder(r.Body).Decode(&mailDetails); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}

	if len(mailDetails.MailSendTo) == 0 || mailDetails.MailBody == nil {
		respondWithError(w, http.StatusBadGateway, "Please enter emailTo and email body")
		return
	}

	if result, err := con.SendEmail(mailDetails); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, map[string]string{
			"message": result,
		})
	}
}

func searchFilter(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid method")
		return
	}

	var search mail.Search

	if err := json.NewDecoder(r.Body).Decode(&search); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}

	if result, err := con.SearchFilter(search); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result)
	}
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func main() {
	http.HandleFunc("/send-email", sendEmail)
	http.HandleFunc("/search", searchFilter)
	fmt.Println("Service Started at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
