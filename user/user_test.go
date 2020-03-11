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

    t.Run("test login returns error with wrong password", func(t *testing.T) {
        userModel := &StubUserModel{
            usernameToPassword: map[string]string{},
        }
        sessionModel := &StubSessionModel{}
        user := &User{userModel, sessionModel}

        username := "jiao.xue"
        password := "888"

        user.Login(username, password)

        assertCalled(t, "UserModel.FromUserName", len(userModel.fromUserNameCalls))
        assertCalled(t, "UserModel.GetPassword", len(userModel.getPasswordCalls))
        assertNotCalled(t, "SessionModel.Create", len(sessionModel.createCalls))
    })

    t.Run("test FromSessionId", func(t *testing.T) {
        userModel := &StubUserModel{
            usernameToPassword: map[string]string{},
        }
        sessionModel := &StubSessionModel{}
        user :=  &User{userModel, sessionModel}

        user.FromSessionId(123)

        assertCalled(t, "SessionModel.Get", len(sessionModel.getCalls))
        assertCalled(t, "SessionModel.GetUserId", len(sessionModel.getUserIdCalls))
        assertCalled(t, "UserModel.Get", len(userModel.getCalls))
    })

    t.Run("test Logout", func(t *testing.T) {
        userModel := &StubUserModel{
            usernameToPassword: map[string]string{},
        }
        sessionModel := &StubSessionModel{}
        user :=  &User{userModel, sessionModel}

        user.Logout()

        assertCalled(t, "SessionModel.Delete", len(sessionModel.deleteCalls))
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

func assertNotCalled(t *testing.T, name string, count int) {
    if count != 0 {
        t.Fatalf("got %d calls to %s Create want %d", count, name, 0)
    }
}
