// Package main is the entry point for the mocks server.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
	"github.com/mockzilla/connexions/v2/pkg/api"
	"github.com/mockzilla/connexions/v2/pkg/loader"
)

func main() {
	appDir := getAppDir()
	_ = godotenv.Load(fmt.Sprintf("%s/.env", appDir), fmt.Sprintf("%s/.env.dist", appDir))

	initLogger()

	router := initRouter()

	addr := fmt.Sprintf(":%s", getEnv("PORT", "2200"))
	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Starting Mock Server on %s", addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}

func getAppDir() string {
	if appDir := os.Getenv("APP_DIR"); appDir != "" {
		return appDir
	}
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(filepath.Dir(b)))
}

func initLogger() {
	var handler slog.Handler
	if os.Getenv("LOG_FORMAT") == "text" {
		handler = tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelInfo,
			TimeFormat: time.Kitchen,
		})
	} else {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	}
	slog.SetDefault(slog.New(handler))
}

func initRouter() *api.Router {
	router := api.NewRouter()
	_ = api.CreateHealthRoutes(router)
	_ = api.CreateHomeRoutes(router)
	_ = api.CreateServiceRoutes(router)
	_ = api.CreateHistoryRoutes(router)
	loader.LoadAll(router)

	services := loader.DefaultRegistry.List()
	if len(services) == 0 {
		log.Println("WARNING: No services discovered!")
	} else {
		log.Printf("Discovered %d service(s): %v", len(services), services)
	}

	return router
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
