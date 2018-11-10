package provider

import (
	"github.com/Journerist/drone-backend/core/task"
	"github.com/Journerist/drone-backend/infrastructure/service"
	"github.com/google/go-cloud/wire"
	"github.com/Journerist/drone-backend/core/command"
	"github.com/kataras/iris"
)

func ProvideDroneTaskRepository() (*task.DroneTaskRepository, error) {
	taskRepository := new(task.DroneTaskRepository)
	return taskRepository, nil
}

func ProvideTaskScheduler(taskRepository *task.DroneTaskRepository) (*service.TaskScheduler, error) {
	taskScheduler := new(service.TaskScheduler)
	taskScheduler.Init(taskRepository)
	return taskScheduler, nil
}

func ProvideIrisApp() (*iris.Application, error) {
	return iris.Default(), nil
}

func ProvideCommandRepository() (*command.CommandRepository, error){
	cmdRepository := new(command.CommandRepository)
	return cmdRepository, nil
}

var SuperSet = wire.NewSet(
	ProvideTaskScheduler,
	ProvideDroneTaskRepository,
	ProvideIrisApp,
	ProvideCommandRepository,
)
