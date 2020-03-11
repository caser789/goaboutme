package user

import "errors"

type IUserModel interface {
    Create(username, password string) error
    FromUserName(username string) error
    GetPassword() string
    GetId() int
}

type ISessionModel interface{
    Create(userId int) error
    GetId() int
}

type User struct {
    userModel IUserModel
    sessionModel ISessionModel
}

func (u *User) Register(username, password string) error {
    err := u.userModel.Create(username, password)
    if err != nil {
        return err
    }
    return nil
}

func (u *User) Login(username, password string) (sessionId int, err error) {
    u.userModel.FromUserName(username)
    // TODO not exists

    if u.userModel.GetPassword() != password {
        return 0, errors.New("password not match")
    }

    u.sessionModel.Create(u.userModel.GetId())
    return u.sessionModel.GetId(), nil
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
