package test

import "github.com/Journerist/drone-backend/core/task"

type TaskRepository struct {
	tasks []task.Task
}

func (tr *TaskRepository) Init(tasks []task.Task) {
	tr.tasks = tasks
}

func (tr TaskRepository) FindAll() []task.Task {
	return tr.tasks
}