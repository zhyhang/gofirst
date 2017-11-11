package main

import (
	"fmt"
	"os"
	"database/sql"
	"github.com/mattn/go-oci8"
)

const (
	tableName2 = "zhyhang.example_insertid"

	sqlCreatetable = "CREATE TABLE " + tableName2 + ` (id NUMBER(13,0)
		GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1)
		primary key, data varchar2(256))`

	msgDsn1 = `Please specifiy connection parameter in GO_OCI8_CONNECT_STRING environment variable,
or as the first argument! (The format is user/pass@host:port/sid)`

	defaultDsn1 = "sys/Ipinyou.com2017@127.0.0.1/orcltest?as=sysdba"
)

func getDSN1() string {
	var dsn string
	if len(os.Args) > 1 {
		dsn = os.Args[1]
		if dsn != "" {
			return dsn
		}
	}
	dsn = os.Getenv("GO_OCI8_CONNECT_STRING")
	if dsn != "" {
		return dsn
	}
	fmt.Fprintln(os.Stderr, msgDsn1)
	return defaultDsn1
}

func main() {
	os.Setenv("NLS_LANG", "")
	db, err := sql.Open("oci8", getDSN1())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	db.Exec("drop table " + tableName2)
	_, err = db.Exec(sqlCreatetable)
	if err != nil {
		fmt.Println("create table " + tableName2 + " error:")
		fmt.Println(err)
	}
	res, err := db.Exec("insert into " + tableName2 + "(data) values ('世界你好')")
	if err != nil {
		fmt.Println("insert row error:")
		fmt.Println(err)
		return
	}
	insertid, err := res.LastInsertId()
	if err != nil {
		fmt.Println("Last insert Id error:")
		fmt.Println(err)
		return
	}
	rowID := oci8.GetLastInsertId(insertid)
	var id int
	err = db.QueryRow("select id from "+tableName2+" where rowid = :1", rowID).Scan(&id)
	if err != nil {
		fmt.Println("oci8 last insert Id error:")
		fmt.Println(err)
		return
	}
	fmt.Printf("last inserted id %d, rowId %s, id from oci8 %d\n", insertid, rowID, id)

}
