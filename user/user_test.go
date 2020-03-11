package user

import "testing"

func TestUserAPIs(t *testing.T) {
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

    t.Run("test register fails if user already exists", func(t *testing.T) {
        userModel := &StubUserModel{
            usernameToPassword: map[string]string{"a": "1"},
        }
        sessionModel := &StubSessionModel{}
        user := &User{userModel, sessionModel}

        username := "a"
        password := "123456"
        err := user.Register(username, password)

        assertCalled(t, "UserModel.Create", len(userModel.createCalls))
        assertNotNil(t, err)
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

    t.Run("test GetProfile", func(t *testing.T) {
        userModel := &StubUserModel{
            usernameToPassword: map[string]string{},
        }
        sessionModel := &StubSessionModel{}
        user :=  &User{userModel, sessionModel}

        user.GetProfile()

        assertCalled(t, "UserModel.GetUsername", len(userModel.getUsernameCalls))
        assertCalled(t, "UserModel.GetNickname", len(userModel.getNicknameCalls))
        assertCalled(t, "UserModel.GetAvatar", len(userModel.getAvatarCalls))
    })

    t.Run("test UpdateProfile", func(t *testing.T) {
        userModel := &StubUserModel{
            usernameToPassword: map[string]string{},
        }
        sessionModel := &StubSessionModel{}
        user :=  &User{userModel, sessionModel}

        nickname := "xxx"
        avatar := []byte{'a', 'b'}
        user.UpdateProfile(nickname, avatar)

        assertCalled(t, "UserModel.SetNickname", len(userModel.setNicknameCalls))
        assertCalled(t, "UserModel.SetAvatar", len(userModel.setAvatarCalls))
    })

}

func TestUserIntegration(t *testing.T) {

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

func assertNotNil(t *testing.T, err error) {
    if err == nil {
        t.Fatalf("got nil error")
    }
}
