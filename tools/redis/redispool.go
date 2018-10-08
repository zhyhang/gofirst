package redis

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/jolestar/go-commons-pool"
	"log"
	"sync"
	"time"
)

var ConnectCounter = 0
var incrMutex sync.Mutex

type ConnPool interface {
	Borrow() redis.Conn
	Return(conn redis.Conn)
	Close()
}

type RedigoPool struct {
	realPool *redis.Pool
}

func (p *RedigoPool) Borrow() (c redis.Conn) {
	c1 := p.realPool.Get()
	return c1
}

func (p *RedigoPool) Return(c redis.Conn) {
	c.Close()
}

func (p *RedigoPool) Close() {
	p.realPool.Close()
}

type CommonPool struct {
	realPool *pool.ObjectPool
}

func (p *CommonPool) Borrow() (c redis.Conn) {
	c1, err := p.realPool.BorrowObject(context.Background())
	if err != nil {
		log.Printf("borrrow from common pool error %v\n", err)
	}
	return c1.(redis.Conn)
}

func (p *CommonPool) Return(c redis.Conn) {
	p.realPool.ReturnObject(context.Background(), c)
}

func (p *CommonPool) Close() {
	p.realPool.Close(context.Background())
}

func NewRedigoPool(addr string, maxActive int) *RedigoPool {
	p := &redis.Pool{
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
	return &RedigoPool{p}
}

func NewCommonPool(addr string, maxActive int) *CommonPool {
	factory := pool.NewPooledObjectFactorySimple(
		func(context.Context) (interface{}, error) {
			incrMutex.Lock()
			ConnectCounter++
			incrMutex.Unlock()
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				log.Printf("create common pool's redis connection error %v\n", err)
			}
			return conn, nil
		})
	ctx := context.Background()
	p := pool.NewObjectPoolWithDefaultConfig(ctx, factory)
	p.Config.MaxTotal = maxActive
	p.Config.MaxIdle = 10
	return &CommonPool{p}
}
