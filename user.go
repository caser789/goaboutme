package main

type User struct {
    model *UserModel
}

func (u *User) Create(username, password string) error {
    err := u.model.Create(username, password)
    if err != nil {
        return err
    }
    return nil
}
