package initialize

import (
	"ccsu/global"
	"github.com/go-redis/redis"
)

func Redis() {
	r := global.Config.Redis
	c := redis.NewClient(&redis.Options{
		Addr: r.Addr,
		Password: r.Password,
		DB: r.DB,
	})

	pong, err := c.Ping().Result()

	if err != nil {
		global.Log.Error(err.Error())
	} else {
		global.Log.Info("redis connect ping response: " + pong)
		global.Redis = c
	}
}