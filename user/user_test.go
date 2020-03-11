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

		if len(userModel.createCalls) != 1 {
			t.Fatalf("got %d calls to Create want %d", len(userModel.createCalls), 1)
		}
    })

    t.Run("test login success", func(t *testing.T) {
        userModel := &StubUserModel{
            usernameToPassword: map[string]string{},
        }
        sessionModel := &StubSessionModel{}
        user := &User{userModel, sessionModel}

        username := "jiao.xue"
        password := "123456"
        user.Login(username, password)
    })

}

func assertContains(t *testing.T, store map[string]string, key string) {
	t.Helper()
    _, ok := store[key]
	if !ok {
		t.Errorf("%v do not contains %v", store, key)
	}
}
