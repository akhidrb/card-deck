package cmd

import (
	"context"
	"github.com/akhidrb/toggl-cards/pkg/db"
	"github.com/akhidrb/toggl-cards/pkg/router"
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	PORT       string `env:"PORT" envDefault:":8080"`
	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBUser     string `env:"DB_USER" envDefault:"toggl"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"toggl"`
	DBName     string `env:"DB_NAME" envDefault:"cards"`
}

func API() *cobra.Command {
	return &cobra.Command{
		Use: "api",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("Starting toggl-cards api")
			runAPI()
		},
	}
}

func runAPI() {
	log.Info("Starting api")
	cfg := config()
	dbConn := db.Connect(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	engine := router.Init(dbConn)
	srv := &http.Server{Addr: cfg.PORT, Handler: engine}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
}

func config() Config {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
