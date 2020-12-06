package main

import (
	"context"
	"fmt"
	"go-project-layout/server/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	svr := http.NewServer()

	// http server
	g.Go(func() error {
		defer func() {
			<-ctx.Done()
			svr.Shutdown(context.Background())
		}()
		return svr.Start()
	})

	// signal
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-sig:
			// do something
			return nil
		}
	})

	err := g.Wait() // first error return
	fmt.Println(err)
}
