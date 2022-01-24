package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"study/mysql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("test")
	db := mysql.DbConnect("antline","antline1!","127.0.0.1","3306","golang")

	fmt.Println("db.Begin")
	tx, err := db.Begin()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer tx.Rollback()

	reader := bufio.NewReader(os.Stdin)
	start := time.Now()

	// prepared 방식
	fmt.Print("tx.Prepare: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	stmt, err := tx.Prepare("SELECT id FROM test1 WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// stmt
	fmt.Print("stmt.Exec: ")
	text, _ = reader.ReadString('\n')
	fmt.Println(text)

	_, err = stmt.Exec(2)
	if err != nil {
		panic(err.Error())
	}

	// commit
	fmt.Print("tx.Commit: ")
	text, _ = reader.ReadString('\n')
	fmt.Println(text)

	err = tx.Commit()
	if err != nil {
		panic(err.Error())
	}

	end := time.Now()
	fmt.Println(end.Sub(start))

}