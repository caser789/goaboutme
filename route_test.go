package main

import "testing"
import "net/http"
import "net/http/httptest"

func TestRegistUser(t *testing.T) {
    t.Run("GET returns 200", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/user/register", nil)
        response := httptest.NewRecorder()

        handleUserRegister(response, request)

        assertStatus(t, response.Code, http.StatusOK)
    })
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
