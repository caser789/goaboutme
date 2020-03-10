package main

type IUserModel interface {
    Create(username, password string) error
}

type User struct {
    model IUserModel
}

func (u *User) Create(username, password string) error {
    err := u.model.Create(username, password)
    if err != nil {
        return err
    }
    return nil
}
