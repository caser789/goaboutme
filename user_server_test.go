package main

import "testing"
import "net/http"
import "net/http/httptest"


type StubUser struct {
    createCalls []string
}

func (u *StubUser) Create(username, password string) error {
    u.createCalls = append(u.createCalls, username)
    return nil
}

func TestRoute(t *testing.T) {
    user := &StubUser{}
    server := NewUserServer(user)

    tests := []struct{
        name string
        method string
        url string
        statusCode int
        cookie *http.Cookie
    }{
        {
            name: "/user/register GET returns 200",
            method: http.MethodGet,
            url: "/user/register",
            statusCode: http.StatusOK,
        },
        {
            name: "/user/register POST returns 302", // if success
            method: http.MethodPost,
            url: "/user/register",
            statusCode: http.StatusFound,
        },
        {
            name: "/user/login GET returns 200",
            method: http.MethodGet,
            url: "/user/login",
            statusCode: http.StatusOK,
        },
        {
            name: "/user/login POST returns 302",
            method: http.MethodPost,
            url: "/user/login",
            statusCode: http.StatusFound,
        },
        {
            name: "/user/logout GET returns 302",
            method: http.MethodGet,
            url: "/user/logout",
            statusCode: http.StatusFound,
        },
        {
            name: "/user/profile GET without cookie returns 302",
            method: http.MethodGet,
            url: "/user/profile",
            statusCode: http.StatusFound,
        },
        {
            name: "/user/profile GET with cookie returns 200",
            method: http.MethodGet,
            url: "/user/profile",
            statusCode: http.StatusOK,
            cookie: &http.Cookie{Name: CookieKey, Value: "abc"},
        },
        {
            name: "/user/profile POST without cookie returns 302",
            method: http.MethodPost,
            url: "/user/profile",
            statusCode: http.StatusFound,
        },
        {
            name: "/user/profile POST with cookie returns 200",
            method: http.MethodPost,
            url: "/user/profile",
            statusCode: http.StatusOK,
            cookie: &http.Cookie{Name: CookieKey, Value: "abc"},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            request, _ := http.NewRequest(tt.method, tt.url, nil)
            if tt.cookie != nil {
                request.AddCookie(tt.cookie)
            }

            response := httptest.NewRecorder()

            server.ServeHTTP(response, request)

            assertStatus(t, response.Code, tt.statusCode)
        })
    }
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func TestRegisterPost(t *testing.T) {
    t.Run("test success", func(t *testing.T) {
        user := &StubUser{
            createCalls: []string{},
        }
        server := NewUserServer(user)

        method :=  http.MethodPost
        url := "/user/register"

        request, _ := http.NewRequest(method, url, nil)
        response := httptest.NewRecorder()
        server.ServeHTTP(response, request)

		if len(user.createCalls) != 1 {
			t.Fatalf("got %d calls to Create want %d", len(user.createCalls), 1)
		}
    })
}
