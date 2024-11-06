package usecase

import (
	"context"

	"github.com/bccfilkom-be/go-example/magic_dependency_injector/db/postgresql"
)

type IBookUsecase interface {
	CreateBook(ctx context.Context, title string) (int64, error)
	ListBooks(ctx context.Context) ([]postgresql.Book, error)
}

type usecase struct {
	postgres *postgresql.Queries
}

func NewBookUsecase(postgres *postgresql.Queries) IBookUsecase {
	return &usecase{postgres}
}

func (u *usecase) CreateBook(ctx context.Context, title string) (int64, error) {
	return u.postgres.CreatePet(ctx, title)
}

func (u *usecase) ListBooks(ctx context.Context) ([]postgresql.Book, error) {
	return u.postgres.ListBooks(ctx)
}
