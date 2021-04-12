package main

import (
	"math/rand"
	"time"

	"github.com/mirsaid-mirzohidov/TaskList/task"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	tlist := task.NewTask()

	t1 := task.Task{
		ID:    14334,
		Title: "New task",
		Body:  "Bla Bla bla",
	}

	t2 := task.Task{
		ID:    13434234,
		Title: "Archive task",
		Body:  "Bla Bla bla",
	}

	t3 := task.Task{
		ID:    17853788,
		Title: "Last task",
		Body:  "Bla Bla bla",
	}

	tlist.Add(&t1)
	tlist.Add(&t2)
	tlist.Add(&t3)
	tlist.Print()

}
