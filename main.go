package main

import (
	"context"
	"os"
	"os/signal"
	"smolathon/config"
	"smolathon/internal/app"
	"smolathon/pkg/sctx"
	"smolathon/pkg/slogger"
	"syscall"
	"time"
)

// @securityDefinitions.apikey ExternalId
// @in header
// @name ExternalId
func main() {
	settings := config.Read()

	var opts []slogger.Option
	if settings.Env != config.EnvProd {
		opts = append(opts, slogger.Development())
	}
	l := slogger.NewLogger("smolathon", opts...)

	mainCtx, mainCtxStop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer mainCtxStop()

	app, err := app.NewApp(l, settings, sctx.DefaultContextProvider(mainCtx))
	if err != nil {
		l.Error(err.Error())
		mainCtxStop()
	} else {
		app.Run()
		l.Debug("Successful started")
	}

	select {
	case <-mainCtx.Done():
	case <-app.Notify():
	}
	l.Debug("Shutting down")

	shutdownCtx, shutdownCtxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCtxCancel()

	app.Stop(shutdownCtx)

	<-shutdownCtx.Done()
	l.Debug("Successful stopped")
}
