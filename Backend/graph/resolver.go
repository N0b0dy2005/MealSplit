package graph

import (
	"context"
	"my-backend/controller"
)

type Resolver struct{}

func GetControllerService(ctx context.Context) *controller.Service {
	return ctx.Value("controllerService").(*controller.Service)
}
