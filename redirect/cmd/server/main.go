package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	_ "github.com/lib/pq" // PostgreSQL driver

	"qr.mandacode.com/redirect/ent"
	"qr.mandacode.com/redirect/internal/router"
)

func main() {
	loadEnv()

	logger := setupLogger()
	defer logger.Sync()

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_SSLMODE"),
	)

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		logger.Fatal("failed to connect to postgres", zap.Error(err))
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Fatal("failed to create schema", zap.Error(err))
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(ginLoggerMiddleware(logger))

	router.RegisterRoutes(r, client)

	if err := r.Run(":8080"); err != nil {
		logger.Fatal("failed to run server", zap.Error(err))
	}
}

// ginLoggerMiddleware logs requests excluding /health
func ginLoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		if c.Request.URL.Path == "/health" {
			return
		}

		if c.Request.URL.Path == "/favicon.ico" {
			return
		}

		logger.Info("request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("client_ip", c.ClientIP()),
			zap.Duration("latency", time.Since(start)),
		)
	}
}

func setupLogger() *zap.Logger {
	env := os.Getenv("APP_ENV")
	if env == "prod" {
		logger, _ := zap.NewProduction()
		return logger
	}
	logger, _ := zap.NewDevelopment()
	return logger
}

func loadEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	if env == "dev" || env == "test" {
		envFile := ".env." + env
		if err := godotenv.Load(envFile); err != nil {
			fmt.Printf("No %s file found. Skipping .env load\n", envFile)
		} else {
			fmt.Printf("Loaded environment variables from %s\n", envFile)
		}
	} else {
		fmt.Println("Skipping .env loading; expecting environment variables from orchestrator.")
	}
}
