// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

import (
	"github.com/Journerist/drone-backend/infrastructure/provider"
	"github.com/Journerist/drone-backend/infrastructure/service"
	"github.com/google/go-cloud/wire"
	"github.com/journerist/drone-backend/interfaces"
)

func InitializeTaskScheduler() (service.TaskScheduler, error) {
	wire.Build(provider.SuperSet)

	return service.TaskScheduler{}, nil
}

func InitializeControllerStarter() (interfaces.ControllerStarter, error) {
	wire.Build(provider.SuperSet)

	return interfaces.ControllerStarter{}, nil
}
