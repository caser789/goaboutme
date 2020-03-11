package main

type IUserModel interface {
    Create(username, password string) error
}

type User struct {
    model IUserModel
}

func (u *User) Register(username, password string) error {
    err := u.model.Create(username, password)
    if err != nil {
        return err
    }
    return nil
}

func (u *User) Login(username, password string) (sessionId string, err error) {
    return "", nil
}

func (u *User) Logout() {}


func (u *User) FromSessionId(sessionId string) error {
    return nil
}

func (u *User) GetProfile() map[string]string {
    return nil
}

func (u *User) UpdateProfile(nickname string, avatar []byte) error {
    return nil
}
