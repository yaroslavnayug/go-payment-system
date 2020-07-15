package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/yaroslavnayug/go-payment-system/internal/persistence"
	"github.com/yaroslavnayug/go-payment-system/internal/usecase"
)

func main() {
	logger := MustLogger()
	postgresConnection := MustPostgres()
	accountRepository := persistence.NewAccountRepository(postgresConnection)
	API := usecase.NewPaymentSystemAPI(logger, accountRepository)

	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/createAccount", API.CreateAccountRequest)

	wg := &sync.WaitGroup{}

	// Start Server
	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	if err := server.Shutdown(context.TODO()); err != nil {
		panic(err)
	}

	// WatchOSSignals
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

	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("app crashed & recovered with: %+v", r)
		}
		postgresConnection.Close()
	}()
}

func MustLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout
	return logger
}

func MustPostgres() *pgxpool.Pool {
	dsn := os.Getenv("POSTGRESQL_URL")
	connection, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(fmt.Errorf("unable to connect to database: %v", err))
	}
	return connection
}
