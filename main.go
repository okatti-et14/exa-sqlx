package main

import (
	"fmt"
	"gosqlx/model"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
)

type post struct {
	ID    int
	Title string
	Body  string
}

func main() {

	db1, mock1, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}

	readDB := sqlx.NewDb(db1, "postgres")
	rows := sqlmock.NewRows([]string{"user_id", "password"}).
		AddRow(1, "hello").
		AddRow(2, "world")

	mock1.ExpectQuery(`^select distinct user_id, password,visitcnt AS "mcustshopvisitcnt.visitcnt" 
	from users
	where user_id = \$1 (.+)`).WithArgs(1, "b").WillReturnRows(rows)
	selects(readDB, 1)
	if err := mock1.ExpectationsWereMet(); err != nil {
		fmt.Print("err::::")
		fmt.Println(err)
	}
	rows = sqlmock.NewRows([]string{"user_id"}).AddRow(1)
	mock1.ExpectQuery(`select user_id from users`).WillReturnRows(rows)
	singleSelect(readDB)
	sqlxExe()
}

func selects(db *sqlx.DB, id int) {

	//IN句含めて名前付きQuery使う方法
	bindParams := map[string]interface{}{
		"userid":   id,
		"password": "b",
	}
	basequery := `
	select distinct user_id, password,visitcnt AS "mcustshopvisitcnt.visitcnt"
	from users
	where user_id = :userid and password = :password`
	query, args, err := sqlx.Named(basequery, bindParams)
	if err != nil {
		fmt.Print("err::::")
		fmt.Println(err)
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		fmt.Print("err::::")
		fmt.Println(err)
	}
	users2 := []*model.Users{}
	query = db.Rebind(query)
	err = db.Select(&users2, query, args...)
	if err != nil {
		fmt.Print("err::::")
		fmt.Println(err)
	}
	fmt.Println(query)
	fmt.Println(args)
}

func singleSelect(db *sqlx.DB) {
	user := &model.Users{}
	err := db.Get(user, `select user_id from users`)
	if err != nil {
		fmt.Println("errerrerr", err)
		return
	}
	fmt.Println(user)
	fmt.Println(*user.UserID)
}

func sqlxExe() {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		"test01",
		"test01",
		"localhost",
		"15433",
		"test01",
		"disable")
	db, err := sqlx.Connect("postgres", dsn)
	//defer db.Close()
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	user := []*model.Users{}
	//err = db.Get(user, `select password from users`)
	err = db.Select(&user, `select password from users where password='afa'`)
	if err != nil {
		fmt.Println("errerrerr", err)
		return
	}
	fmt.Println(user)
	if len(user) != 0 {
		fmt.Println(*user[0].Password)
	}
}
