package main

import (
	"context"
	"fmt"
	matchValidator "github.com/RaphaelParment/ppio-api/internal/domain/match/validator"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RaphaelParment/ppio-api/internal/application/pp_service"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/config"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/persistence"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/persistence/postgres"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/rest"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	appLogger := log.New(os.Stdout, "", log.LstdFlags)
	err := run(appLogger)
	if err != nil {
		appLogger.Println(err)
		os.Exit(1)
	}
}

func run(logger *log.Logger) error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	dbCfg := persistence.Config{
		User:       cfg.DB.User,
		Password:   cfg.DB.Password,
		Host:       cfg.DB.Host + cfg.DB.Port,
		Name:       cfg.DB.Name,
		DisableTLS: cfg.DB.DisableTLS,
	}

	db, dbTidy, err := persistence.ConnectAndTidy(&dbCfg)
	if err != nil {
		return fmt.Errorf("setup database %w", err)
	}
	defer dbTidy(logger)

	matchStore := postgres.NewMatchStore(logger, db)
	matchService := pp_service.NewMatchService(matchStore, matchValidator.NewMatchValidator())

	server := rest.NewServer(logger, matchService)

	e := echo.New()
	e.GET("/matches/:id", server.HandleGetOneMatch)
	e.GET("/matches", server.HandleGetAllMatches)
	e.POST("/matches", server.HandleAddOneMatch)
	e.PATCH("/matches/:id", server.HandleUpdateOneMatch)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	errs := make(chan error, 1)

	logger.Println("Starting http server")
	go func() {
		err = e.Start(cfg.Http.Port)
		if err != nil {
			switch err {
			case http.ErrServerClosed:
				return
			default:
				errs <- err
				return
			}
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case <-stop:
		err = e.Shutdown(ctx)
		if err != nil {
			logger.Printf("failed to gracefully shutdown http server; %s", err)
			return err
		}

		logger.Printf("shutdown http server gracefully")
		return nil
	case err := <-errs:
		return err
	}
}
