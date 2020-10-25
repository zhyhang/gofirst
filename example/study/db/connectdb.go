package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(:3306)/test?charset=utf8mb4,utf8&autocommit=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	// insert
	result, err := db.Exec("INSERT INTO godb (name,age) VALUES (?,?)", "Hello 世界!", 19)
	if err != nil {
		log.Println(err)
	}
	affectRow, _ := result.RowsAffected()
	insertedId, _ := result.LastInsertId()
	fmt.Printf("insert affectted rows: %d\n", affectRow)
	// query
	rows, err := db.Query("SELECT name FROM godb WHERE id = ?", insertedId)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d is %s\n", insertedId, name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}
