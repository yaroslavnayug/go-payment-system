package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/buaazp/fasthttprouter"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"github.com/yaroslavnayug/go-payment-system/internal/config"
	"github.com/yaroslavnayug/go-payment-system/internal/handler/v1.0"
	"github.com/yaroslavnayug/go-payment-system/internal/postgres"
	"github.com/yaroslavnayug/go-payment-system/internal/usecase"
	"go.uber.org/zap"
)

func main() {
	// Build deps
	cfg := config.Read()
	logger := MustLogger()
	defer func() {
		_ = logger.Sync()
	}()
	logger.Info(fmt.Sprintf("starting service with config %+v", cfg))

	postgresConnection := MustPostgres(cfg, logger)
	defer func() {
		if r := recover(); r != nil {
			logger.Error(fmt.Sprintf("app crashed & recovered with: %+v", r))
		}

		logger.Info("close postgres connection")
		postgresConnection.Close()
	}()

	repository := postgres.NewCustomerRepository(postgresConnection)
	customerUseCase := usecase.NewCustomerUseCase(repository)
	customerHandler := v1.NewCustomerHandlerV1(
		logger.With(zap.String("handler", "customerV1")),
		customerUseCase,
		v1.NewJSONResponseWriter(logger),
	)

	// Assign handlers
	router := fasthttprouter.New()
	router.POST("/customer", customerHandler.Create)
	router.GET("/customer/:id", customerHandler.Find)
	router.PUT("/customer/:id", customerHandler.Update)
	router.DELETE("/customer/:id", customerHandler.Delete)

	// Start server
	server := &fasthttp.Server{
		Handler: router.Handler,
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		logger.Info("start server on port 8080")
		if err := server.ListenAndServe(":8080"); err != nil {
			logger.Fatal(fmt.Sprintf("fail to start server: %s", err.Error()))
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
		err := server.Shutdown()
		if err != nil {
			logger.Error(fmt.Sprintf("unable to stop http server: %s", err.Error()))
		}
		logger.Info("server stopped")
	}()

	wg.Wait()
}

func MustLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("unable to create logger: %s", err.Error()))
	}
	return logger
}

func MustPostgres(config config.Config, logger *zap.Logger) *pgxpool.Pool {
	pgxCfg, _ := pgx.ParseConfig(config.PostgresConfig.HostString)
	pgxCfg.Logger = zapadapter.NewLogger(logger)
	pgxCfg.LogLevel = config.PostgresConfig.LogLevel
	pgxCfg.PreferSimpleProtocol = true

	pgxPoolCfg, _ := pgxpool.ParseConfig("")
	pgxPoolCfg.ConnConfig = pgxCfg
	pgxPoolCfg.MaxConns = config.PostgresConfig.MaxConnections
	pgxPoolCfg.MinConns = config.PostgresConfig.MinConnections

	connection, err := pgxpool.ConnectConfig(context.Background(), pgxPoolCfg)
	if err != nil {
		logger.Fatal(fmt.Sprintf("unable to connect to database: %s", err.Error()))
	}
	return connection
}
