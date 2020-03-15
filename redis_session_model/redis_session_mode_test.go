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
}

func assertIntEqual(t *testing.T, a, b int) {
    if a != b{
        t.Fatalf("%d doesn't equal to %d", a, b)
    }
}
