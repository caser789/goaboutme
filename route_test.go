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
            name: "/user/register GET returns 200",
            method: http.MethodGet,
            url: "/user/register",
            handler: HandleUserRegister,
        },
        {
            name: "/user/register POST returns 200",
            method: http.MethodPost,
            url: "/user/register",
            handler: HandleUserRegister,
        },
        {
            name: "/user/login GET returns 200",
            method: http.MethodGet,
            url: "/user/login",
            handler: HandleUserLogin,
        },
        {
            name: "/user/login POST returns 200",
            method: http.MethodPost,
            url: "/user/login",
            handler: HandleUserLogin,
        },
        {
            name: "/user/logout GET returns 200",
            method: http.MethodGet,
            url: "/user/logout",
            handler: HandleUserLogout,
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
