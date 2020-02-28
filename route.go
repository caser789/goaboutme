package main

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
}

func handleUserRegisterPost(w http.ResponseWriter, r *http.Request) {
}

func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        handleUserLoginGet(w, r)
    default:
        http.Error(w, "Method not supported", http.StatusInternalServerError)
    }
}

func handleUserLoginGet(w http.ResponseWriter, r *http.Request) {
}
