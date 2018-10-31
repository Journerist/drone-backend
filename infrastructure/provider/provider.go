package provider

import (
	"github.com/Journerist/drone-backend/core/task"
	"github.com/Journerist/drone-backend/infrastructure/service"
	"github.com/google/go-cloud/wire"
	"github.com/journerist/drone-backend/core/command"
	"github.com/journerist/drone-backend/interfaces"
	"github.com/kataras/iris"
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

func ProvideIrisApp() (*iris.Application, error) {
	return iris.Default(), nil
}

func ProvideCommandRepository() (*command.CommandRepository, error){
	cmdRepository := new(command.CommandRepository)
	return cmdRepository, nil
}

func ProvideCommandController(app *iris.Application, cmdRepository *command.CommandRepository) (*interfaces.CommandController, error){
	cmdController := new(interfaces.CommandController)
	cmdController.Init(app, cmdRepository)

	return cmdController, nil
}

func ProvideControllerStarter(cmdController *interfaces.CommandController) (interfaces.ControllerStarter, error) {
	controllerStarter := new (interfaces.ControllerStarter)
	controllerStarter.Init(cmdController)

	return *controllerStarter, nil
}

var SuperSet = wire.NewSet(
	ProvideTaskScheduler,
	ProvideTaskRepository,
	ProvideIrisApp,
	ProvideCommandRepository,
	ProvideCommandController,
	ProvideControllerStarter,
	)
