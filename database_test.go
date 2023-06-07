package golang_database

import (
	"testing"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang-db")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}