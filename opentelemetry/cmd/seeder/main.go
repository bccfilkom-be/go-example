package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/bccfilkom-be/go-example/opentelemetry/db/postgresql"
	"github.com/go-faker/faker/v4"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}
}

func main() {
	ctx := context.Background()
	uri := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	cfg, err := pgx.ParseConfig(uri)
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := postgresql.NewConn(ctx, cfg)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer conn.Close(ctx)

	pets := pets()

	_, err = conn.CopyFrom(
		context.Background(),
		pgx.Identifier{"pets"},
		[]string{"name", "photo_url", "sold"},
		pgx.CopyFromSlice(len(pets), func(i int) ([]any, error) {
			return []any{pets[i].Name, pets[i].PhotoUrl, pets[i].Sold}, nil
		}),
	)
	if err != nil {
		log.Fatalf("Error to seed pets: %v\n", err)
	}

}

func users() []postgresql.User {
	users := make([]postgresql.User, 1000)
	for i := 0; i < 1000; i++ {
		users[i] = postgresql.User{
			Email:    faker.Email(),
			Username: faker.Username(),
			Provider: postgresql.OauthProviderGoogle,
		}
	}
	return users
}

func buyers() []postgresql.Buyer {
	buyers := make([]postgresql.Buyer, 1000)
	for i := 0; i < 1000; i++ {
		buyers[i] = postgresql.Buyer{
			ID:       int64(i + 1),
			UserID:   pgtype.Int8{Int64: int64(rand.Intn(1000) + 1), Valid: true},
			Username: faker.Username(),
		}
	}
	return buyers
}

func categories() []postgresql.Category {
	categories := make([]postgresql.Category, 1000)
	for i := 0; i < 1000; i++ {
		categories[i] = postgresql.Category{
			ID:   int64(i + 1),
			Name: faker.Word(),
		}
	}
	return categories
}

func orders() []postgresql.Order {
	orders := make([]postgresql.Order, 1000)
	statusOptions := []postgresql.OrderStatus{postgresql.OrderStatusSubmitted, postgresql.OrderStatusPaid, postgresql.OrderStatusShipped, postgresql.OrderStatusCancelled}
	for i := 0; i < 1000; i++ {
		orders[i] = postgresql.Order{
			ID:              int64(i + 1),
			BuyerID:         pgtype.Int8{Int64: int64(rand.Intn(1000) + 1), Valid: true},
			Status:          statusOptions[rand.Intn(len(statusOptions))],
			Description:     pgtype.Text{String: faker.Sentence(), Valid: true},
			ShippingAddress: faker.Word(),
		}
	}
	return orders
}

func pets() []postgresql.Pet {
	pets := make([]postgresql.Pet, 1000)
	for i := 0; i < 1000; i++ {
		pets[i] = postgresql.Pet{
			Name:     faker.Name(),
			PhotoUrl: faker.URL(),
			Sold:     rand.Intn(2) == 1,
		}
	}
	return pets
}

func tags() []postgresql.Tag {
	tags := make([]postgresql.Tag, 1000)
	for i := 0; i < 1000; i++ {
		tags[i] = postgresql.Tag{
			ID:   int32(i + 1),
			Name: faker.Word(),
		}
	}
	return tags
}

func baskets() []postgresql.Basket {
	baskets := make([]postgresql.Basket, 1000)
	for i := 0; i < 1000; i++ {
		baskets[i] = postgresql.Basket{
			ID:      int64(i + 1),
			BuyerID: pgtype.Int8{Int64: int64(rand.Intn(1000) + 1), Valid: true},
		}
	}
	return baskets
}

func basketItems() []postgresql.BasketItem {
	items := make([]postgresql.BasketItem, 1000)
	for i := 0; i < 1000; i++ {
		items[i] = postgresql.BasketItem{
			ID:       int64(i + 1),
			BasketID: pgtype.Int8{Int64: int64(rand.Intn(1000) + 1), Valid: true},
			PetID:    pgtype.Int8{Int64: int64(rand.Intn(1000) + 1), Valid: true},
		}
	}
	return items
}

func orderItems() []postgresql.OrderItem {
	items := make([]postgresql.OrderItem, 1000)
	for i := 0; i < 1000; i++ {
		items[i] = postgresql.OrderItem{
			ID:       int64(i + 1),
			Name:     faker.Name(),
			PhotoUrl: faker.URL(),
			Category: faker.Word(),
		}
	}
	return items
}

func petTags() []postgresql.PetTag {
	tags := make([]postgresql.PetTag, 1000)
	for i := 0; i < 1000; i++ {
		tags[i] = postgresql.PetTag{
			PetID: pgtype.Int8{Int64: int64(rand.Intn(1000) + 1), Valid: true},
			TagID: pgtype.Int8{Int64: int64(rand.Intn(1000) + 1), Valid: true},
		}
	}
	return tags
}
