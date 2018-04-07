package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Person model
type Person struct {
	ID        int
	Firstname string
	Lastname  string
}

// PersonIndexPage model
type PersonIndexPage struct {
	PageTitle string
	Persons   []Person
}

var persons []Person

func loadName() {
	db, _ := sql.Open("mysql", "root:abcde@/testdb")
	defer db.Close()

	rows, err := db.Query("select * from persons")

	for rows.Next() {
		var id int
		var firstname string
		var lastname string
		err = rows.Scan(&id, &firstname, &lastname)

		if err != nil {
			panic(err.Error)
		}
		persons = append(persons, Person{ID: id, Firstname: firstname, Lastname: lastname})
	}
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	loadName()
	tmpl := template.Must(template.ParseFiles("index.html"))
	data := PersonIndexPage{
		PageTitle: "All persons",
		Persons:   persons,
	}
	tmpl.Execute(w, data)
}

func main() {
	fmt.Println("Web application")
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8000", nil)
}
