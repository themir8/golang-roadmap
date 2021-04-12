package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/mirsaid-mirzohidov/TaskList/task"
)

func main() {

	DATABASE_URL := "postgres://mirzohidov:coder@localhost:6432/django_examples?sslmode=disable"

	db, err := sqlx.Connect("postgres", DATABASE_URL)
	if err != nil {
		log.Fatalln(err, " Bla Bla1")
	}
	defer db.Close()

	rand.Seed(time.Now().UnixNano())

	t1 := task.Task{
		Title: "New task",
		Body:  "Bla Bla bla",
	}

	t2 := task.Task{
		Title: "Archive task",
		Body:  "Bla Bla bla",
	}

	t3 := task.Task{
		Title: "Last task",
		Body:  "Bla Bla bla",
	}

	t1.Add(&t1, db)
	t2.Add(&t2, db)
	t3.Add(&t3, db)
}
