package main

import (
	"context"
	"fmt"
	"log"
	"os"

	handler "github.com/bccfilkom-be/go-example/magic_dependency_injector/book/http"
	"github.com/bccfilkom-be/go-example/magic_dependency_injector/book/usecase"
	"github.com/bccfilkom-be/go-example/magic_dependency_injector/db/postgresql"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kataras/iris/v12"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			context.Background,
			newPostgresqlConfig,
			newRouter,
			postgresql.NewPool,
			postgresql.New,
			usecase.NewBookUsecase,
		),
		fx.Invoke(
			handler.RegisterBookHTTP,
			newServer,
		),
	).Run()
}

func newRouter() *iris.Application {
	r := iris.New()
	r.Use(iris.Compression)
	return r
}

// FIX: decouple server from handler
func newServer(lc fx.Lifecycle, r *iris.Application) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			port := 8080
			fmt.Println("starting server on port:", port)
			go r.Listen(fmt.Sprintf(":%d", port))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return r.Shutdown(ctx)
		},
	})
}

func newPostgresqlConfig() *pgxpool.Config {
	uri := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	cfg, err := pgxpool.ParseConfig(uri)
	if err != nil {
		log.Fatalf("Unable to parse database config: %v\n", err)
	}
	return cfg
}
