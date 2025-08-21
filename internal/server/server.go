package server

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/go-fuego/fuego"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"

	"bookstore-be/internal/config"
	"bookstore-be/internal/controllers/ping"
	middleware "bookstore-be/internal/server/middlewares"
)

type Server struct {
	db *pgxpool.Pool
}

func NewServer() *fuego.Server {
	cfg, _ := env.ParseAs[config.ServerConfig]()

	s := fuego.NewServer(
		fuego.WithAddr(fmt.Sprintf(":%d", cfg.Port)),
		fuego.WithGlobalMiddlewares(middleware.Cors),
		fuego.WithEngineOptions(
			fuego.WithOpenAPIConfig(fuego.OpenAPIConfig{
				DisableDefaultServer: true,
				DisableMessages:      true,
			}),
		),
	)

	ping.RegisterRoutes(s)

	return s
}
