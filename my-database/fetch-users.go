package mydb

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// 全員のユーザデータを取得
func FetchUsers() ([]User, error) {
	// DBと接続
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("接続失敗", err)
		return make([]User, 0), err
	}
	defer db.Close()

	// nameで検索する
	cmd := "select * from login;"
	rows, err := db.Query(cmd)
	if err != nil {
		log.Fatalln("クエリ取得失敗", err)
		return make([]User, 0), err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var tmp User
		err := rows.Scan(&tmp.Id, &tmp.Name, &tmp.Message, &tmp.Point)
		if err != nil {
			log.Fatal(err)
			return []User{}, err
		}
		users = append(users, tmp)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return make([]User, 0), err
	}
	return users, nil
}
