package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"online-election-system/controller"
	"strings"
	"testing"
)

func TestAddUser(t *testing.T) {
	//
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	// here we write our expected response, in this case, we return a
	// 	// JSON string which is typical when dealing with REST APIs
	// 	io.WriteString(w, "{ \"success\": true}")

	// }

	payload := strings.NewReader(`{
		"role": "user",
		"name": "Lorem5",
		"email": "Lorem5",
		"password": "Lorem",
		"phone_number": "Lorem",
		"personal_info": {
			"name": "Lorem5",
			"father_name": "Lorem5",
			"dob": "2016-04-08",
			"age": 21,
			"document_type": "Lorem",
			"address": {
				"street": "Lorem",
				"city": "Lorem",
				"state": "Lorem",
				"zip_code": "Lorem",
				"country": "Lorem"
			}
		},
		"uploaded_docs": {
			"document_type": "Lorem",
			"document_identification_no": "Lorem",
			"document_path": "C:/Users/Vidhi/Downloads/529_3_334359_1666357514_Databricks - Generic.pdf"
		}
	}`)

	req := httptest.NewRequest("POST", "/api/user/add", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	// handler(w, req)
	// http.NewServeMux().HandleFunc("localhost:8080/api/user/search/63635d75a68e40fe497eac67", controller.SearchOneUser)
	handler := http.HandlerFunc(controller.AddUser)

	handler.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"success": true}`
	// if w.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		w.Body.String(), expected)
	// }
}

func TestVerifyUser(t *testing.T) {
	//
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	// here we write our expected response, in this case, we return a
	// 	// JSON string which is typical when dealing with REST APIs
	// 	io.WriteString(w, "{ \"success\": true}")

	// }

	payload := strings.NewReader(`{
		"role": "admin",
		"name": "Lorem1update",
		"email": "Lorem1update",
		"is_verified": false
	}`)

	req := httptest.NewRequest("PUT", "/api/user/verify/6363c5ccdef093ff71212320", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	// handler(w, req)
	// http.NewServeMux().HandleFunc("localhost:8080/api/user/search/63635d75a68e40fe497eac67", controller.SearchOneUser)
	handler := http.HandlerFunc(controller.VerifyUser)

	handler.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"success": true}`
	// if w.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		w.Body.String(), expected)
	// }
}

func TestUpdateUser(t *testing.T) {
	//
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	// here we write our expected response, in this case, we return a
	// 	// JSON string which is typical when dealing with REST APIs
	// 	io.WriteString(w, "{ \"success\": true}")

	// }

	payload := strings.NewReader(`{
		"role": "admin",
		"name": "Lorem1update",
		"email": "Lorem1update",
		"password": "Lorem",
		"phone_number": "Lorem",
		"personal_info": {
			"name": "Lorem",
			"father_name": "Lorem",
			"dob": "2016-04-08",
			"age": 21,
			"voter_id": "Lorem",
			"document_type": "Lorem",
			"address": {
				"street": "Lorem",
				"city": "Lorem",
				"state": "Lorem",
				"zip_code": "Lorem",
				"country": "Lorem"
			}
		},
		"is_verified": true,
		"verified_by": {
			"_id": "6c6c2d84de0f09d5007bdef8",
			"name": "Lorem"
		},
		"uploaded_docs": {
			"document_type": "Lorem",
			"document_identification_no": "Lorem",
			"document_path": "C:/Users/Vidhi/Downloads/MicrosoftTeams-image.png"
		}
	}`)

	req := httptest.NewRequest("PUT", "/api/user/update/63635d75a68e40fe497eac67", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	// handler(w, req)
	// http.NewServeMux().HandleFunc("localhost:8080/api/user/search/63635d75a68e40fe497eac67", controller.SearchOneUser)
	handler := http.HandlerFunc(controller.UpdateUser)

	handler.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"success": true}`
	// if w.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		w.Body.String(), expected)
	// }
}

func TestSearchOneUser(t *testing.T) {
	//
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	// here we write our expected response, in this case, we return a
	// 	// JSON string which is typical when dealing with REST APIs
	// 	io.WriteString(w, "{ \"success\": true}")

	// }

	req := httptest.NewRequest("GET", "/api/user/search/63635d75a68e40fe497eac67", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	// handler(w, req)
	// http.NewServeMux().HandleFunc("localhost:8080/api/user/search/63635d75a68e40fe497eac67", controller.SearchOneUser)
	handler := http.HandlerFunc(controller.SearchOneUser)

	handler.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"success": true}`
	// if w.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		w.Body.String(), expected)
	// }
}

func TestSearchMultipleUser(t *testing.T) {
	//
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	// here we write our expected response, in this case, we return a
	// 	// JSON string which is typical when dealing with REST APIs
	// 	io.WriteString(w, "{ \"success\": true}")

	// }

	payload := strings.NewReader(`{
		"role": "user",
		"city": "Lorem",
		"state": "Lorem",
		"zip_code": "Lorem",
		"country": "Lorem",
		"is_verified": true
	}`)

	req := httptest.NewRequest("POST", "/api/user/search-by-filter", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	// handler(w, req)
	// http.NewServeMux().HandleFunc("localhost:8080/api/user/search/63635d75a68e40fe497eac67", controller.SearchOneUser)
	handler := http.HandlerFunc(controller.SearchMultipleUser)

	handler.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"success": true}`
	// if w.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		w.Body.String(), expected)
	// }
}

func TestDeleteUser(t *testing.T) {
	//
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	// here we write our expected response, in this case, we return a
	// 	// JSON string which is typical when dealing with REST APIs
	// 	io.WriteString(w, "{ \"success\": true}")

	// }

	req := httptest.NewRequest("DELETE", "/api/user/delete/6369cf4d1047aa38379d6206", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	// handler(w, req)
	// http.NewServeMux().HandleFunc("localhost:8080/api/user/search/63635d75a68e40fe497eac67", controller.SearchOneUser)
	handler := http.HandlerFunc(controller.DeleteUser)

	handler.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"success": true}`
	// if w.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		w.Body.String(), expected)
	// }
}

