package main

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
)

type post struct {
	ID    int
	Title string
	Body  string
}

func main() {
	/*dsn := fmt.Sprintf(
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
	fmt.Println(*users2[0].UserID)*/

	db1, mock1, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}
	defer db1.Close()
	mock1.ExpectBegin()
	mock1.ExpectExec("UPDATE products SET").WillReturnResult(sqlmock.NewResult(1, 1))
	mock1.ExpectExec("INSERT INTO product_viewers").WithArgs(4, 10, 5, 11).WillReturnResult(sqlmock.NewResult(2, 2))
	mock1.ExpectCommit()
	fmt.Println(mock1)
	if err = recordStats(db1, 4, 10); err != nil {
		fmt.Println(err)
	}
	fmt.Println("-----")
	if err := mock1.ExpectationsWereMet(); err != nil {
		fmt.Println(err)
		fmt.Println("-----")
	}

	db1, mock1, err = sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}
	rows := sqlmock.NewRows([]string{"id", "title", "body"}).
		AddRow(1, "post 1", "hello").
		AddRow(2, "post 2", "world")

	mock1.ExpectQuery("^SELECT (.+),s FROM posts$").WillReturnRows(rows)
	selects(db1)
	if err := mock1.ExpectationsWereMet(); err != nil {
		fmt.Print("err::::")
		fmt.Println(err)
	}
}

func recordStats(db *sql.DB, userID, productID int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		fmt.Println("defer")
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	result, err := tx.Exec("UPDATE products SET views = views + 200")
	if err != nil {
		return err
	}
	fmt.Println(result)
	i, err := result.LastInsertId()
	fmt.Println(i)
	i, err = result.RowsAffected()
	fmt.Println(i)
	result, err = tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?), (?, ?)", userID, productID, userID+1, productID+1)
	if err != nil {
		return
	}
	fmt.Println(result)
	i, err = result.LastInsertId()
	fmt.Println(i)
	i, err = result.RowsAffected()
	fmt.Println(i)
	return
}

func selects(db *sql.DB) {
	rows, err := db.Query("SELECT id ,title FROM posts")
	if err != nil {
		fmt.Print("err::::")
		fmt.Println(err)
		return
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	var posts []*post
	for rows.Next() {
		p := &post{}
		if err := rows.Scan(&p.ID, &p.Title, &p.Body); err != nil {
			return
		}
		posts = append(posts, p)
		fmt.Println(*p)
	}
	fmt.Println(*rows)
}
