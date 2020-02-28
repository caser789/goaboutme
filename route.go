package main

import "net/http"

func handleUserRegister(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        handleUserRegisterGet(w, r)
    }
}

func handleUserRegisterGet(w http.ResponseWriter, r *http.Request) {
}
