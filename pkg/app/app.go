package app

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dbut2/shortener/pkg/config"
	"github.com/dbut2/shortener/pkg/model"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	db *sql.DB
}

func New(c config.Config) *App {
	connStr := fmt.Sprintf("%s:%s@(%s)/%s", c.DB.Username, c.DB.Password, c.DB.Host, c.DB.Database)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	a := &App{
		db: db,
	}
	fmt.Println("a: ", a)
	return a
}

func (a *App) Shorten(request model.Shorten) {
	_, err := a.db.Query("INSERT INTO links (code, url) VALUES (?, ?)", request.Code, request.Url)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func (a *App) Lengthen(code string) model.Shorten {
	row := a.db.QueryRow("SELECT code as Code, url as URl FROM links WHERE code = ?", code)
	var s model.Shorten
	err := row.Scan(&s.Code, &s.Url)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return s
}
