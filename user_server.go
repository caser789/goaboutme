package main

import "html/template"
import "net/http"
import "bytes"
import "io"
import "strconv"

const CookieKey = "sessionId"

type IUser interface {
    Register(username, password string) error
    Login(username, password string) (sessionId int, err error)
    FromSessionId(sessionId int) error
    Logout()
    GetProfile() map[string]string
    UpdateProfile(nickname string, avatar []byte) error
}

type UserServer struct{
    user IUser
    http.Handler
}

func NewUserServer(user IUser) *UserServer {
	p := new(UserServer)

	router := http.NewServeMux()
	router.Handle("/user/register", http.HandlerFunc(p.handleUserRegister))
	router.Handle("/user/login", http.HandlerFunc(p.handleUserLogin))
	router.Handle("/user/logout", http.HandlerFunc(p.handleUserLogout))
	router.Handle("/user/profile", http.HandlerFunc(p.handleUserProfile))

	p.Handler = router
    p.user = user
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
    _ = r.ParseForm()
    username := r.PostFormValue("username")
    password := r.PostFormValue("password")

    // 2. Create user
    _ = u.user.Register(username, password)
    // TODO user exists error

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
    _ = r.ParseForm() // TODO handle error
    username := r.PostFormValue("username")
    password := r.PostFormValue("password")

    sessionId, _ := u.user.Login(username, password)
    // 1. TODO User not found error
    // 2. TODO Password not right error

    cookie := http.Cookie{
        Name: CookieKey,
        Value: string(sessionId),
        HttpOnly: true,
    }
    http.SetCookie(w, &cookie)
    http.Redirect(w, r, "/user/profile", 302)
}


func (u *UserServer) handleUserLogout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(CookieKey)

	if err != http.ErrNoCookie {
        sessionId, err := strconv.Atoi(cookie.Value)
        u.user.FromSessionId(sessionId)
        u.user.Logout()
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
	cookie, err := r.Cookie(CookieKey)
    if err != nil {
        http.Redirect(w, r, "/user/login", 302)
        return
    }

    sessionId, _ := strconv.Atoi(cookie.Value)
    u.user.FromSessionId(sessionId)
    // TODO session expires

    profile := u.user.GetProfile()
    // TODO Profile type

    t, _ := template.ParseFiles("templates/profile.html")
    t.Execute(w, profile)
}

func (u *UserServer) handleUserProfilePost(w http.ResponseWriter, r *http.Request) {
    // 1. Get sessionId from cookie
	cookie, err := r.Cookie(CookieKey)
    if err != nil {
        http.Redirect(w, r, "/user/login", 302)
        return
    }

    sessionId, _ := strconv.Atoi(cookie.Value)
    u.user.FromSessionId(sessionId) // TODO user not exists

    r.ParseMultipartForm(10485760) // max body in memory is 10MB
    file, _, _ := r.FormFile("avatar")
    buf := bytes.NewBuffer(nil)

    if file != nil {
        defer file.Close()
        io.Copy(buf, file)
    }

    avatar := buf.Bytes()
    nickname := r.PostFormValue("nickname")
    u.user.UpdateProfile(nickname, avatar)

    t, _ := template.ParseFiles("templates/profile.html")
    t.Execute(w, nil)
}
