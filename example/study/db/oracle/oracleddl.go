package main

import (
	"os"
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-oci8"
	"time"
	"github.com/zhyhang/gofirst/example/study/db/dbconst"
)

const (

	tableNameTmp = "zhyhang.tmp_godb1"

	createAutoincr = "CREATE TABLE " + tableNameTmp + ` (id NUMBER(13,0)
		GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1),
		name VARCHAR2(128),birthday date)`
)

func getDsnDdl() string {
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
	fmt.Fprintln(os.Stderr, dbconst.MsgDsn)
	return dbconst.DefaultDsn
}

func main() {
	os.Setenv("NLS_LANG", "Simplified Chinese_china.AL32UTF8")
	db, err := sql.Open("oci8", getDsnDdl())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	_, err = db.Exec("drop table " + tableNameTmp)
	if err != nil {
		fmt.Println("drop table " + tableNameTmp + " error:")
		fmt.Println(err)
	}
	_, err = db.Exec(createAutoincr)
	if err != nil {
		fmt.Println("create table " + tableNameTmp + " error:")
		fmt.Println(err)
	}
	db.Exec("insert into " + tableNameTmp + "(name,birthday) values ('z',sysdate)")
	row := db.QueryRow("select name,id,birthday from " + tableNameTmp)
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
