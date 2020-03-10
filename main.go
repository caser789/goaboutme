package main

import "net/http"
import "log"

func main() {
    userModel := &UserModel{}
    user := &User{userModel}
	server := NewUserServer(user)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
