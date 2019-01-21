package main

import (
	"crypto/md5"
	"github.com/tecbot/gorocksdb"
	"io/ioutil"
	"log"
	"strconv"
)

/**
####################################################################
# pre-require:                                                     #
#                                                                  #
# refer https://github.com/facebook/rocksdb/blob/master/INSTALL.md #
#     1. install zip libs by the instructs writing up url          #
#     2. cd /home/zhyhang/code/                                    #
#     3. git clone https://github.com/facebook/rocksdb.git         #
#     4. cd rocksdb                                                #
#     5. make static_lib(if librocksdb.a not exists)               #
#     6. run this sh                                               #
####################################################################
*/

// run in GoLand:
//     run-edit configuration...
//     click Environment, add following:
//     CGO_CFLAGS="-I/home/zhyhang/code/rocksdb/include"
//     CGO_LDFLAGS="-L/home/zhyhang/code/rocksdb -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd"

// run in cmd:
//     export CGO_CFLAGS="-I/home/zhyhang/code/rocksdb/include"
//     export CGO_LDFLAGS="-L/home/zhyhang/code/rocksdb -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd"
//     go run rocksdb-cli.go
//
func main() {

	options := gorocksdb.NewDefaultOptions()
	options.SetCreateIfMissing(true)
	tempDir, _ := ioutil.TempDir("", "gorocksdb-test")
	db, err := gorocksdb.OpenDb(options, tempDir)
	if err != nil {
		panic(err)
	}
	log.Println("begin write rocksdb")
	parallel := 1
	que := make(chan int, parallel)
	for i := 0; i < 10000; i++ {
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
