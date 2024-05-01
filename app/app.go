package app

import (
	"context"

	"go-enum-example/controller"
)

func Run(ctx context.Context) (shutdown func(ctx context.Context)) {
	ctl := controller.NewController()

	go ctl.Start(":0")

	return func(ctx context.Context) {
		_ = ctl.Shutdown(ctx)
	}
}
