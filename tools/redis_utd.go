package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"github.com/zhyhang/gofirst/tools/redis"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	// dump stack info by http://localhost:6060/debug/pprof/goroutine?debug=2
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	if len(os.Args) < 5 {
		log.Printf("using: redis_utd <total request> <fields per key> <parallism> <redis addr>")
		log.Printf("\tredis addr: host:port")
		os.Exit(-1)
	}
	total, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	idCount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Panic(err)
	}
	parallel, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Panic(err)
	}

	redisAddr := os.Args[4]

	//total := 100
	//idCount := 20
	//parallel := 100
	//redisAddr :="192.168.144.50:10300"

	today := time.Now().Format("20060102")
	rand.Seed(time.Now().Unix())
	//ssds10300 := redis.NewRedigoPool(":6401", 100)
	ssds10300 := redis.NewCommonPool(redisAddr, parallel)
	defer ssds10300.Close()
	log.Println("begin write to redis")
	for i := 0; i < total; {
		batchWg := &sync.WaitGroup{}
		for j := 0; j < parallel/idCount && i < total; j++ {
			uuid := uuidUtd()
			for k := 0; k < idCount; k++ {
				idInt := rand.Intn(1000000)
				id := strconv.Itoa(idInt)
				field := ":" + today + ":" + id
				batchWg.Add(1)
				go zaddUtd(batchWg, ssds10300, uuid, field, 1.0)
			}
			i++
		}
		batchWg.Wait()
	}
	log.Println("end write to redis")
}

func uuidUtd() string {
	randInt := rand.Int63()
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, uint64(randInt))
	return fmt.Sprintf("%x", md5.Sum(intBytes))
}

func zaddUtd(wg *sync.WaitGroup, pool redis.ConnPool, uid string, field string, delta float64) {
	defer wg.Done()
	conn := pool.Borrow()
	if conn != nil {
		defer pool.Return(conn)
	} else {
		return
	}
	_, err := conn.Do("zadd", uid, strconv.FormatFloat(delta, 'f', 2, 64), field)
	if err != nil {
		log.Printf("zadd error %v", err)
	}

}
