package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/yaroslavnayug/go-payment-system/internal/config"
	"github.com/yaroslavnayug/go-payment-system/internal/handler/v1.0"
	"github.com/yaroslavnayug/go-payment-system/internal/postgres"
	"github.com/yaroslavnayug/go-payment-system/internal/usecase"
)

func main() {
	// Build deps
	cfg := config.GetConfig()
	logger := MustLogger(cfg)
	logger.Infof("starting service with config %+v", cfg)

	//fasthttp.Server

	postgresConnection := MustPostgres(cfg, logger)
	repository := postgres.NewCustomerRepository(postgresConnection)
	accountService := usecase.NewCustomerUseCase(repository)
	customerHandler := v1.NewCustomerHandlerV1(logger, accountService)

	// Assign handlers
	http.HandleFunc("/customer", customerHandler.Create)

	wg := &sync.WaitGroup{}

	// Start server
	server := &http.Server{Addr: ":8080"}
	wg.Add(1)
	go func() {
		defer wg.Done()

		logger.Info("start server on port 8080")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Watch OS signals for shutdown
	wg.Add(1)
	go func() {
		defer wg.Done()

		termChannel := make(chan os.Signal, 1)
		signal.Notify(termChannel, syscall.SIGTERM, syscall.SIGINT)
		<-termChannel

		logger.Info("receive sigterm")
		logger.Info("trying to stop server with grace")
		err := server.Shutdown(context.Background())
		if err != nil {
			logger.Fatalf("unable to stop http server: %s", err)
		}
	}()

	// Close connections
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("app crashed & recovered with: %+v", r)
		}

		logger.Info("close postgres connection")
		postgresConnection.Close()
	}()

	wg.Wait()
}

func MustLogger(config config.Config) *logrus.Logger {
	logger := logrus.New()
	logger.Level = config.LogConfig.Level
	logger.Out = os.Stdout
	return logger
}

func MustPostgres(config config.Config, logger *logrus.Logger) *pgxpool.Pool {
	connection, err := pgxpool.Connect(context.Background(), config.PostgresConfig.HostString)
	if err != nil {
		logger.Fatalf("unable to connect to database: %v", err)
	}
	return connection
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return nil
}
