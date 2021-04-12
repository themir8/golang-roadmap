package task

type Task struct {
	ID    uint64
	Title string
	Body  string
}

func NewTask() TaskI {
	return &TaskList{}
}

type TaskList struct {
	Tasks []*Task
}
