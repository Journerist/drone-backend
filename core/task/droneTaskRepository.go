package task

type DroneTaskRepository struct {
}

func (tr DroneTaskRepository) FindAll() []Task {
	tasks := make([]Task, 2)
	tasks[0] = new (CalibrateMotorTask)
	tasks[1] = new (StartMotorTask)
	tasks[2] = new (StartMotorTask)
	return tasks
}