package task

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func (t *Task) Add(task *Task, db *sqlx.DB) {
	// tlist.Tasks = append(tlist.Tasks, task)
	err := db.QueryRowx("INSERT INTO taskList(title, body) values($1, $2) RETURNING id", t.Title, t.Body).Scan(&t.ID)
	if err != nil {
		log.Println(err, " Bla Bla2")
		return
	}
}

func (t *Task) Update(id uint64, db *sqlx.DB) error {
	err := db.QueryRowx("UPDATE taskList SET title = $2, body = $3 WHERE id = $1;", id, t.Title, t.Body)
	if err != nil {
		log.Println(err, " Bla Bla2")
		return nil
	}
	return nil
}

func (t *Task) Delete(id uint64, db *sqlx.DB) error {
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM taskList WHERE id = $1", id)
	tx.Commit()
	return nil
}

func (t *Task) Get(id uint64, db *sqlx.DB) *Task {
	row, err := db.Queryx("SELECT id, title, body FROM taskList WHERE id = $1 LIMIT 1", id)
	if err != nil {
		log.Println(err, " Bla Bla2")
		return nil
	}
	if row.Next() {
		err := row.Scan(&t.ID, &t.Title, &t.Body)
		if err != nil {
			log.Println(err, " Bla Bla2")
			return nil
		}
	}
	return nil
}

func Print(db *sqlx.DB) (tasks []Task) {
	rows, _ := db.Queryx("SELECT * FROM taskList")
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Title, &t.Body)
		if err != nil {
			log.Println(err, " Bla Bla2")
		}
		tasks = append(tasks, t)
	}
	fmt.Println(tasks)
	return
}
