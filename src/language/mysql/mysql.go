package mysql

import (
	"database/sql"
	"language/error"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//DB 연결
func DbConnect(id string, pw string, ip string, port string, dbName string) *sql.DB {

	db,err := sql.Open("mysql",id+":"+pw+"@tcp("+ip+":"+port+")/"+dbName)
	if err != nil {
		log.Print(error.ErrDBConnect.Error())
	}

	return db;
}

// DB에서 한줄만 끌어와서 사용하는 경우 
func SelectId(db *sql.DB, id string) string{
	var name string
	err := db.QueryRow("SELECT name FROM language WHERE id="+id).Scan(&name)

	if err !=nil{
		log.Fatal(err)
	}
	return name
}

type Language struct {
	id      int
	title   string
	name    string
	seqence int
	t_type  string
}

// DB에서 여러줄을 끌어와서 사용하는 경우
func SelectSeq(db *sql.DB, seq string) []Language {
	var id int
	var name string
	result := []Language{};
	att := Language{};
	rows, err := db.Query("SELECT id,name FROM language WHERE sequence=?",seq)
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()

	for i:=0 ; rows.Next() ; i++ {
        err := rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
		att.id = id
		att.name = name
		result = append(result,att);
    }
	return result
}