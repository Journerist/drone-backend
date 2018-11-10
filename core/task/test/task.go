package test

type Task struct {
	executed bool
}

func (tt *Task) Execute() {
	tt.executed = true
}

func (tt *Task) IsExecuted() bool {
	return tt.executed
}