package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", ":6400")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	result, err := redis.String(conn.Do("info"))
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
