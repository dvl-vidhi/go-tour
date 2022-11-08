package controller

import (
	"net/http/httptest"
	"testing"
)

func TestAddElection(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/user/add", nil)
		response := httptest.NewRecorder()

		AddUser(response, request)
		got := response.Body.String()
		want := "Method not allowed\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}

func TestAddCandidate(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/user/verify/", nil)
		response := httptest.NewRecorder()

		AddUser(response, request)
		got := response.Body.String()
		want := "Method not allowed\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}

func TestVerifyCandidate(t *testing.T) {
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

func TestUpdateElection(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/user/update/", nil)
		response := httptest.NewRecorder()

		AddUser(response, request)
		got := response.Body.String()
		want := "Method not allowed\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}

func TestSearchOneElection(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/api/user/search/", nil)
		response := httptest.NewRecorder()

		AddUser(response, request)
		got := response.Body.String()
		want := "Method not allowed\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}

func TestSearchMultipleElection(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/user/search-by-filter", nil)
		response := httptest.NewRecorder()

		AddUser(response, request)
		got := response.Body.String()
		want := "Method not allowed\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}

func TestDeactivateElection(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/user/delete/", nil)
		response := httptest.NewRecorder()

		AddUser(response, request)
		got := response.Body.String()
		want := "Method not allowed\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
