package service

import (
	"log"
	"reflect"
	"time"

	"github.com/Journerist/drone-backend/core/task"
)

type TaskScheduler struct {
	taskRepository task.TaskRepository
}

func (ts TaskScheduler) Init(taskRepository task.TaskRepository) {
	ts.taskRepository = taskRepository
}

func (ts TaskScheduler) Start() {
	log.Println("Start Task-Scheduler")
	log.Println("Load first task")

	tasks := ts.taskRepository.FindAll()
	for _, element := range tasks {
		taskName := getType(element)

		println("Start " + taskName)
		element.Execute()
		println("End " + taskName)
		time.Sleep(time.Second * 1)
	}
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
