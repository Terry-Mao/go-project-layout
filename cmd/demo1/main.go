package main

import (
	"context"
	kratos "go-project-layout"
	"go-project-layout/server/http"
	"log"
	"os"
	"syscall"
)

func main() {
	svr := http.NewServer()
	app := kratos.New(kratos.Signal(
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	), kratos.SignalFn(
		func(a *kratos.App, sig os.Signal) {
			switch sig {
			case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
				a.Stop()
			default:
			}
		},
	))
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
