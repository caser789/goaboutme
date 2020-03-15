package redis_session_model

import "github.com/go-redis/redis"
import "time"
import "fmt"

type RedisSessionModel struct {
    sessionId int
    userId int
}


func (r *RedisSessionModel) Create(userId int) error {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
    result, err := client.Incr("session_id").Result()
    sessionId := int(result)
    if err != nil {
        panic(err)
    }

    err = client.Set(getUserIdToSessionIdKey(userId), sessionId, time.Hour).Err()
    if err != nil {
        panic(err)
    }

    err = client.Set(getSessionIdToUserIdKey(sessionId), userId, time.Hour).Err()
    if err != nil {
        panic(err)
    }
    r.sessionId = sessionId
    r.userId = userId

    return err
}

func (r *RedisSessionModel) Get(sessionId int) error {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
    userId, err := client.Get(getSessionIdToUserIdKey(sessionId)).Int()
    if err != nil {
        panic(err)
    }
    r.userId = userId
    r.sessionId = sessionId

    return nil
}

func (r *RedisSessionModel) Delete() {
}

func (r *RedisSessionModel) GetId() int {
    return r.sessionId
}

func (r *RedisSessionModel) GetUserId() int {
    return r.userId
}

func getUserIdToSessionIdKey(userId int) string {
    return fmt.Sprintf("user.%d.session", userId)
}

func getSessionIdToUserIdKey(sessionId int) string {
    return fmt.Sprintf("session.%d.user", sessionId)
}
