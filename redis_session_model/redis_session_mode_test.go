package redis_session_model

import "testing"


func TestRedisSessionModel(t *testing.T) {
    t.Run("test Create", func(t *testing.T) {
        model := RedisSessionModel{}
        userId := 100
        model.Create(userId)
    })

    t.Run("test Get", func(t *testing.T) {
        model := RedisSessionModel{}
        userId := 100
        model.Create(userId)
        sessionId := model.sessionId

        model = RedisSessionModel{}
        model.Get(sessionId)

        assertIntEqual(t, sessionId, model.sessionId)
        assertIntEqual(t, userId, model.userId)
    })

    t.Run("test Delete", func(t *testing.T) {
        model := RedisSessionModel{}
        userId := 100
        model.Create(userId)
        sessionId := model.sessionId

        model.Delete()

        model = RedisSessionModel{}
        err := model.Get(sessionId)
        assertNotNil(t, err)
    })
}

func assertIntEqual(t *testing.T, a, b int) {
    if a != b{
        t.Fatalf("%d doesn't equal to %d", a, b)
    }
}

func assertNotNil(t *testing.T, err interface{}) {
    if err == nil {
        t.Fatalf("got nil error")
    }
}

