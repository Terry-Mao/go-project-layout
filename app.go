package kratos

import (
	"context"
	"os"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

// Hook is a pair of start and stop callbacks.
type Hook struct {
	OnStart func(context.Context) error
	OnStop  func(context.Context) error
}

// Option is a life cycle option.
type Option func(o *options)

// options is a life cycle options.
type options struct {
	startTimeout time.Duration
	stopTimeout  time.Duration
	signals      []os.Signal
}

// App is manage the application component life cycle.
type App struct {
	opts  options
	hooks []Hook

	eg     *errgroup.Group
	cancel func()
}

// New new a application manage.
func New(opts ...Option) *App {
	options := options{
		startTimeout: time.Second * 30,
		stopTimeout:  time.Second * 30,
		signals: []os.Signal{
			syscall.SIGTERM,
			syscall.SIGQUIT,
			syscall.SIGINT,
		},
	}
	for _, o := range opts {
		o(&options)
	}
	return &App{opts: options}
}

// Append register callbacks that are executed on application start and stop.
func (a *App) Append(hook Hook) {
	a.hooks = append(a.hooks, hook)
}

// Run executes all OnStart hooks registered with the application's Lifecycle.
func (a *App) Run() error {
	var ctx context.Context
	ctx, a.cancel = context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)
	for _, hook := range a.hooks {
		hook := hook
		if hook.OnStop != nil {
			a.eg.Go(func() error {
				<-ctx.Done() // wait for stop signal
				stopCtx, cancel := context.WithTimeout(context.Background(), a.opts.startTimeout)
				defer cancel()
				return hook.OnStop(stopCtx)
			})
		}
		if hook.OnStart != nil {
			g.Go(func() error {
				startCtx, cancel := context.WithTimeout(context.Background(), a.opts.startTimeout)
				defer cancel()
				return hook.OnStart(startCtx)
			})
		}
	}
	return g.Wait()
}

// Stop gracefully stops the application.
func (a *App) Stop() {
	a.cancel()
}
