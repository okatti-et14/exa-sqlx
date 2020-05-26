package main

import (
	"fmt"
	"gosqlx/model"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		"test01",
		"test01",
		"localhost",
		"15432",
		"test01",
		"disable")
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	users := []model.Users{}
	db.Select(&users, "select * from users limit 1")
	fmt.Println(*users[0].UserId)
}
