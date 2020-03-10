package main

import "testing"

type StubModel struct {
    usernameToPassword map[string]string
}

func ( u *StubModel) Create(username, password string) error {
    u.usernameToPassword[username] = password
    return nil
}

func TestCreate(t *testing.T) {
    t.Run("test success", func(t *testing.T) {
        model := &StubModel{
            usernameToPassword: map[string]string{},
        }
        user := &User{model}

        username := "jiao.xue"
        password := "123456"
        user.Create(username, password)

        assertContains(t, model.usernameToPassword, username)
    })
}



func assertContains(t *testing.T, store map[string]string, key string) {
	t.Helper()
    _, ok := store[key]
	if !ok {
		t.Errorf("%v do not contains %v", store, key)
	}
}
