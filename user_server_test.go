package main

import "testing"
import "net/http"
import "net/http/httptest"


type StubUser struct {
    registerCalls []string
    loginCalls []string
    logoutCalls []string
    fromSessionIdCalls []string
}

func (u *StubUser) Register(username, password string) error {
    u.registerCalls = append(u.registerCalls, username)
    return nil
}

func (u *StubUser) Login(username, password string) (sessionId string, err error) {
    u.loginCalls = append(u.loginCalls, username)
    return "", nil
}

func (u *StubUser) Logout() {
    u.logoutCalls = append(u.logoutCalls, "")
}

func (u *StubUser) FromSessionId(sessionId string) error {
    u.fromSessionIdCalls = append(u.fromSessionIdCalls, sessionId)
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
    t.Run("test register success", func(t *testing.T) {
        user := &StubUser{
            registerCalls: []string{},
        }
        server := NewUserServer(user)

        method :=  http.MethodPost
        url := "/user/register"

        request, _ := http.NewRequest(method, url, nil)
        response := httptest.NewRecorder()
        server.ServeHTTP(response, request)

		if len(user.registerCalls) != 1 {
			t.Fatalf("got %d calls to Register want %d", len(user.registerCalls), 1)
		}
    })
}

func TestLogin(t *testing.T) {
    t.Run("test login success", func(t *testing.T) {
        user := &StubUser{
            registerCalls: []string{},
        }
        server := NewUserServer(user)

        method :=  http.MethodPost
        url := "/user/login"

        request, _ := http.NewRequest(method, url, nil)
        response := httptest.NewRecorder()
        server.ServeHTTP(response, request)

		if len(user.loginCalls) != 1 {
			t.Fatalf("got %d calls to Login want %d", len(user.loginCalls), 1)
		}
    })
}

func TestLogout(t *testing.T) {
    t.Run("test logout success", func(t *testing.T) {
        user := &StubUser{
            registerCalls: []string{},
        }
        server := NewUserServer(user)

        method :=  http.MethodPost
        url := "/user/logout"

        request, _ := http.NewRequest(method, url, nil)
        request.AddCookie(&http.Cookie{Name: CookieKey, Value: "abc"})

        response := httptest.NewRecorder()
        server.ServeHTTP(response, request)

		if len(user.fromSessionIdCalls) != 1 {
			t.Fatalf("got %d calls to FromSessionId want %d", len(user.fromSessionIdCalls), 1)
		}

		if len(user.logoutCalls) != 1 {
			t.Fatalf("got %d calls to Logout want %d", len(user.logoutCalls), 1)
		}
    })
}
