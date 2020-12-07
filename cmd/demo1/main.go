package main

import (
	"context"
	kratos "go-project-layout"
	"go-project-layout/server/http"
	"time"
)

func main() {
	app := kratos.New()
	svr := http.NewServer()
	app.Append(kratos.Hook{
		OnStart: func(ctx context.Context) error {
			return svr.Start()
		},
		OnStop: func(ctx context.Context) error {
			return svr.Shutdown(ctx)
		},
	})

	app.Start()

	time.Sleep(5 * time.Second)
	app.Stop()
	time.Sleep(5 * time.Second)
}
