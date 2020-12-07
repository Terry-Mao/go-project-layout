package main

import (
	"context"
	kratos "go-project-layout"
	"go-project-layout/server/http"
	"log"
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

	// handle signal

	if err := app.Run(); err != nil {
		log.Printf("app failed: %v\n", err)
		return
	}
}
