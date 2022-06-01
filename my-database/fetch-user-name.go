package mydb

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	Id      int
	Name    string
	Message string
	Point   int
}

// 全てのユーザ名を取得する
func FetchUserName() ([]string, error) {
	// DBと接続
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("接続失敗", err)
		return []string{}, err
	}
	defer db.Close()

	// nameで検索する
	cmd := "select name from login;"
	rows, err := db.Query(cmd)
	if err != nil {
		log.Fatalln("クエリ取得失敗", err)
		return []string{}, err
	}
	defer rows.Close()

	var names []string
	var tmp string
	for rows.Next() {
		err := rows.Scan(&tmp)
		if err != nil {
			return []string{}, err
		}
		names = append(names)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return []string{}, err
	}
	return names, nil
}
