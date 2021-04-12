package task

import (
	"fmt"
)

func (tlist *TaskList) Add(task *Task) {
	tlist.Tasks = append(tlist.Tasks, task)
}

func (tlist *TaskList) Update(id uint64, task *Task) error {
	for _, val := range tlist.Tasks {
		if id == val.ID {
			val.Title = task.Title
			val.Body = task.Body
		}
	}
	return nil
}

func (tlist *TaskList) Delete(id uint64) error {
	for ind, val := range tlist.Tasks {
		if id == val.ID {
			tlist.Tasks = append(tlist.Tasks[:ind], tlist.Tasks[ind+1:]...)
		}
	}
	return nil
}

func (tlist *TaskList) Get(id uint64) *Task {
	var t Task
	for _, val := range tlist.Tasks {
		if id == val.ID {
			t.ID = val.ID
			t.Title = val.Title
			t.Body = val.Body
		}
	}

	return &t
}

func (tlist *TaskList) Print() {
	for _, val := range tlist.Tasks {
		fmt.Println(val)
	}
}
