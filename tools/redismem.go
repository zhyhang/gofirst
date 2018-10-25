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
	"os/signal"
	"runtime"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	// for kill -10 pid dump stack info
	setupSigusr1Trap()
	// dump stack info by http://localhost:6060/debug/pprof/goroutine?debug=2
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	total := 1000000
	idCount := 10
	parallel := 100
	today := time.Now().Format("20060102")
	rand.Seed(time.Now().Unix())
	//pool6400 := redis.NewRedigoPool(":6400", 100)
	//pool6401 := redis.NewRedigoPool(":6401", 100)
	pool6400 := redis.NewCommonPool(":6400", parallel)
	pool6401 := redis.NewCommonPool(":6401", parallel)
	defer pool6400.Close()
	defer pool6401.Close()
	log.Println("begin write to redis")
	for i := 0; i < total; {
		batchWg := &sync.WaitGroup{}
		for j := 0; j < parallel/idCount && i < total; j++ {
			uuid := uuid()
			for k := 0; k < idCount; k++ {
				idInt := rand.Intn(1000000)
				id := strconv.Itoa(idInt)
				field := ":" + today + ":" + id
				batchWg.Add(2)
				go incr(batchWg, pool6400, uuid+field)
				go hincr(batchWg, pool6401, uuid, field, 1)
			}
			i++
		}
		batchWg.Wait()
		if redis.ConnectCounter > 240 {
			log.Printf("connect to redis times: %d", redis.ConnectCounter)
		}
	}
	log.Println("end write to redis")
}

func uuid() string {
	randInt := rand.Int63()
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, uint64(randInt))
	return fmt.Sprintf("%x", md5.Sum(intBytes))
}

func incr(wg *sync.WaitGroup, pool redis.ConnPool, key string) {
	defer wg.Done()
	conn := pool.Borrow()
	if conn != nil {
		defer pool.Return(conn)
	} else {
		return
	}
	_, err := conn.Do("incr", key)
	if err != nil {
		log.Printf("incr error %v", err)
	}
}

func hincr(wg *sync.WaitGroup, pool redis.ConnPool, uid string, field string, delta int) {
	defer wg.Done()
	conn := pool.Borrow()
	if conn != nil {
		defer pool.Return(conn)
	} else {
		return
	}
	_, err := conn.Do("hincrby", uid, field, strconv.Itoa(delta))
	if err != nil {
		log.Printf("hincr error %v", err)
	}

}

//for debug
func setupSigusr1Trap() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			DumpStacks()
		}
	}()
}

func DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}
