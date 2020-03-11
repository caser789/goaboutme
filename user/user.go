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
    userModel IUserModel
    sessionModel ISessionModel
}

func (u *User) Register(username, password string) error {
    err := u.userModel.Create(username, password)
    if err != nil {
        return errors.New("user exits")
    }
    return nil
}

func (u *User) Login(username, password string) (sessionId int, err error) {
    err = u.userModel.FromUserName(username)
    if err != nil {
        return 0, errors.New("user exits")
    }

    if u.userModel.GetPassword() != password {
        return 0, errors.New("password not match")
    }

    u.sessionModel.Create(u.userModel.GetId())
    return u.sessionModel.GetId(), nil
}

func (u *User) Logout() {
    u.sessionModel.Delete()
}

func (u *User) FromSessionId(sessionId int) error {
    u.sessionModel.Get(sessionId)  // error
    userId := u.sessionModel.GetUserId()

    u.userModel.Get(userId)
    return nil
}

func (u *User) GetProfile() map[string]string {
    return map[string]string {
        "username": u.userModel.GetUsername(),
        "nickname": u.userModel.GetNickname(),
        "avatar": base64.StdEncoding.EncodeToString(u.userModel.GetAvatar()),
    }
}

func (u *User) UpdateProfile(nickname string, avatar []byte) error {
    u.userModel.SetNickname(nickname)
    u.userModel.SetAvatar(avatar)
    // TODO batch set
    return nil
}
