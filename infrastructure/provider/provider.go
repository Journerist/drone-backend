package provider

import (
	"github.com/Journerist/drone-backend/core/task"
	"github.com/Journerist/drone-backend/infrastructure/service"
	"github.com/google/go-cloud/wire"
)

func ProvideTaskRepository() (task.TaskRepository, error) {
	taskRepository := new(task.TaskRepository)
	return *taskRepository, nil
}

func ProvideTaskScheduler(taskRepository task.TaskRepository) (service.TaskScheduler, error) {
	taskScheduler := new(service.TaskScheduler)
	taskScheduler.Init(taskRepository)
	return *taskScheduler, nil
}

var SuperSet = wire.NewSet(ProvideTaskScheduler, ProvideTaskRepository)
