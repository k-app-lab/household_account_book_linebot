package mypkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func mysqlSample() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("mysql", os.Getenv("DB_ROLE")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := 2
	var name string
	err = db.QueryRow("SELECT name FROM users WHERE opening_id = ?", id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}
