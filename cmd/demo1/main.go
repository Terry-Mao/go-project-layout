package main

import (
	"context"
	kratos "go-project-layout"
	"go-project-layout/server/http"
	"log"
)

func main() {
	svr := http.NewServer()
	app := kratos.New()
	app.Append(kratos.Hook{
		OnStart: func(ctx context.Context) error {
			return svr.Start()
		},
		OnStop: func(ctx context.Context) error {
			return svr.Shutdown(ctx)
		},
	})

	if err := app.Run(); err != nil {
		log.Printf("app failed: %v\n", err)
		return
	}
}
