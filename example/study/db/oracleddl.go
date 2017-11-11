package main

import (
	"os"
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-oci8"
	"time"
)

const (
	tableName1 = "zhyhang.tmp_godb1"

	createAutoincr = "CREATE TABLE " + tableName1 + ` (id NUMBER(13,0)
		GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1),
		name VARCHAR2(128),birthday date)`

	msgDsn = `Please specifiy connection parameter in GO_OCI8_CONNECT_STRING environment variable,
or as the first argument! (The format is user/pass@host:port/sid)`

	defaultDsn = "sys/Ipinyou.com2017@127.0.0.1/orcltest?as=sysdba"
)

func getDSN2() string {
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
	fmt.Fprintln(os.Stderr, msgDsn)
	return defaultDsn
}

func main() {
	os.Setenv("NLS_LANG", "")
	db, err := sql.Open("oci8", getDSN2())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	_, err = db.Exec("drop table " + tableName1)
	if err != nil {
		fmt.Println("drop table " + tableName1 + " error:")
		fmt.Println(err)
	}
	_, err = db.Exec(createAutoincr)
	if err != nil {
		fmt.Println("create table " + tableName1 + " error:")
		fmt.Println(err)
	}
	db.Exec("insert into " + tableName1 + "(name,birthday) values ('z',sysdate)")
	row := db.QueryRow("select name,id,birthday from " + tableName1)
	var name string
	var id int
	var birthday time.Time
	err = row.Scan(&name, &id, &birthday)
	if err != nil {
		fmt.Println("query error:")
		fmt.Println(err)
	} else {
		fmt.Printf("query result %d, %s %v", id, name, birthday)
	}
}
