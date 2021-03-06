package main

import "testing"
import "net/http"
import "net/http/httptest"

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
            cookie: &http.Cookie{Name: CookieKey, Value: "2345"},
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
            cookie: &http.Cookie{Name: CookieKey, Value: "1234"},
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
        user := &StubUser{}
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
        user := &StubUser{}
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
        user := &StubUser{}
        server := NewUserServer(user)

        method :=  http.MethodGet
        url := "/user/logout"

        request, _ := http.NewRequest(method, url, nil)
        request.AddCookie(&http.Cookie{Name: CookieKey, Value: "234"})

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


func TestProfile(t *testing.T) {
    t.Run("test get profile success", func(t *testing.T) {
        user := &StubUser{}
        server := NewUserServer(user)

        method :=  http.MethodGet
        url := "/user/profile"

        request, _ := http.NewRequest(method, url, nil)
        request.AddCookie(&http.Cookie{Name: CookieKey, Value: "123"})

        response := httptest.NewRecorder()
        server.ServeHTTP(response, request)

		if len(user.fromSessionIdCalls) != 1 {
			t.Fatalf("got %d calls to FromSessionId want %d", len(user.fromSessionIdCalls), 1)
		}

		if len(user.getProfileCalls) != 1 {
			t.Fatalf("got %d calls to GetProfile want %d", len(user.getProfileCalls), 1)
		}
    })
    t.Run("test post profile success", func(t *testing.T) {
        user := &StubUser{}
        server := NewUserServer(user)

        method :=  http.MethodPost
        url := "/user/profile"

        request, _ := http.NewRequest(method, url, nil)
        request.AddCookie(&http.Cookie{Name: CookieKey, Value: "888"})

        response := httptest.NewRecorder()
        server.ServeHTTP(response, request)

		if len(user.fromSessionIdCalls) != 1 {
			t.Fatalf("got %d calls to FromSessionId want %d", len(user.fromSessionIdCalls), 1)
		}

		if len(user.updateProfileCalls) != 1 {
			t.Fatalf("got %d calls to UpdateProfile want %d", len(user.updateProfileCalls), 1)
		}
    })
}
