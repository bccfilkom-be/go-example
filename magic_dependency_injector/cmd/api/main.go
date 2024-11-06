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
	"github.com/kataras/iris/v12/core/router"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			context.Background,
			newPostgresqlConfig,
			newServer,
			postgresql.NewPool,
			postgresql.New,
			usecase.NewBookUsecase,
		),
		fx.Invoke(
			handler.RegisterBookHTTP,
		),
	).Run()
}

func newServer(lc fx.Lifecycle) router.Party {
	r := iris.New()
	r.Use(iris.Compression)
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
	return r.Party("/api")
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
