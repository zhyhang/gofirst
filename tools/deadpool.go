package main

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jolestar/go-commons-pool"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"sync"
	"time"
)

func main() {
	// dump stack info by http://localhost:6060/debug/pprof/goroutine?debug=2
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	// create commons pool
	factory := pool.NewPooledObjectFactorySimple(
		func(context.Context) (interface{}, error) {
			return redis.Dial("tcp", ":6400")
		})
	ctx := context.TODO()
	p := pool.NewObjectPoolWithDefaultConfig(ctx, factory)
	p.Config.MaxTotal = 100
	p.Config.MaxIdle = 100
	defer p.Close(ctx)
	// write to redis parallel
	total := 100000
	idCount := 10
	// when <= 100 (MaxTotal) is OK, or else dead lock
	parallel := 200
	today := time.Now().Format("20060102")
	rand.Seed(time.Now().Unix())
	log.Println("begin write to redis")
	for i := 0; i < total; {
		batchWg := &sync.WaitGroup{}
		for j := 0; j < parallel/idCount && i < total; j++ {
			uuid := createUuid()
			for k := 0; k < idCount; k++ {
				idInt := rand.Intn(1000000)
				id := strconv.Itoa(idInt)
				field := ":" + today + ":" + id
				batchWg.Add(1)
				go incrCounter(batchWg, p, uuid+field)
			}
			i++
		}
		batchWg.Wait()
	}
	log.Println("end write to redis")
}

func incrCounter(wg *sync.WaitGroup, pool *pool.ObjectPool, key string) {
	defer wg.Done()
	ctx := context.TODO()
	conn, _ := pool.BorrowObject(ctx)
	if conn != nil {
		defer pool.ReturnObject(ctx, conn)
	} else {
		return
	}
	redisConn := conn.(redis.Conn)
	_, err := redisConn.Do("incr", key)
	if err != nil {
		log.Printf("incr error %v", err)
	}
}

func createUuid() string {
	randInt := rand.Int63()
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, uint64(randInt))
	md5Bytes := md5.Sum(intBytes)
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%x", md5Bytes)
	return buf.String()
}
