package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/hammer-code/moonlight/app/certificates/controller"
	"github.com/hammer-code/moonlight/app/certificates/repository"
	"github.com/hammer-code/moonlight/app/certificates/usecase"
	"github.com/hammer-code/moonlight/app/route"
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

	defer func() {
		if r := recover(); r != nil {
			panicCause := fmt.Sprintf("Recovered panic, %v", r)
			logging.Info(ctx, panicCause)
		}
	}()

	logging.Info(ctx, "success load config", logging.Fields{
		"api":               cfg.API,
		"database_postgres": cfg.DBPostgres,
	})

	db, err := config.NewDatabase(cfg.DBPostgres)
	if err != nil {
		logging.Error(ctx, err, "failed to init config")
	}

	repo := repository.NewRepository(db)
	usecase := usecase.Newusecase(repo)
	ctrl := controller.NewController(usecase)

	handler := route.NewRoute(ctrl)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.API.Host, strconv.Itoa(cfg.API.Port)),
		Handler: handler,
	}
	cls := make(chan struct{})

	go func() {
		for {
			time.Sleep(1 * time.Second)
			resp, err := http.Get("https://hammercode.org")
			if err != nil {
				logging.Error(ctx, err, "failed to init config")
				return
			}

			if resp.StatusCode >= 500 {
				// auto recover hammercodeweb
				cmd := exec.CommandContext(ctx, "pm2", "restart", "hammercodewebsite")
				if err = cmd.Run(); err != nil {
					return
				}
				// auto recover moonligt
				cmd = exec.CommandContext(ctx, "sudo", "systemctl", "restart", "hmc-cert-go.service")
				if err = cmd.Run(); err != nil {
					return
				}
			}
		}
	}()

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