func TestAddElection(t *testing.T) {
	//
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	// here we write our expected response, in this case, we return a
	// 	// JSON string which is typical when dealing with REST APIs
	// 	io.WriteString(w, "{ \"success\": true}")

	// }

	payload := strings.NewReader(`{
		"location": "Lorem5",
		"election_date": "2023-04-08",
		"result_date": "2023-04-08",
		"result": "",
		"election_status": "pending"
	}`)

	req := httptest.NewRequest("POST", "/api/election/add", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	// handler(w, req)
	// http.NewServeMux().HandleFunc("localhost:8080/api/user/search/63635d75a68e40fe497eac67", controller.SearchOneUser)
	handler := http.HandlerFunc(controller.AddElection)

	handler.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"success": true}`
	// if w.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		w.Body.String(), expected)
	// }
}

func TestAddCandidate(t *testing.T) {
	//
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	// here we write our expected response, in this case, we return a
	// 	// JSON string which is typical when dealing with REST APIs
	// 	io.WriteString(w, "{ \"success\": true}")

	// }

	payload := strings.NewReader(`{
		"election_id": "6368b8691503ded80405a7a8",
		"name": "Lorem1",
		"user_id": "6369cf4d1047aa38379d6206",
		"vote_count": "",
		"vote_sign": "C:/Users/Vidhi/Downloads/529_3_334359_1666357514_Databricks - Generic.pdf"
	}`)

	req := httptest.NewRequest("PUT", "/api/candidate/add", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	// handler(w, req)
	// http.NewServeMux().HandleFunc("localhost:8080/api/user/search/63635d75a68e40fe497eac67", controller.SearchOneUser)
	handler := http.HandlerFunc(controller.AddCandidate)

	handler.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"success": true}`
	// if w.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		w.Body.String(), expected)
	// }
}

func TestVerifyCandidate(t *testing.T) {
	//
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	// here we write our expected response, in this case, we return a
	// 	// JSON string which is typical when dealing with REST APIs
	// 	io.WriteString(w, "{ \"success\": true}")

	// }

	payload := strings.NewReader(`{
		"role": "admin",
		"name": "Lorem1update",
		"email": "Lorem1update",
		"user_id": "63635d75a68e40fe497eac67",
		"nomination_status": "not verified"
	}`)

	req := httptest.NewRequest("PUT", "/api/candidate/verify/636893521d1a023b7e38aa50", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	// handler(w, req)
	// http.NewServeMux().HandleFunc("localhost:8080/api/user/search/63635d75a68e40fe497eac67", controller.SearchOneUser)
	handler := http.HandlerFunc(controller.VerifyCandidate)

	handler.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"success": true}`
	// if w.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		w.Body.String(), expected)
	// }
}

func TestUpdateElection(t *testing.T) {
	//
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	// here we write our expected response, in this case, we return a
	// 	// JSON string which is typical when dealing with REST APIs
	// 	io.WriteString(w, "{ \"success\": true}")

	// }

	payload := strings.NewReader(`{
		"election_date": "2022-04-08 ",
		"result_date": " 2022-04-08",
		"election_staus": "voting "
	}`)

	req := httptest.NewRequest("PUT", "/api/election/update/636893521d1a023b7e38aa50", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	// handler(w, req)
	// http.NewServeMux().HandleFunc("localhost:8080/api/user/search/63635d75a68e40fe497eac67", controller.SearchOneUser)
	handler := http.HandlerFunc(controller.UpdateElection)

	handler.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"success": true}`
	// if w.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		w.Body.String(), expected)
	// }
}

func TestSearchOneElection(t *testing.T) {
	//
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	// here we write our expected response, in this case, we return a
	// 	// JSON string which is typical when dealing with REST APIs
	// 	io.WriteString(w, "{ \"success\": true}")

	// }

	req := httptest.NewRequest("GET", "/api/election/search/6368b8691503ded80405a7a8", nil)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	// handler(w, req)
	// http.NewServeMux().HandleFunc("localhost:8080/api/user/search/63635d75a68e40fe497eac67", controller.SearchOneUser)
	handler := http.HandlerFunc(controller.SearchOneUser)

	handler.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"success": true}`
	// if w.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		w.Body.String(), expected)
	// }
}

func TestSearchMultipleElection(t *testing.T) {
	//
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	// here we write our expected response, in this case, we return a
	// 	// JSON string which is typical when dealing with REST APIs
	// 	io.WriteString(w, "{ \"success\": true}")

	// }

	payload := strings.NewReader(`{
		"location": "Lorem",
		"election_date": "2022-04-08 ",
		"result_date": " 2022-04-08",
		"result": "Lorem",
		"election_status": "Lorem"
}`)

	req := httptest.NewRequest("POST", "/api/user/search-by-filter", payload)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	// handler(w, req)
	// http.NewServeMux().HandleFunc("localhost:8080/api/user/search/63635d75a68e40fe497eac67", controller.SearchOneUser)
	handler := http.HandlerFunc(controller.SearchMultipleElection)

	handler.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"success": true}`
	// if w.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		w.Body.String(), expected)
	// }
}
