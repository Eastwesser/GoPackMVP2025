package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	".../internal/banks"
	".../internal/logger"
	"golang.org/x/sync/errgroup"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	r := banks.New(bankStore)

	wrappedRouter := logger.AddLoggerMid(log, logger.Middleware(r))

	log.Info("server starting on port :8080")

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 5 * time.Second,
		Handler:           wrappedRouter,
	}

	// GRACEFUL SHUTDOWN LOGIC
	errGroup, errGrpCtx := errgroup.WithContext(context.Background())

	// FIRST GOROUTINE (ERRORS)
	errGroup.Go(func() error {
		if err := server.ListenAndServe(); err != nil {
			log.Error("failed to start server", "error", err)
			return fmt.Errorf("error starting server: %w", err)
		}
		return nil
	})

	// SECOND GOROUTINE (LISTENING + SHUTDOWN)
	errGroup.Go(func() error {
		sigch := make(chan os.Signal, 1)
		signal.Notify(
			sigch,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)
		select {
		case sig := <-sigch: // it's a blocking signal
			log.Info("signal recieved", "signal", sig) // if any recieved - we initiate the shutdown of the server

		case <-errGrpCtx.Done(): // signal 2 shuts down the server
		}

		ctxWithTimeout, cancelFn := context.WithTimeout(errGrpCtx, 5*time.Second)
		defer cancelFn()

		log.Info("initiating graceful shutdown")

		if err := server.Shutdown(ctxWithTimeout); err != nil {
			return fmt.Errorf("error graceful shutdown: %w", err)
		}

		return nil
	})

	// WAIT() for running errors
	if err := errGroup.Wait(); err != nil {
		log.Error("error running", "err", err)
	}
}
