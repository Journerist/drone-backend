package service

import (
	"testing"

	"github.com/journerist/drone-backend/core/task"
)

func getRemoteController() *RemoteController {
	remoteController := new(RemoteController)
	remoteController.Init()
	return remoteController
}

func TestInitiallyUnpaused(t *testing.T) {
	remoteController := getRemoteController()

	if remoteController.paused != false {
		t.Error("Expected unpaused, got paused")
	}
}

func TestCanBePaused(t *testing.T) {
	remoteController := getRemoteController()
	remoteController.PauseExecution()

	if !(remoteController.paused == true) {
		t.Error("Expected to be paused after Pause()")
	}
}

func TestCommandCanNotBeExecutedWhenNotPaused(t *testing.T) {
	remoteController := getRemoteController()

	err := remoteController.ExecuteTask(new(task.StartTask))

	if !(err != nil) {
		t.Error("Expected to throw an error when executing command when not paused")
	}
}

type mockTask struct {
	executed bool
}

func (mt *mockTask) Execute() {
	mt.executed = true
}

func TestCommandGetsExecuted(t *testing.T) {
	remoteController := getRemoteController()
	mockTask := mockTask{executed: false}

	remoteController.PauseExecution()
	err := remoteController.ExecuteTask(&mockTask)

	if !(err == nil) {
		t.Error("Expected to return no error on command execution")
	}

	if !(mockTask.executed == true) {
		t.Error("Expected to execute given task")
	}
}
