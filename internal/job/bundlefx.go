package job

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"go.uber.org/fx"
)

func registerHooks(
	lifecycle fx.Lifecycle,
) {
	idleConnsClosed := make(chan struct{})
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					sigint := make(chan os.Signal, 1)
					signal.Notify(sigint, os.Interrupt) // Catch OS signals.
					<-sigint
					// slog.Info("Gracefully shutting down...")

					// Received an interrupt signal, shutdown.
					// if err := app.Shutdown(); err != nil {
					// 	// Error from closing listeners, or context timeout:
					// 	slog.Errorf("Server is not shutting down! Reason: %v", err)
					// }
					fmt.Println(333)

					close(idleConnsClosed)
				}()

				go func() {
					fmt.Println(222)
					// slog.Infof("Server running, port=%s", cfg.ApplicationConfig.Address)
					// if err := app.Listen(cfg.ApplicationConfig.Address); err != nil {
					// 	slog.Errorf("Server is not running! Reason: %v", err)
					// }
				}()

				return nil
			},
			OnStop: func(context.Context) error {
				// db.Close()
				<-idleConnsClosed
				// return slog.Sync()
				return nil
			},
		},
	)
}
