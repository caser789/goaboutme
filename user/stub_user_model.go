package user

type StubUserModel struct {
    usernameToPassword map[string]string

    createCalls []string
    fromUserNameCalls []string
    getPasswordCalls []string
    getIdCalls []string
}

func (u *StubUserModel) Create(username, password string) error {
    u.usernameToPassword[username] = password
    u.createCalls = append(u.createCalls, username)
    return nil
}

func (u *StubUserModel) FromUserName(username string) error {
    u.fromUserNameCalls = append(u.fromUserNameCalls, username)
    return nil
}

func (u *StubUserModel) GetPassword() string {
    u.getPasswordCalls = append(u.getPasswordCalls, "")
    return "123"
}

func (u *StubUserModel) GetId() int {
    u.getIdCalls = append(u.getIdCalls, "")
    return 123
}
