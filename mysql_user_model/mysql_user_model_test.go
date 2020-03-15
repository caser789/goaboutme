package mysql_user_model

import "testing"
import "math/rand"
import "fmt"
import "bytes"

func TestMysqlUserModel(t *testing.T) {
    t.Run("test Create success", func(t *testing.T) {
        n := rand.Intn(10000)
        username := fmt.Sprintf("xue%d", n)
        password := "1111"
        model := &MysqlUserModel{}

        err := model.Create(username, password)

        assertNil(t, err)

        model.Delete()
    })

    t.Run("test Create failed", func(t *testing.T) {
        n := rand.Intn(10000)
        username := fmt.Sprintf("xue%d", n)
        password := "1111"
        model := &MysqlUserModel{}

        model.Create(username, password)
        err := model.Create(username, password)

        assertNotNil(t, err)

        model.Delete()
    })

    t.Run("test FromUsername failed", func(t *testing.T) {
        n := rand.Intn(10000)
        username := fmt.Sprintf("xue%d", n)

        model := &MysqlUserModel{}
        err := model.FromUsername(username)

        assertNotNil(t, err)
    })

    t.Run("test FromUsername", func(t *testing.T) {
        n := rand.Intn(10000)
        username := fmt.Sprintf("xue%d", n)
        password := "1111"
        model := &MysqlUserModel{}

        model.Create(username, password)

        model = &MysqlUserModel{}
        err := model.FromUsername(username)

        assertNil(t, err)
        model.Delete()
    })

    t.Run("test Get error", func(t *testing.T) {
        userId := 2
        model := &MysqlUserModel{}
        err := model.Get(userId)
        assertNotNil(t, err)
    })

    t.Run("test Get OK", func(t *testing.T) {
        n := rand.Intn(10000)
        username := fmt.Sprintf("xue%d", n)
        password := "1111"
        model := &MysqlUserModel{}

        model.Create(username, password)
        userId := model.id

        model = &MysqlUserModel{}
        err := model.Get(userId)
        assertNil(t, err)

        model.Delete()
    })

    t.Run("test SetNickname", func(t *testing.T) {
        n := rand.Intn(10000)
        username := fmt.Sprintf("xue%d", n)
        password := "1111"
        model := &MysqlUserModel{}
        model.Create(username, password)
        userId := model.id

        model = &MysqlUserModel{}
        _ = model.Get(userId)

        nickname := "lalal"
        model.SetNickname(nickname)

        model = &MysqlUserModel{}
        _ = model.Get(userId)
        assertEqual(t, model.nickname, nickname)

        model.Delete()
    })

    t.Run("test SetAvatar", func(t *testing.T) {
        n := rand.Intn(10000)
        username := fmt.Sprintf("xue%d", n)
        password := "1111"
        model := &MysqlUserModel{}
        model.Create(username, password)
        userId := model.id

        model = &MysqlUserModel{}
        _ = model.Get(userId)

        avatar := []byte{'a', 'b'}
        model.SetAvatar(avatar)

        model = &MysqlUserModel{}
        _ = model.Get(userId)
        assertNotNil(t, model.avatar)
        assertAvatarEqual(t, model.avatar, avatar)

        model.Delete()
    })
}


func assertNotNil(t *testing.T, err interface{}) {
    if err == nil {
        t.Fatalf("got nil error")
    }
}

func assertNil(t *testing.T, err error) {
    if err != nil {
        t.Fatalf("got error %v expecting nil", err)
    }
}

func assertEqual(t *testing.T, a, b string) {
    if a != b{
        t.Fatalf("%v doesn't equal to %v", a, b)
    }
}

func assertAvatarEqual(t *testing.T, a, b []byte) {
    res := bytes.Compare(a, b)
    if res != 0 {
        t.Fatalf("%v doesn't equal to %v", a, b)
    }
}
