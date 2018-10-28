package provider

import (
	"github.com/google/go-cloud/wire"
	"github.com/sebarthel/waypoint-handler/core/task"
	"github.com/sebarthel/waypoint-handler/infrastructure/service"
)

func ProvideTaskRepository() (task.TaskRepository, error) {
	taskRepository := new (task.TaskRepository)
	return *taskRepository, nil
}

func ProvideTaskScheduler(taskRepository task.TaskRepository) (service.TaskScheduler, error) {
	taskScheduler := new (service.TaskScheduler)
	taskScheduler.Init(taskRepository)
	return *taskScheduler, nil
}

var SuperSet = wire.NewSet(ProvideTaskScheduler, ProvideTaskRepository)