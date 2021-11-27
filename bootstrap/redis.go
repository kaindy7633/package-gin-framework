package bootstrap

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/kaindy7633/package-gin-framework/global"
	"go.uber.org/zap"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Host + ":" + strconv.Itoa(global.App.Config.Redis.Port),
		Password: global.App.Config.Redis.Password, // no password set
		DB:       global.App.Config.Redis.DB,       // // use defeault DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.App.Log.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}

func String(i int) {
	panic("unimplemented")
}
