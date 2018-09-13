package main

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

const key = "test_key_1"

func main() {
	incrCount := 1000
	p := &redis.Pool{
		MaxIdle:   1,
		MaxActive: 100,
		Dial:      dial,
		Wait:      true,
	}
	defer p.Close()
	delConn := p.Get()
	delConn.Do("del", key)
	delConn.Close()
	waitCh := make(chan int, incrCount)
	for i := 0; i < incrCount; i++ {
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Microsecond)
		conn, err1 := p.GetContext(ctx)
		if err1 != nil {
			log.Printf("%s\n", err1.Error())
		}
		go incr(conn, waitCh)
	}
	for i := 0; i < incrCount; i++ {
		<-waitCh
	}
	reply, err2 := p.Get().Do("get", key)
	if err2 != nil {
		panic(err2)
	}
	count, _ := redis.Int64(reply, err2)
	fmt.Printf("final count %d", count)
}

func dial() (redis.Conn, error) {
	return redis.Dial("tcp", ":6400")
}

func incr(conn redis.Conn, waitCh chan int) {
	if conn != nil {
		conn.Do("incr", key)
		conn.Close()
	}
	waitCh <- 1
}
