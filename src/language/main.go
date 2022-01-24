package main

import (
	"fmt"
	"language/mysql"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db := mysql.DbConnect("root","antline1!","127.0.0.1","3306","golang")
	
	name := mysql.SelectId(db,"2")
	fmt.Println(name)
	result := mysql.SelectSeq(db,"4")
	fmt.Println(result)
	
	db.Close();
}