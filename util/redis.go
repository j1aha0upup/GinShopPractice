package util

import (
	"fmt"

	"github.com/go-ini/ini"
	"github.com/gomodule/redigo/redis"
)

type RedisClient struct {
	redis.Conn
}

var Redis RedisClient

func init() {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("ini.Load(\"./config/config.ini\") error")
		panic(err)
	}
	cfgMysqlSection := cfg.Section("redis")
	DatebaseHost := cfgMysqlSection.Key("Host").String()
	DatebasePort := cfgMysqlSection.Key("Port").String()
	openStr := DatebaseHost + ":" + DatebasePort

	tempRedis, err := redis.Dial("tcp", openStr)
	Redis = RedisClient{tempRedis}
	if err != nil {
		fmt.Println("Redis open error")
		panic(err)
	}
}

func (redisClient RedisClient) SetRedisStringValue(key string, value string, expire int) error {
	redisClient.Do("SET", key, value)
	if expire != 0 {
		redisClient.Do("EXPIRE", key, expire)
	}

	return nil
}

func (redisClient RedisClient) GetRedisStringValue(key string) (string, error) {
	result, err := redis.String(redisClient.Do("GET", key))
	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}
