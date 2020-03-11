package user

import "testing"

func TestUser(t *testing.T) {
    t.Run("test register success", func(t *testing.T) {
        userModel := &StubUserModel{
            usernameToPassword: map[string]string{},
        }
        sessionModel := &StubSessionModel{}
        user := &User{userModel, sessionModel}

        username := "jiao.xue"
        password := "123456"
        user.Register(username, password)

        assertContains(t, userModel.usernameToPassword, username)
        assertCalled(t, "UserModel.Create", len(userModel.createCalls))
    })

    t.Run("test login success", func(t *testing.T) {
        userModel := &StubUserModel{
            usernameToPassword: map[string]string{},
        }
        sessionModel := &StubSessionModel{}
        user := &User{userModel, sessionModel}

        username := "jiao.xue"

        user.Login(username, correctPassword)

        assertCalled(t, "UserModel.FromUserName", len(userModel.fromUserNameCalls))
        assertCalled(t, "UserModel.GetPassword", len(userModel.getPasswordCalls))
        assertCalled(t, "SessionModel.Create", len(sessionModel.createCalls))
    })

}

func assertContains(t *testing.T, store map[string]string, key string) {
	t.Helper()
    _, ok := store[key]
	if !ok {
		t.Errorf("%v do not contains %v", store, key)
	}
}

func assertCalled(t *testing.T, name string, count int) {
    if count != 1 {
        t.Fatalf("got %d calls to %s Create want %d", count, name, 1)
    }
}
