package mydb

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

// 指定したユーザのポイントを取得する
func UpdatePoint(name string) (int, error) {
	// DBと接続
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("接続失敗", err)
		return 0, err
	}

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
	// ポイントを更新するクエリ文
	updateCmd := "update login set point ='"
	updateCmd += strconv.Itoa(addPoint)
	updateCmd += "' where name='" + name + "';"

	_, err = db.Exec(updateCmd)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return addPoint, nil
}
