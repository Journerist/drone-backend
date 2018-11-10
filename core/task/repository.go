package task

type Repository interface {
	FindAll() []Task
}
