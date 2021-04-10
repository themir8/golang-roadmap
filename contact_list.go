package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var persons = []Person{}
var person = Person{}

type Person struct {
	ID         uint16 `db:id`
	first_name string
	last_name  string
	phone      string
}

func (p *Person) add(db *sqlx.DB) {
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO contactList(first_name, last_name, phone) values($1, $2, $3)", p.first_name, p.last_name, p.phone)
	tx.Commit()
}

func (p *Person) delete(db *sqlx.DB, id int) {
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM contactList WHERE id = $1", id)
	tx.Commit()
}

func (p *Person) get(db *sqlx.DB) {
	row, _ := db.Queryx("SELECT id, first_name, last_name, phone FROM contactList WHERE id = $1 LIMIT 1", p.ID)
	if row.Next() {
		err := row.Scan(&p.ID, &p.first_name, &p.last_name, &p.phone)
		if err != nil {
			log.Println(err, " Bla Bla2")
		}
	}
	fmt.Println(p)
}

func (p *Person) list(db *sqlx.DB) {
	rows, _ := db.Queryx("SELECT * FROM contactList")
	for rows.Next() {
		var person Person
		err := rows.Scan(&person.ID, &person.first_name, &person.last_name, &person.phone)
		if err != nil {
			log.Println(err, " Bla Bla2")
		}
		persons = append(persons, person)
	}
	fmt.Println(persons)
}

func main() {

	DATABASE_URL := "postgres://mirzohidov:coder@localhost:6432/django_examples?sslmode=disable"

	db, err := sqlx.Connect("postgres", DATABASE_URL)
	if err != nil {
		log.Fatalln(err, " Bla Bla1")
	}
	defer db.Close()

	// person.delete(db)
	person := Person{first_name: "Mirsaid", last_name: "Mirzohidov", phone: "+998998141352"}

	var str string
	fmt.Print("Command: ")
	fmt.Scan(&str)
	if str == "add" {
		person.add(db)
	} else if str == "posts" {
		person.list(db)
	} else if str == "get" {
		// var i int
		// fmt.Print("Person id: ")
		// fmt.Scan(&i)
		// person.list(db)
		p1 := Person{ID: 3}

		p1.get(db)
	} else if str == "delete" {
		var i int
		fmt.Print("Person id: ")
		fmt.Scan(&i)
		person.list(db)
		person.delete(db, i)
		person.list(db)
	}

}

// func person_list(db *sqlx.DB) {
// 	rows, _ := db.Queryx("SELECT * FROM contactList")
// 	for rows.Next() {
// 		var person Person
// 		err := rows.Scan(&person.ID, &person.first_name, &person.last_name, &person.phone)
// 		if err != nil {
// 			log.Println(err, " Bla Bla2")
// 		}
// 		persons = append(persons, person)
// 	}
// 	fmt.Println(persons)
// }

// func add_person(db *sqlx.DB) {
// 	var str string
// 	fmt.Print("First name: ")
// 	input, _ := fmt.Scan(&str)
// 	if input == 1 {
// 		person.first_name = str
// 		fmt.Println("Ok")
// 		last_name(db)
// 	}
// }

// func last_name(db *sqlx.DB) {
// 	var str string
// 	fmt.Print("Last name: ")
// 	input, _ := fmt.Scan(&str)
// 	if input == 1 {
// 		person.last_name = str
// 		fmt.Println("Ok")
// 		phone(db)
// 	}
// }

// func phone(db *sqlx.DB) {
// 	var str string
// 	fmt.Print("Phone: ")
// 	input, _ := fmt.Scan(&str)
// 	if input == 1 {
// 		person.phone = str
// 		fmt.Println("Ok")
// 		// Save person
// 		save(db)
// 	}
// }
// func save(db *sqlx.DB) {
// 	tx := db.MustBegin()
// 	tx.MustExec("INSERT INTO contactList(first_name, last_name, phone) values($1, $2, $3)", person.first_name, person.last_name, person.phone)
// 	tx.Commit()
// }
