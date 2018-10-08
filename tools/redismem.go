package main

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/zhyhang/gofirst/tools/redis"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	total := 1000000
	idCount := 10
	parallel := 100
	today := time.Now().Format("20060102")
	rand.Seed(time.Now().Unix())
	pool6400 := redis.NewPool(":6400", 100)
	pool6401 := redis.NewPool(":6401", 100)
	defer pool6400.Close()
	defer pool6401.Close()
	log.Println("begin write to redis")
	batchWg := &sync.WaitGroup{}
	for i := 0; i < total; {
		for j := 0; j < parallel && i < total; j++ {
			uuid := uuid()
			for k := 0; k < idCount; k++ {
				idInt := rand.Intn(1000000)
				id := strconv.Itoa(idInt)
				field := ":" + today + ":" + id
				batchWg.Add(2)
				go incr(batchWg, pool6400, uuid+field)
				go zincr(batchWg, pool6401, uuid, field, 1.0)
			}
			i++
		}
		batchWg.Wait()
		//log.Printf("connect to redis times: %d",redis.ConnectCounter)
	}
	log.Println("end write to redis")
}

func uuid() string {
	randInt := rand.Int63()
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, uint64(randInt))
	md5Bytes := md5.Sum(intBytes)
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%x", md5Bytes)
	return buf.String()
}

func incr(wg *sync.WaitGroup, pool *redigo.Pool, key string) {
	conn := pool.Get()
	_, err := conn.Do("incr", key)
	if err != nil {
		log.Printf("incr error %v", err)
	}
	conn.Close()
	defer wg.Done()
}

func zincr(wg *sync.WaitGroup, pool *redigo.Pool, uid string, field string, delta float64) {
	conn := pool.Get()
	_, err := conn.Do("zincrby", uid, strconv.FormatFloat(delta, 'f', 2, 64), field)
	if err != nil {
		log.Printf("incr error %v", err)
	}
	conn.Close()
	defer wg.Done()
}
