package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"fmt"
	"sync"
)

func main() {
	//db, err := sql.Open("mysql", "optimus:ipinyou.com@tcp(192.168.144.55:3306)/test?charset=utf8mb4,utf8&autocommit=true")
	db, err := sql.Open("mysql", "root:@tcp(:3306)/test?charset=utf8mb4,utf8&autocommit=true")
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(24)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			records,_:=db.Query("SELECT name,age FROM godb WHERE id = ?", i)
			//must to trigger close for releasing connection
			for records.Next() {
			}
			fmt.Printf("pool holding %d connections\n",  db.Stats().OpenConnections)
			wg.Done()
		}()
	}
	wg.Wait()

}
