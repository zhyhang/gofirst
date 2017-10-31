package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "optimus:ipinyou.com@tcp(192.168.144.55:3306)/test?charset=utf8mb4,utf8")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	// insert
	result, err := db.Exec("INSERT INTO godb (name) VALUES (?)", "北京你好")
	if err != nil {
		log.Println(err)
	}
	affectRow, _ := result.RowsAffected()
	fmt.Printf("insert affectted rows: %d\n", affectRow)
	// query
	rows, err := db.Query("SELECT name FROM godb WHERE id = ?", 2)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s is %d\n", name, 2)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}
