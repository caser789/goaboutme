package main

import "testing"
import "net/http"
import "net/http/httptest"


func TestRoute(t *testing.T) {
    tests := []struct{
        name string
        method string
        url string
        handler func(http.ResponseWriter, *http.Request)
    }{
        {
            name: "GET returns 200",
            method: http.MethodGet,
            url: "/user/register",
            handler: HandleUserRegister,
        },
        {
            name: "POST returns 200",
            method: http.MethodPost,
            url: "/user/register",
            handler: HandleUserRegister,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            request, _ := http.NewRequest(tt.method, tt.url, nil)
            response := httptest.NewRecorder()

            tt.handler(response, request)

            assertStatus(t, response.Code, http.StatusOK)
        })
    }
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
