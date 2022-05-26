package mypkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	id   int
	name string
}

func SQLSample() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	defer db.Close()

	// 取得件数を条件[$1]にする
	// ?だとエラーが発生
	//cmd := "select id, order_id from final_sales where id like $1"
	cmd := "select * from mybook;"
	//取得するデータが1件の場合は、QueryRowも利用できる
	rows, err := db.Query(cmd, "T00%")
	if err != nil {
		fmt.Println("クエリ取得失敗")
	}
	defer rows.Close()

}