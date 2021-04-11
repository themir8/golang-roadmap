package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	DATABASE_URL := "postgres://mirzohidov:coder@localhost:6432/django_examples?sslmode=disable"

	db, err := sqlx.Connect("postgres", DATABASE_URL)
	if err != nil {
		log.Fatalln(err, " Bla Bla1")
	}
	defer db.Close()

	// person.delete(db)
	person := Person{FirstName: "Mirsaid", LastName: "Mirzohidov", Phone: "+998998141352"}

	person.save(db)

	var p2 Person

	p2.get(db, person.ID)

	fmt.Println(person)
	fmt.Println(p2)

	fmt.Println(PersonList(db))

}
