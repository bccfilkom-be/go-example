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

	"github.com/bccfilkom-be/go-example/opentelemetry/common"
	"github.com/bccfilkom-be/go-example/opentelemetry/db/postgresql"
	"github.com/bccfilkom-be/go-example/opentelemetry/pet/rest"
	"github.com/bccfilkom-be/go-example/opentelemetry/pet/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}
}

func mux() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	uri := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	cfg, err := pgxpool.ParseConfig(uri)
	if err != nil {
		log.Fatalln(err)
	}
	poll, err := common.NewPostgreSQLPool(cfg)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer poll.Close()

	petAPI := chi.NewRouter()
	postgresql := postgresql.New(poll)
	petUsecase := usecase.NewPetUsecase(postgresql)
	petHandler := rest.NewPetHandler(petUsecase)
	petHandler.Register(petAPI)
	r.Mount("/api/v1/pets", petAPI)

	return r
}

func main() {
	server := &http.Server{Addr: ":8080", Handler: mux()}
	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		shutdownCtx, cancelShutdownCtx := context.WithTimeout(ctx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		cancelShutdownCtx()
		cancel()
	}()

	fmt.Printf("server running on %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-ctx.Done()
}
