package main

import (
	"github.com/journerist/drone-backend/infrastructure/injector-gen"
)

func main() {
	taskScheduler, _ := injector_gen.InitializeTaskScheduler()
	taskScheduler.Start()
}
