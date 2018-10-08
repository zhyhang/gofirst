package redis

import (
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

var ConnectCounter = 0
var incrMutex sync.Mutex

func NewPool(addr string, maxActive int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxActive,
		IdleTimeout: 240 * time.Second,
		MaxActive:   maxActive,
		Dial: func() (redis.Conn, error) {
			incrMutex.Lock()
			ConnectCounter++
			incrMutex.Unlock()
			return redis.Dial("tcp", addr)
		},
		Wait: true,
	}
}
