package service

import (
	"github.com/Journerist/drone-backend/core/task"
	"github.com/pkg/errors"
)

type RemoteController struct {
	paused bool
}

func (rt *RemoteController) Init() {
	rt.paused = false
}

func (rt *RemoteController) PauseExecution() error {
	if rt.paused {
		return errors.New("Remote controller is already paused")
	}
	rt.paused = true

	return nil
}

func (rt *RemoteController) ContinueExecution() error {
	if !rt.paused {
		return errors.New("Remote controller is not paused")
	}
	rt.paused = false

	return nil
}

func (rt *RemoteController) ExecuteTask(task task.Task) error {
	if !rt.paused {
		return errors.New("Can not execute task because Execution is not paused")
	} else {
		task.Execute()
	}

	return nil
}
