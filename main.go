// @title Northwind API
// @version 1.0
// @description RESTful API for Northwind database
// @in header
// @name Authorization
package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "northwind-api/docs"

	"northwind-api/internal/config"
	"northwind-api/internal/logging"
	"northwind-api/internal/routes"
	"northwind-api/internal/server"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	// 1) Load env
	if err := godotenv.Load(); err != nil {
		log.Warn().Err(err).Msg(".env not found; falling back to environment")
	}

	// 2) Config & dependencies
	cfg := config.LoadConfig() // pastikan struct ini punya Env, APIVer, Port, DBPath
	db := config.SetupDB(cfg.DBPath)
	defer func() {
		if err := db.Close(); err != nil {
			log.Error().Err(err).Msg("closing DB")
		}
	}()

	logFile, err := logging.InitLogger("app.log")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init logger")
	}
	defer func() {
		if err := logFile.Close(); err != nil {
			log.Error().Err(err).Msg("closing log file")
		}
	}()

	// 3) Build engine
	if v := os.Getenv("GIN_MODE"); v != "" {
		gin.SetMode(v)
	}
	engine := server.NewEngine()

	// Healthcheck
	engine.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	// Register versioned routes (+ swagger non-prod, + auth gate inside)
	routes.Register(engine, routes.Deps{
		DB:     db,
		Config: cfg,
	})

	// 4) HTTP server
	addr := ":" + cfg.Port
	srv := &http.Server{
		Addr:         addr,
		Handler:      engine,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// 5) Run & graceful shutdown
	log.Info().Msgf("Server running at http://%s (env=%s api=%s)", srv.Addr, cfg.Env(), cfg.APIVer())

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("listen")
		}
	}()

	// Wait for interrupt
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	log.Info().Msg("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error().Err(err).Msg("graceful shutdown failed; forcing close")
		_ = srv.Close()
	}
	log.Info().Msg("Server exited")
}
