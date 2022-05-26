package mypkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	Id   int
	Name string
}

func SQLSample() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	defer db.Close()

	cmd := "select * from mybook;"
	rows, err := db.Query(cmd)
	if err != nil {
		log.Fatalln("クエリ取得失敗", err)
	}
	defer rows.Close()

	var user User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", user.Id, user.Name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
