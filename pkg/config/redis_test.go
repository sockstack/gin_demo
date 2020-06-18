package config

import (
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	Redis = (&redis{}).Load("../../conf/redis.ini").Init()
	fmt.Println(Redis)
	if Mysql == nil {
		t.Fail()
	}
}