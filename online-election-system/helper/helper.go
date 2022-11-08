package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Response struct {
	Success    bool        `json:"success"`
	SuccessMsg string      `json:"success_msg"`
	Data       interface{} `json:"data,omitempty"`
}

func (r *Response) ToJson() []byte {
	json_data, _ := json.Marshal(r)
	return json_data
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, msg, false, map[string]string{})
}

func RespondWithJson(w http.ResponseWriter, code int, message string, success bool, payload interface{}) {

	resp := Response{Data: payload, Success: success, SuccessMsg: message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp.ToJson())
}

func UploadFile(path string, uploadPath string) (string, error) {
	err := os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	fileURL, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	segments := strings.Split(fileURL.Path, "/")
	fileName := segments[len(segments)-1]
	fileName = uploadPath + fileName
	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	resp, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Close()
	size, err := io.Copy(file, resp)
	defer file.Close()
	return "File Downloaded with size :" + fmt.Sprintf("%v", size), nil
}
