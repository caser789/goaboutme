package user

import "errors"

var userId = 1000
var usernameToPassword = map[string]string{}
var usernameToUserId = map[string]int{}
var usernameToNickname = map[string]string{}
var usernameToAvatar = map[string][]byte{}
var userIdToUsername = map[int]string{}

type StubUserModel struct {
    username string
    password string
    nickname string
    userId int
    avatar []byte

    createCalls []string
    fromUserNameCalls []string
    getPasswordCalls []string
    getIdCalls []string
    getCalls []int
    getUsernameCalls []string
    getNicknameCalls []string
    getAvatarCalls []string
    setAvatarCalls []string
    setNicknameCalls []string
}

func (u *StubUserModel) Create(username, password string) error {
    u.createCalls = append(u.createCalls, username)
    _, ok := usernameToPassword[username]
    if ok {
        return errors.New("user exits")
    }

    userId += 1
    usernameToPassword[username] = password
    usernameToUserId[username] = userId
    userIdToUsername[userId] = username

    u.username = username
    u.password = password
    u.userId = userId
    return nil
}

func (u *StubUserModel) FromUserName(username string) error {
    u.fromUserNameCalls = append(u.fromUserNameCalls, username)
    password, ok := usernameToPassword[username]
    if !ok {
        return errors.New("user not exits")
    }

    u.username = username
    u.password = password
    u.userId = usernameToUserId[username]
    u.nickname, _ = usernameToNickname[username]
    u.avatar, _ = usernameToAvatar[username]
    return nil
}

func (u *StubUserModel) GetPassword() string {
    u.getPasswordCalls = append(u.getPasswordCalls, "")
    return u.password
}

func (u *StubUserModel) GetId() int {
    u.getIdCalls = append(u.getIdCalls, "")
    return u.userId
}

func (u *StubUserModel) Get(userId int) error {
    u.getCalls = append(u.getCalls, userId)
    username, ok := userIdToUsername[userId]
    if !ok {
        return errors.New("user ID not exits")
    }
    u.username = username
    u.password = usernameToPassword[username]
    u.userId = usernameToUserId[username]
    u.nickname, _ = usernameToNickname[username]
    u.avatar, _ = usernameToAvatar[username]
    return nil
}

func (u *StubUserModel) GetUsername() string {
    u.getUsernameCalls = append(u.getUsernameCalls, "")
    return u.username
}

func (u *StubUserModel) GetNickname() string {
    u.getNicknameCalls = append(u.getNicknameCalls, "")
    return u.nickname
}

func (u *StubUserModel) GetAvatar() []byte {
    u.getAvatarCalls = append(u.getAvatarCalls, "")
    return u.avatar
}

func (u *StubUserModel) SetAvatar(avatar []byte) error {
    u.setAvatarCalls = append(u.setAvatarCalls, "")
    usernameToAvatar[u.username] = avatar
    u.avatar = avatar
    return nil
}

func (u *StubUserModel) SetNickname(nickname string) error {
    u.setNicknameCalls = append(u.setNicknameCalls, nickname)
    usernameToNickname[u.username] = nickname
    u.nickname = nickname
    return nil
}
