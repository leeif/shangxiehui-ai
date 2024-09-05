package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"shangxiehui-ai/config"
	"shangxiehui-ai/internal/controller"
	"shangxiehui-ai/internal/controller/server"
	"shangxiehui-ai/internal/controller/server/route"
	"shangxiehui-ai/internal/infrastructure"
	"time"

	"go.uber.org/fx"
)

// VERSION 版本号(一般编译时注入)
var VERSION = ""

func registerRoute(
	route *route.Route,
	apiserver *server.APIServer,
) error {

	// register api route
	route.RegisterApiV1(apiserver.Engine)

	return nil
}

// @title shangxiehui-ai API
// @Version 0.0.1
// @description Client-side API intended for general users.
//
// @Contact.Name API Support
func main() {

	app := fx.New(
		fx.Provide(
			func() []string {
				return os.Args
			},
			func() string {
				return VERSION
			},
			config.NewConfig,
		),
		infrastructure.Module,
		controller.ServerModule,
		fx.NopLogger,
		fx.Invoke(registerRoute),
		// fx.Invoke(initRecSys),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
