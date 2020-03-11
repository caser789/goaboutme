package user

const correctPassword = "1234"
const correctUserId = 1234

type StubUserModel struct {
    usernameToPassword map[string]string

    createCalls []string
    fromUserNameCalls []string
    getPasswordCalls []string
    getIdCalls []string
    getCalls []int
    getUsernameCalls []string
    getNicknameCalls []string
    getAvatarCalls []string
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
    return correctPassword
}

func (u *StubUserModel) GetId() int {
    u.getIdCalls = append(u.getIdCalls, "")
    return correctUserId
}

func (u *StubUserModel) Get(userId int) error {
    u.getCalls = append(u.getCalls, userId)
    return nil
}

func (u *StubUserModel) GetUsername() string {
    u.getUsernameCalls = append(u.getUsernameCalls, "")
    return "username"
}

func (u *StubUserModel) GetNickname() string {
    u.getNicknameCalls = append(u.getNicknameCalls, "")
    return "nickname"
}

func (u *StubUserModel) GetAvatar() []byte {
    u.getAvatarCalls = append(u.getAvatarCalls, "")
    return []byte{'a'}
}
