package main

import "html/template"
import "net/http"

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
    _ = r.ParseForm()
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
}

func HandleUserProfilePost(w http.ResponseWriter, r *http.Request) {
}
