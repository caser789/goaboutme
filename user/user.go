package user

import "errors"
import "encoding/base64"

type IUserModel interface {
    // Creater
    Create(username, password string) error
    FromUserName(username string) error
    Get(userId int) error

    // Getter
    GetPassword() string
    GetId() int
    GetUsername() string
    GetNickname() string
    GetAvatar() []byte

    // Setter
    SetNickname(nickname string) error
    SetAvatar(avatar []byte) error
}

type ISessionModel interface{
    Create(userId int) error
    GetId() int
    GetUserId() int
    Get(sessionId int) error
    Delete()
}

type User struct {
    UserModel IUserModel
    SessionModel ISessionModel
}

func (u *User) Register(username, password string) error {
    err := u.UserModel.Create(username, password)
    if err != nil {
        return errors.New("user exits")
    }
    return nil
}

func (u *User) Login(username, password string) (sessionId int, err error) {
    err = u.UserModel.FromUserName(username)
    if err != nil {
        return 0, errors.New("user exits")
    }

    if u.UserModel.GetPassword() != password {
        return 0, errors.New("password not match")
    }

    u.SessionModel.Create(u.UserModel.GetId())
    return u.SessionModel.GetId(), nil
}

func (u *User) Logout() {
    u.SessionModel.Delete()
}

func (u *User) FromSessionId(sessionId int) error {
    u.SessionModel.Get(sessionId)  // error
    userId := u.SessionModel.GetUserId()

    u.UserModel.Get(userId)
    return nil
}

func (u *User) GetProfile() map[string]string {
    return map[string]string {
        "username": u.UserModel.GetUsername(),
        "nickname": u.UserModel.GetNickname(),
        "avatar": base64.StdEncoding.EncodeToString(u.UserModel.GetAvatar()),
    }
}

func (u *User) UpdateProfile(nickname string, avatar []byte) error {
    u.UserModel.SetNickname(nickname)
    u.UserModel.SetAvatar(avatar)
    // TODO batch set
    return nil
}
