package main

import "net/http"
import "log"
import "github.com/caser789/goaboutme/user"

func main() {
    UserModel := &user.StubUserModel{}
    SessionModel := &user.StubSessionModel{}
    user := &user.User{
        UserModel: UserModel,
        SessionModel: SessionModel,
    }

    server := NewUserServer(user)

    if err := http.ListenAndServe(":5000", server); err != nil {
        log.Fatalf("could not listen on port 5000 %v", err)
    }
}
