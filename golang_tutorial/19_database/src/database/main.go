package main

import (
	"database/sql"

	"fmt"

	_ "github.com/lib/pq"
)

/*
CREATE TABLE accounts (
id SERIAL PRIMARY KEY,
username VARCHAR NOT NULL,
name VARCHAR NOT NULL,
email VARCHAR,
UNIQUE(email)
)
*/

func main() {
	db, err := sql.Open("postgres", "user=emc password=vmware dbname=emccode")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//SELECT
	rows, err := db.Query(`SELECT id, username, name, email FROM
    accounts WHERE username = $1`, "vonthd")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var username string
		var name string
		var email string
		err = rows.Scan(&id, &username, &name, &email)
		if err != nil {
			continue
		}
		fmt.Println("id:", id, " username:", username, " name:", name, " email:", email)
	}

	//INSERT
	var userid int
	err = db.QueryRow(`INSERT INTO accounts (username, name, email)
    VALUES('test', 'MyTest', 'test@test.com') RETURNING id`).Scan(&userid)
	if err != nil {
		panic(err)
	}
	fmt.Println("INSERT:", userid)

	//SELECT
	rows, err = db.Query("SELECT id, username, name, email FROM accounts")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var username string
		var name string
		var email string
		err = rows.Scan(&id, &username, &name, &email)
		if err != nil {
			continue
		}
		fmt.Println("id:", id, " username:", username, " name:", name, " email:", email)
	}

	//DELETE
	userid = 0
	err = db.QueryRow("DELETE FROM accounts WHERE username = $1 RETURNING id", "test").Scan(&userid)
	if err != nil {
		panic(err)
	}
	fmt.Println("DELETE:", userid)
}
