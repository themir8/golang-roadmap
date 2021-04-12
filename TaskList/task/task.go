package task

type Task struct {
	ID    uint64
	Title string
	Body  string
}

func NewTask() TaskInter {
	return &TaskList{}
}

type TaskList struct {
	Tasks []*Task
}
