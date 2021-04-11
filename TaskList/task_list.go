package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Task struct {
	ID   int
	Text string
}

func (p Person) save(db *sqlx.DB) id {
	err := db.QueryRowx("INSERT INTO contactList(first_name, last_name, phone) values($1, $2, $3) RETURNING id", p.FirstName, p.LastName, p.Phone).Scan(&p.ID)
	if err != nil {
		log.Println(err, " Bla Bla2")
		return
	}
	fmt.Println("person saved")
}

func (p *Person) get(db *sqlx.DB, id int) {
	row, err := db.Queryx("SELECT id, first_name, last_name, phone FROM contactList WHERE id = $1 LIMIT 1", id)
	if err != nil {
		log.Println(err, " Bla Bla2")
		return
	}
	if row.Next() {
		err := row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Phone)
		if err != nil {
			log.Println(err, " Bla Bla2")
			return
		}
	}
	fmt.Println(p)
	fmt.Println("person found")
}

func (p *Person) delete(db *sqlx.DB) {
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM contactList WHERE id = $1", p.ID)
	tx.Commit()
}

func PersonList(db *sqlx.DB) (persons []Person) {
	rows, _ := db.Queryx("SELECT * FROM contactList")
	for rows.Next() {
		var person Person
		err := rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Phone)
		if err != nil {
			log.Println(err, " Bla Bla2")
		}
		persons = append(persons, person)
	}
	fmt.Println(persons)
	fmt.Println("person list")
	return
}

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
