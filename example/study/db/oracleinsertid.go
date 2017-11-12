package main

import (
	"fmt"
	"os"
	"database/sql"
	"github.com/mattn/go-oci8"
	"github.com/zhyhang/gofirst/example/study/db/dbconst"
)

const (

	tableNameInsertid = "zhyhang.example_insertid"

	sqlTableInsertid = "CREATE TABLE " + tableNameInsertid + ` (id NUMBER(13,0)
		GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1)
		primary key, data varchar2(256))`
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
	fmt.Fprintln(os.Stderr, dbconst.MsgDsn)
	return dbconst.DefaultDsn
}

func main() {
	os.Setenv("NLS_LANG", "Simplified Chinese_china.AL32UTF8")
	db, err := sql.Open("oci8", getDSN1())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	//db.Exec("drop table " + tableName2)
	_, err = db.Exec(sqlTableInsertid)
	if err != nil {
		fmt.Println("create table " + tableNameInsertid + " error:")
		fmt.Println(err)
	}
	res, err := db.Exec("insert into " + tableNameInsertid + "(data) values ('世界你好，我很好！')")
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
	err = db.QueryRow("select id from "+tableNameInsertid+" where rowid = :1", rowID).Scan(&id)
	if err != nil {
		fmt.Println("oci8 last insert Id error:")
		fmt.Println(err)
		return
	}
	fmt.Printf("last inserted id %d, rowId %s, id from oci8 %d\n", insertid, rowID, id)

}
