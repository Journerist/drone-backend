package main

import (
	"github.com/sebarthel/waypoint-handler/infrastructure/injector-gen"
)

func main() {
	taskScheduler, _ := injector_gen.InitializeTaskScheduler()
	taskScheduler.Start()
}