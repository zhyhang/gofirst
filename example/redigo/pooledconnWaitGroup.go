package main

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"sync"
	"time"
)

const key1 = "test_key_1"

var poolConnWg sync.WaitGroup

func main() {
	incrCount := 1000
	p := &redis.Pool{
		MaxIdle:   1,
		MaxActive: 100,
		Dial:      dial1,
		Wait:      true,
	}
	defer p.Close()
	delConn := p.Get()
	delConn.Do("del", key1)
	delConn.Close()
	for i := 0; i < incrCount; i++ {
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Microsecond)
		conn, err1 := p.GetContext(ctx)
		if err1 != nil {
			log.Printf("%s\n", err1.Error())
		} else {
			poolConnWg.Add(1)
			go incr1(conn)
		}
	}
	poolConnWg.Wait()
	reply, err2 := p.Get().Do("get", key1)
	if err2 != nil {
		panic(err2)
	}
	count, _ := redis.Int64(reply, err2)
	fmt.Printf("final count %d", count)
}

func dial1() (redis.Conn, error) {
	return redis.Dial("tcp", ":6400")
}

func incr1(conn redis.Conn) {
	conn.Do("incr", key1)
	conn.Close()
	defer poolConnWg.Done()
}
