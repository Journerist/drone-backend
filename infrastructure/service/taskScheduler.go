package service

import (
	"github.com/Journerist/drone-backend/core/task"
	"log"
	"reflect"
)

type TaskScheduler struct {
	taskRepository task.Repository
}

func (ts *TaskScheduler) Init(taskRepository task.Repository) {
	ts.taskRepository = taskRepository
}

func (ts *TaskScheduler) Start() {
	log.Println("Start Task-Scheduler")
	log.Println("Load first task")

	tasks := ts.taskRepository.FindAll()
	for _, element := range tasks {
		taskName := getType(element)

		println("Start " + taskName)
		element.Execute()
		println("End " + taskName)
	}
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
