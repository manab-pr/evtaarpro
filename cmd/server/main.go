package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/manab-pr/evtaarpro/internal/config"
	"github.com/manab-pr/evtaarpro/internal/datastore"
	"github.com/manab-pr/evtaarpro/internal/httpx"
)

// @title EvtaarPro API
// @version 1.0
// @description AI-Powered Collaboration & Payroll Platform API
// @termsOfService http://evtaarpro.com/terms/

// @contact.name API Support
// @contact.url http://evtaarpro.com/support
// @contact.email support@evtaarpro.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token

func main() {
	// Load .env file in development
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load configuration
	appCfg, pgCfg, redisCfg, err := config.Load(
		"config/app.yaml",
		"config/postgres.yaml",
		"config/redis.yaml",
	)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set Gin mode
	if appCfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize PostgreSQL
	pgStore, err := datastore.NewPostgresStore(pgCfg)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer pgStore.Close()
	log.Println("✓ Connected to PostgreSQL")

	// Initialize Redis
	redisStore, err := datastore.NewRedisStore(redisCfg)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisStore.Close()
	log.Println("✓ Connected to Redis")

	// Initialize router
	router := httpx.NewRouter(appCfg, pgStore, redisStore)

	// Create HTTP server
	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", appCfg.App.Host, appCfg.App.Port),
		Handler:        router,
		ReadTimeout:    appCfg.Server.ReadTimeout,
		WriteTimeout:   appCfg.Server.WriteTimeout,
		IdleTimeout:    appCfg.Server.IdleTimeout,
		MaxHeaderBytes: appCfg.Server.MaxHeaderBytes,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("🚀 Server starting on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("✓ Server exited gracefully")
}
