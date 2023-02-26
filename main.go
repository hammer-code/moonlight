package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/hammer-code/moonlight/config"
	"github.com/hammer-code/moonlight/pkg/logging"
)

func init() {
	logging.InitLogging()
}
func main() {
	ctx := context.Background()
	cfg, err := config.InitConfig()
	if err != nil {
		logging.Error(ctx, err, "failed to init config")
	}

	logging.Info(ctx, "success load config", logging.Fields{
		"api":               cfg.API,
		"database_postgres": cfg.DBPostgres,
	})

	err = config.NewDatabase(cfg.DBPostgres)
	if err != nil {
		logging.Error(ctx, err, "failed to init config")
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.API.Host, strconv.Itoa(cfg.API.Port)),
		Handler: nil,
	}
	cls := make(chan struct{})

	go grafulShutdonw(ctx, server, cls)

	logging.Info(ctx, "server running", logging.Fields{
		"host": cfg.API.Host,
		"port": cfg.API.Port,
	})
	server.ListenAndServe()
	<-cls
}

func grafulShutdonw(ctx context.Context, server *http.Server, cls chan struct{}) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-done

	if err := server.Shutdown(ctx); err == context.DeadlineExceeded {
		logging.Error(ctx, err, "server cannot shutdown")
	}
	close(cls)
}
