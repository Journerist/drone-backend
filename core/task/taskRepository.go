package task

type TaskRepository struct {
}

func (tr TaskRepository) FindAll() []Task {
	tasks := make([]Task, 2)
	tasks[0] = new (StartTask)
	tasks[1] = new (EndTask)
	return tasks
}