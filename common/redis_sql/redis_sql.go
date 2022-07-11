package redis_sql

import (
	"fmt"
	"gin-app-example/common/util"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var RedisClient *redis.Client

func init() {
	envConfig := util.GetEnvConfig()

	strRedisConnect := fmt.Sprintf("%s:%d", envConfig.Redis.Address, envConfig.Redis.Port)
	redisClient := redis.NewClient(&redis.Options{
		Addr:         strRedisConnect,
		PoolSize:     1000,
		Password:     envConfig.Redis.Password,
		ReadTimeout:  time.Second * time.Duration(3),
		WriteTimeout: time.Second * time.Duration(3),
		IdleTimeout:  time.Minute * time.Duration(5),
	})
	_, errRedis := redisClient.Ping().Result()
	if errRedis != nil {
		logrus.Printf("redis error  %v", errRedis)
		panic("init redis error")
	} else {
		logrus.Println("init redis ok")
	}

	RedisClient = redisClient
}
