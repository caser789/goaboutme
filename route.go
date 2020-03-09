package main

import "html/template"
import "net/http"

const CookieKey = "sessionId"

func HandleUserRegister(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        handleUserRegisterGet(w, r)
    case "POST":
        handleUserRegisterPost(w, r)
    default:
        http.Error(w, "Method not supported", http.StatusInternalServerError)
    }
}

func handleUserRegisterGet(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/register.html")
    t.Execute(w, nil)
}

func handleUserRegisterPost(w http.ResponseWriter, r *http.Request) {
    // 1. Parse request
    _ = r.ParseForm() // TODO handle error
    _ = r.PostFormValue("username")
    _ = r.PostFormValue("password")

    // 2. TODO If user exists, return error
    // 3. TODO Create user
    http.Redirect(w, r, "/user/login", 302)
}

func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        handleUserLoginGet(w, r)
    case "POST":
        handleUserLoginPost(w, r)
    default:
        http.Error(w, "Method not supported", http.StatusInternalServerError)
    }
}

func handleUserLoginGet(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/login.html")
    t.Execute(w, nil)
}

func handleUserLoginPost(w http.ResponseWriter, r *http.Request) {
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


func HandleUserLogout(w http.ResponseWriter, r *http.Request) {
}

func HandleUserProfile(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        HandleUserProfileGet(w, r)
    case "POST":
        HandleUserProfilePost(w, r)
    default:
        http.Error(w, "Method not supported", http.StatusInternalServerError)
    }
}

func HandleUserProfileGet(w http.ResponseWriter, r *http.Request) {
    // 1. Get sessionid from cookie
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

func HandleUserProfilePost(w http.ResponseWriter, r *http.Request) {
}
