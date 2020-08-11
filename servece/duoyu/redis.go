package duoyu

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var Pool *redis.Pool

func InitPool(address string, maxIdle int, maxActive int, idleTimeout time.Duration) {
	Pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
