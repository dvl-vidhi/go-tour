package helper

import (
	"net/http/httptest"
	"testing"
)

func TestRespondWithError(t *testing.T) {
	statusCode := 500

	w := httptest.NewRecorder()

	RespondWithError(w, statusCode, "Internal Server Error")

	body_got := w.Body.String()
	body_want := `{"success":false,"success_msg":"Internal Server Error","data":{}}`

	assertBody(t, body_got, body_want)

	assertStatusCode(t, w.Code, statusCode)
}

func TestRespondWithJson(t *testing.T) {
	w := httptest.NewRecorder()

	data := map[string]string{"id": "abcd"}
	statusCode := 200

	RespondWithJson(w, statusCode, "Operation completed successfully", true, data)

	body_got := w.Body.String()
	body_want := `{"success":true,"success_msg":"Operation completed successfully","data":{"id":"abcd"}}`

	assertBody(t, body_got, body_want)
	assertStatusCode(t, w.Code, statusCode)
}

func assertStatusCode(t *testing.T, got_status int, want_status int) {
	t.Helper()

	if got_status != want_status {
		t.Errorf("got %d want %d", got_status, want_status)
	}
}

func assertBody(t *testing.T, body_got string, body_want string) {
	t.Helper()

	if body_got != body_want {
		t.Errorf("got %s want %s", body_got, body_want)
	}
}
