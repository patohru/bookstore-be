package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"bookstore-be/internal/config"
	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

var instance *pgxpool.Pool

func NewPool() *pgxpool.Pool {
	ctx := context.Background()
	if instance != nil {
		return instance
	}

	cfg, _ := env.ParseAs[config.DatabaseConfig]()
	pgxConfig, _ := pgxpool.ParseConfig(cfg.DatabaseURL())
	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	instance, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	if err = instance.Ping(ctx); err != nil {
		fmt.Printf("Unable to ping database\n")
	}

	return instance
}
