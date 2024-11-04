package main

import (
	"context"
	"errors"
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
	_ "github.com/joho/godotenv/autoload"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func main() {
	ctx := context.Background()

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
	poll, err := common.NewPostgreSQLPool(cfg)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	if err := poll.Ping(ctx); err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer poll.Close()

	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("goexample"),
		),
	)
	if errors.Is(err, resource.ErrPartialResource) || errors.Is(err, resource.ErrSchemaURLConflict) {
		log.Println(err)
	} else if err != nil {
		log.Fatalf("Failed to create opentelemetry resource: %v\n", err)
	}
	expTrace, err := common.NewTraceExporterConsole()
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	tp := common.NewTracerProvider(expTrace, res)
	defer tp.Shutdown(ctx)
	otel.SetTracerProvider(tp)
	tracer := tp.Tracer("goexample")
	expMetric, err := common.NewMetricExporterConsole()
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	mp := common.NewMeterProvider(expMetric, res)
	defer mp.Shutdown(ctx)
	otel.SetMeterProvider(mp)

	r := chi.NewRouter()
	r.Use(
		middleware.RequestID,
		middleware.Logger,
		otelhttp.NewMiddleware("goexample"),
	)

	petAPI := chi.NewRouter()
	postgresql := postgresql.New(poll)
	petUsecase := usecase.NewPetUsecase(postgresql, tracer)
	petHandler := rest.NewPetHandler(petUsecase, tracer)
	petHandler.Register(petAPI)
	r.Mount("/api/v1/pets", petAPI)

	server := &http.Server{Addr: ":8080", Handler: r}

	// graceful shutdown
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

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-ctx.Done()
}
