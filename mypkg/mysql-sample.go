package mypkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type User struct {
	Id      int
	Name    string
	Message string
	Point   int
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
		err := rows.Scan(&user.Id, &user.Name, &user.Message, &user.Point)
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

// 指定したユーザのポイントを取得する
func UpdatePoint(name string) (int, error) {
	// DBと接続
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("接続失敗", err)
		return 0, err
	}

	// defer文はreturn前に呼ばれる（スタック実装）
	defer db.Close()

	// nameで検索する
	cmd := "select point from login where name='" + name + "';"
	row := db.QueryRow(cmd)

	var point int
	err = row.Scan(&point)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	addPoint := point + 1
	updateCmd := "update login set point ='"
	updateCmd += strconv.Itoa(addPoint)
	updateCmd += "' where name='" + name + "';"

	fmt.Println(updateCmd)
	_, err = db.Exec(updateCmd)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return addPoint, nil
}

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
			return make([]User, 0), err
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
