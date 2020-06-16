package main

import (
	"fmt"

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
	rows := sqlmock.NewRows([]string{"id", "title", "body"}).
		AddRow(1, "post 1", "hello").
		AddRow(2, "post 2", "world")

	mock1.ExpectQuery(`^SELECT (.+) FROM posts WHERE id = \$1 AND (.+)`).WithArgs(1, 1).WillReturnRows(rows)
	selects(readDB, 1)
	if err := mock1.ExpectationsWereMet(); err != nil {
		fmt.Print("err::::")
		fmt.Println(err)
	}
}

func selects(db *sqlx.DB, id int) {
	rows, err := db.Queryx(`SELECT id ,title FROM posts WHERE id = $1 AND tiltle = $2`, id, id)
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
