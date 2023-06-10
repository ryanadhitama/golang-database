package golang_database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO user (first_name, last_name) VALUES('ryan', 'adhitama')"

	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data into user")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, first_name, last_name FROM user"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, first_name, last_name string
		err = rows.Scan(&id, &first_name, &last_name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("First Name:", first_name)
		fmt.Println("Last Name:", last_name)
	}
}