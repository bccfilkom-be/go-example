package usecase

import (
	"context"

	"github.com/bccfilkom-be/go-example/opentelemetry/db/postgresql"
	"github.com/bccfilkom-be/go-example/opentelemetry/pet/dto"
)

type IPetUsecase interface {
	ListPets(ctx context.Context, offset, limit int32) ([]dto.Pet, error)
	GetPet(ctx context.Context, id int64) (dto.Pet, error)
	CreatePet(ctx context.Context, pet dto.Pet) error
	UpdatePet(ctx context.Context, pet dto.Pet) error
	DeletePet(ctx context.Context, id int64) error
}

type usecase struct {
	postgresql postgresql.Queries
}

func NewPetUsecase(postgresql postgresql.Queries) IPetUsecase {
	return &usecase{postgresql}
}

func (u *usecase) ListPets(ctx context.Context, offset, limit int32) ([]dto.Pet, error) {
	pets, err := u.postgresql.ListPets(ctx, postgresql.ListPetsParams{Offset: offset, Limit: limit})
	if err != nil {
		return nil, err
	}
	var _pets []dto.Pet
	for _, pet := range pets {
		_pets = append(_pets, dto.Pet{
			ID:       pet.ID,
			Name:     pet.Name,
			PhotoURL: pet.Photourl,
			Sold:     pet.Sold,
		})
	}
	return _pets, nil
}

func (u *usecase) GetPet(ctx context.Context, id int64) (dto.Pet, error) {
	pet, err := u.postgresql.GetPet(ctx, id)
	if err != nil {
		return dto.Pet{}, err
	}
	_pet := dto.Pet{
		ID:       pet.ID,
		Name:     pet.Name,
		PhotoURL: pet.Photourl,
		Sold:     pet.Sold,
	}
	return _pet, nil
}

func (u *usecase) CreatePet(ctx context.Context, pet dto.Pet) error {
	if _, err := u.postgresql.CreatePet(ctx, postgresql.CreatePetParams{Name: pet.Name, Photourl: pet.PhotoURL}); err != nil {
		return err
	}
	return nil
}

func (u *usecase) UpdatePet(ctx context.Context, pet dto.Pet) error {
	if err := u.postgresql.UpdatePet(ctx, postgresql.UpdatePetParams{ID: pet.ID, Name: pet.Name}); err != nil {
		return err
	}
	return nil
}

func (u *usecase) DeletePet(ctx context.Context, id int64) error {
	return u.postgresql.DeletePet(ctx, id)
}
