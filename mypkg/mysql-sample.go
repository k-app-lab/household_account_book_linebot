package mypkg

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Sale struct {
	id   int
	name string
}

func SQLSample() {
	db, err := sql.Open("postgres", "postgres://user:pass@host:port/dbname")
	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	defer db.Close()

	// 取得件数を条件[$1]にする
	// ?だとエラーが発生
	//cmd := "select id, order_id from final_sales where id like $1"
	cmd := "select id, order_id from final_sales where id like $1"
	//取得するデータが1件の場合は、QueryRowも利用できる
	rows, _ := db.Query(cmd, "T00%")
	defer rows.Close()

	var sales []Sale // 取得するデータの構造体を用意（複数取得するのでスライス）
	for rows.Next() {
		var tmp Sale // 取得の格納用
		err := rows.Scan(&tmp.id, &tmp.name)
		if err != nil {
			log.Fatalln("取得失敗", err)
		}
		sales = append(sales, tmp)
	}
	for _, sale := range sales {
		fmt.Println(sale.id, sale.name)
	}
}
