package main

import "html/template"
import "net/http"
// import "bytes"
// import "io"

const CookieKey = "sessionId"

type UserServer struct{
    http.Handler
}

func NewUserServer() *UserServer {
	p := new(UserServer)

	router := http.NewServeMux()
	router.Handle("/user/register", http.HandlerFunc(p.handleUserRegister))
	router.Handle("/user/login", http.HandlerFunc(p.handleUserLogin))
	router.Handle("/user/logout", http.HandlerFunc(p.handleUserLogout))
	router.Handle("/user/profile", http.HandlerFunc(p.handleUserProfile))

	p.Handler = router
	return p
}

func (u *UserServer) handleUserRegister(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        u.handleUserRegisterGet(w, r)
    case "POST":
        u.handleUserRegisterPost(w, r)
    default:
        http.Error(w, "Method not supported", http.StatusInternalServerError)
    }
}

func (u *UserServer) handleUserRegisterGet(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/register.html")
    t.Execute(w, nil)
}

func (u *UserServer) handleUserRegisterPost(w http.ResponseWriter, r *http.Request) {
    // 1. Parse request
    _ = r.ParseForm() // TODO handle error
    _ = r.PostFormValue("username")
    _ = r.PostFormValue("password")

    // 2. TODO If user exists, return error
    // 3. TODO Create user
    http.Redirect(w, r, "/user/login", 302)
}

func (u *UserServer) handleUserLogin(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        u.handleUserLoginGet(w, r)
    case "POST":
        u.handleUserLoginPost(w, r)
    default:
        http.Error(w, "Method not supported", http.StatusInternalServerError)
    }
}

func (u *UserServer) handleUserLoginGet(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/login.html")
    t.Execute(w, nil)
}

func (u *UserServer) handleUserLoginPost(w http.ResponseWriter, r *http.Request) {
    // 1. Parse request
    _ = r.ParseForm() // TODO handle error

    _ = r.PostFormValue("username")
    _ = r.PostFormValue("password")

    // 2. Get user by username (200 and 400)
    // 3. Auth user against password (200 and 400)
    // 4. Create or update session
    // 5. render with cookie
    cookie := http.Cookie{
        Name: CookieKey,
        Value: "xxx",
        HttpOnly: true,
    }
    http.SetCookie(w, &cookie)
    http.Redirect(w, r, "/user/profile", 302)
}


func (u *UserServer) handleUserLogout(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie(CookieKey)
	if err != http.ErrNoCookie {
        // expire session
	}
    http.Redirect(w, r, "/user/login", 302)
}

func (u *UserServer) handleUserProfile(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        u.handleUserProfileGet(w, r)
    case "POST":
        u.handleUserProfilePost(w, r)
    default:
        http.Error(w, "Method not supported", http.StatusInternalServerError)
    }
}

func (u *UserServer) handleUserProfileGet(w http.ResponseWriter, r *http.Request) {
    // 1. Get sessionId from cookie
	_, err := r.Cookie(CookieKey)
    if err != nil {
        http.Redirect(w, r, "/user/login", 302)
        return
    }

    // 2. Get session by sesionid (200 or 404)
    // 3. Get user by session
    // 4. render the profile
    t, _ := template.ParseFiles("templates/profile.html")
    t.Execute(w, nil)
}

func (u *UserServer) handleUserProfilePost(w http.ResponseWriter, r *http.Request) {
    // 1. Get sessionId from cookie
	_, err := r.Cookie(CookieKey)
    if err != nil {
        http.Redirect(w, r, "/user/login", 302)
    }

    // 2. Get session by sesionid (200 or 404)
    // 3. Get user by session
    // 4. Parse file
    r.ParseMultipartForm(10485760) // max body in memory is 10MB
    file, _, _ := r.FormFile("avatar")
    if file != nil {
        defer file.Close()
    }
    // buf := bytes.NewBuffer(nil)
    // io.Copy(buf, file)
    // _ = r.PostFormValue("nickname")
    // 5. Update profile
    // 6. render profile
    t, _ := template.ParseFiles("templates/profile.html")
    t.Execute(w, nil)
}
