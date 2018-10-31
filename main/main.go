package main

import (
	"github.com/Journerist/drone-backend/infrastructure/injector-gen"
)

func main() {
	taskScheduler, _ := injector_gen.InitializeTaskScheduler()
	go taskScheduler.Start()

	controllerStarter, _ := injector_gen.InitializeControllerStarter()
	controllerStarter.StartAllController()
}
