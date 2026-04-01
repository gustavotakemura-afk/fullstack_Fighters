package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	return sql.Open("mysql", "root:senha@tcp(localhost:3306)/seubanco")
}
