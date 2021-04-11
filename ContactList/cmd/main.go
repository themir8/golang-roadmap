package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/mirsaid-mirzohidov/ContactList/contact"
)

func main() {

	DATABASE_URL := "postgres://mirzohidov:coder@localhost:6432/django_examples?sslmode=disable"

	db, err := sqlx.Connect("postgres", DATABASE_URL)
	if err != nil {
		log.Fatalln(err, " Bla Bla1")
	}
	defer db.Close()

	// person.delete(db)
	person := contact.Person{FirstName: "Mirsaid", LastName: "Mirzohidov", Phone: "+998998141352"}

	person.Save(db)

	var p2 contact.Person

	p2.get(db, person.ID)

	fmt.Println(person)
	fmt.Println(p2)

	// fmt.Println(contact.List(db))

}
