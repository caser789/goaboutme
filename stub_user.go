package main

type StubUser struct {
    registerCalls []string
    loginCalls []string
    logoutCalls []string
    fromSessionIdCalls []int
    getProfileCalls []string
    updateProfileCalls []string
}

func (u *StubUser) Register(username, password string) error {
    u.registerCalls = append(u.registerCalls, username)
    return nil
}

func (u *StubUser) Login(username, password string) (sessionId int, err error) {
    u.loginCalls = append(u.loginCalls, username)
    return 0, nil
}

func (u *StubUser) Logout() {
    u.logoutCalls = append(u.logoutCalls, "")
}

func (u *StubUser) FromSessionId(sessionId int) error {
    u.fromSessionIdCalls = append(u.fromSessionIdCalls, sessionId)
    return nil
}

func (u *StubUser) GetProfile() map[string]string {
    u.getProfileCalls = append(u.getProfileCalls, "")
    return nil
}

func (u *StubUser) UpdateProfile(nickname string, avatar []byte) error {
    u.updateProfileCalls = append(u.updateProfileCalls, "")
    return nil
}
