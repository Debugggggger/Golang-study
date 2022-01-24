package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	id := "root"
	pw := "antline1!"
	ipadd := "127.0.0.1"
	port := "3306"
	dbName := "golang"

	db, err := sql.Open("mysql", id+":"+pw+"@tcp("+ipadd+":"+port+")/"+dbName)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
	
	var name string
    err = db.QueryRow("SELECT name FROM test1 WHERE id = 1").Scan(&name)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(name)
	
}
