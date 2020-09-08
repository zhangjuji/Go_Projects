package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func initPool(address string, maxIdle int, maxActive int, idleTimeout time.Duration) {

	pool = &redis.Pool{
		MaxIdle:     maxIdle,     // 最大空闲链接数
		MaxActive:   maxActive,   // 和数据库的最大链接数, 0表示没有限制
		IdleTimeout: idleTimeout, // 表示最大空闲处
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
