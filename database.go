package golang_database

import (
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang-db")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10) // set minimal connection
	db.SetMaxOpenConns(100) // set maximal connection
	db.SetConnMaxIdleTime(5 * time.Minute) // set how long connection not used will erased
	db.SetConnMaxLifetime(60 * time.Minute) // set how long connection could used

	return db
}