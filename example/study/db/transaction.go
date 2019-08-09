package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strconv"
	"fmt"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(:3306)/test?charset=utf8mb4,utf8&autocommit=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// note: must exec in tx not in db
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	age := rnd.Intn(120) + 1
	result, err := tx.Exec("INSERT INTO godb (name,age) VALUES (?,?)", "Hello guy"+strconv.Itoa(age), age)
	if err != nil {
		log.Println(err)
	}
	lastId, err := result.LastInsertId()
	record := tx.QueryRow("select id,name,age from godb where id=?", lastId)
	var pk int
	var name string
	var age1 int
	record.Scan(&pk, &name, &age1)
	fmt.Printf("just inserted id[%d], name[%s],age[%d]\n", pk, name, age1)
	tx.Rollback()
	fmt.Println("your inserted is rollbacked")
}
