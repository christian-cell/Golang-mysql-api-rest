package commons

import (
	"database/sql"
	_ "time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() (*sql.DB, error) {
	return sql.Open("mysql", "tester:secret@tcp(db:3306)/test")
}
