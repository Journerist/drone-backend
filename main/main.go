package main

import (
	"github.com/Journerist/drone-backend/infrastructure/injector-gen"
)

func main() {
	taskScheduler, _ := injector_gen.InitializeTaskScheduler()
	taskScheduler.Start()
}
