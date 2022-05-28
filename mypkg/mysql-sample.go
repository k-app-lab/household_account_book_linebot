package mypkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	Id      int
	Name    string
	Message string
}

func SQLSample() {
	// DBと接続
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	// defer文はreturn前に呼ばれる（スタック実装）
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

// 入力した文字列でクエリを検索する
func FetchLoginMessage(key string) (string, error) {
	// DBと接続
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("接続失敗", err)
		return "", err
	}

	// defer文はreturn前に呼ばれる（スタック実装）
	defer db.Close()

	// nameで検索する
	cmd := "select * from login where name='" + key + "';"
	// 複数行取得する事がない場合は、db.QueryRow(cmd)でもいいみたい
	rows, err := db.Query(cmd)
	if err != nil {
		log.Fatalln("クエリ取得失敗", err)
		return "", err
	}
	defer rows.Close()

	var user User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Message)
		if err != nil {
			log.Fatal(err)
			return "", err
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return user.Message, nil
}
