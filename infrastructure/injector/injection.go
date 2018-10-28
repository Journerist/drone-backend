// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

import (
	"github.com/google/go-cloud/wire"
	"github.com/sebarthel/waypoint-handler/infrastructure/provider"
	"github.com/sebarthel/waypoint-handler/infrastructure/service"
)

func InitializeTaskScheduler() (service.TaskScheduler, error) {
	wire.Build(provider.SuperSet)

	return service.TaskScheduler{}, nil
}
