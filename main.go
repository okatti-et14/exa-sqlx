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
		"15433",
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

	users := []*model.Users{}
	db.Select(&users,
		`select 
			users.user_id, 
			users.password, 
			users.insert_date, 
			users.update_date, 
			user_names.user_name as "user_names.user_name"
		from users 
		left join user_names
			on users.user_id = user_names.user_id`)

	//IN句含めて名前付きQuery使う方法
	bindParams := map[string]interface{}{
		"userid":   "bbb",
		"password": []string{"bbb", "ccc"},
	}
	basequery := `
	select * 
	from users
	where user_id = :userid 
		and password in (:password)`
	query, args, err := sqlx.Named(basequery, bindParams)
	query, args, err = sqlx.In(query, args...)
	users2 := []*model.Users{}
	query = db.Rebind(query)
	err = db.Select(&users2, query, args...)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println(*users2[0].UserID)
}
