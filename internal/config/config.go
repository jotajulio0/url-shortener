package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DatabaseURL       string
	HTTPAddr          string
	BaseURL           string
	MaxConns          int32
	MinConns          int32
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
	HealthCheckPeriod time.Duration
	MigrationsDir     string
	LogFormat         string
}

func Load() (Config, error) {
	cfg := Config{
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		HTTPAddr:          getEnv("HTTP_ADDR", ":8080"),
		BaseURL:           getEnv("BASE_URL", "http://localhost:8080"),
		MaxConns:          int32(getEnvInt("DB_MAX_CONNS", 10)),
		MinConns:          int32(getEnvInt("DB_MIN_CONNS", 2)),
		MaxConnLifetime:   getEnvTime("DB_MAX_CONN_LIFETIME", time.Minute*10),
		MaxConnIdleTime:   getEnvTime("DB_MAX_CONN_IDLETIME", time.Minute*20),
		HealthCheckPeriod: getEnvTime("DB_HEALTHCHECK_PERIOD", time.Minute),
		MigrationsDir:     getEnv("MIGRATIONS_DIR", "./migrations"),
		LogFormat:         getEnv("LOG_FORMAT", "development"),
	}

	// if cfg.DatabaseURL == "" {
	// 	return Config{}, fmt.Errorf("The Environment Variable 'DATABASE_URL' is required.")
	// }

	if cfg.MaxConns <= 0 {
		return Config{}, fmt.Errorf("DB_MAX_CONNS must be greater than zero.")
	}

	if cfg.MinConns < 0 || cfg.MinConns > cfg.MaxConns {
		return Config{}, fmt.Errorf("DB_MIN_CONNS must be between 0 and DB_MAX_CONNS.")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	n, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return n
}

func getEnvTime(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	time, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return time
}
