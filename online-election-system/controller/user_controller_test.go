package controller

import (
	"net/http/httptest"
	"testing"
)

func TestAddUser(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/add-user", nil)
		response := httptest.NewRecorder()

		AddUser(response, request)
		got := response.Body.String()
		want := "Method not allowed\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}

func TestVerifyUser(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/verify-user", nil)
		response := httptest.NewRecorder()

		AddUser(response, request)
		got := response.Body.String()
		want := "Method not allowed\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}

func TestSearchOneUser(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/search-one-user", nil)
		response := httptest.NewRecorder()

		AddUser(response, request)
		got := response.Body.String()
		want := "Method not allowed\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}

func TestSearchMultipleUser(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/search-multiple-user", nil)
		response := httptest.NewRecorder()

		AddUser(response, request)
		got := response.Body.String()
		want := "Method not allowed\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/delete-user", nil)
		response := httptest.NewRecorder()

		AddUser(response, request)
		got := response.Body.String()
		want := "Method not allowed\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
