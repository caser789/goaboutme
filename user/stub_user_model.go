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
