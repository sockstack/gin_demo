package redis

import (
	"cn.sockstack/gin_demo/pkg/config"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var (
	Redis redis.Conn
	err error
)

func init() {
	address := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)
	Redis, err = redis.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
}
