package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Account struct {
	Id       int
	Username string
	Name     string
	Email    string
}

type Accounts struct {
	AccountData []Account
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user=emc password=vmware dbname=emccode")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//SELECT
	rows, err := db.Query("SELECT id, username, name, email FROM accounts")
	if err != nil {
		panic(err)
	}

	accts := make([]Account, 0)

	for rows.Next() {
		var id int
		var username string
		var name string
		var email string
		err = rows.Scan(&id, &username, &name, &email)
		if err != nil {
			continue
		}

		fmt.Println("id:", id, " username:", username, " name:", name,
			" email:", email)

		append(accts, Account{id, username, name, email})
	}

	response, err := json.MarshalIndent(Accounts{accts}, "", "  ")
	if err != nil {
		panic(err) //not expecting error... just a short cut
	}

	fmt.Println("response:", string(response))

	fmt.Fprintf(w, string(response))
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hello %v\n", vars["id"])
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hello %v\n", vars["id"])
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/user", ListUsers).Methods("GET")
	mux.HandleFunc("/user", AddUser).Methods("POST")
	mux.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":9000")
}
