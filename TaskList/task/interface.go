package task

type TaskI interface {
	Add(tlist *Task)
	Update(id uint64, tlist *Task) error
	Delete(id uint64) error
	Get(id uint64) *Task
	Print()
}
