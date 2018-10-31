package main

import "C"
import (
	"crypto/md5"
	"github.com/tecbot/gorocksdb"
	"log"
	"strconv"
)

func main() {

	options := gorocksdb.NewDefaultOptions()
	options.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(options, "/tmp/gorocks-test")
	if err != nil {
		panic(err)
	}
	log.Println("begin write rocksdb")
	parallel := 1
	que := make(chan int, parallel)
	for i := 0; i < 10000000; i++ {
		que <- 1
		go put(db, que, i)
	}
	for i := 0; i < parallel; i++ {
		que <- 1
	}
	log.Println("end write rocksdb")
	defer db.Close()
}

func put(db *gorocksdb.DB, que chan int, i int) {
	db.Put(gorocksdb.NewDefaultWriteOptions(), strmd5("k"+strconv.Itoa(i)), []byte("v"+strconv.Itoa(i)))
	<-que
}

func strmd5(s string) []byte {
	data := md5.Sum([]byte(s))
	return data[0:]
}
