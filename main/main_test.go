package main

import (
	"github.com/Journerist/drone-backend/core/task"
	"github.com/Journerist/drone-backend/core/task/test"
	"github.com/Journerist/drone-backend/infrastructure/service"
	"testing"
)

func TestAllRepositoryTasksAreExecuted(t *testing.T) {
	scheduler := new(service.TaskScheduler)

	testTask1 := new(test.Task)
	testTask2 := new(test.Task)

	repository := new(test.TaskRepository)
	repository.Init([]task.Task{testTask1, testTask2})

	scheduler.Init(repository)

	scheduler.Start()

	if !testTask1.IsExecuted()  || !testTask2.IsExecuted() {
		t.Error("not all test tasks were automatically executed")
	}
}
