package helper

import (
	"encoding/json"
	"net/http"
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
