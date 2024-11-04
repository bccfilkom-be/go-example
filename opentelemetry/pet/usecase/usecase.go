package usecase

import (
	"context"

	"github.com/bccfilkom-be/go-example/opentelemetry/db/postgresql"
	"github.com/bccfilkom-be/go-example/opentelemetry/pet/dto"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var paginationSize = int32(25)

type IPetUsecase interface {
	ListPets(ctx context.Context, offset, limit int32) ([]dto.Pet, error)
	GetPet(ctx context.Context, id int64) (dto.Pet, error)
	CreatePet(ctx context.Context, pet *dto.Pet) error
	UpdatePet(ctx context.Context, pet *dto.Pet) error
	DeletePet(ctx context.Context, id int64) error
}

type usecase struct {
	postgresql *postgresql.Queries
	tracer     trace.Tracer
	_log       *zap.Logger
}

func NewPetUsecase(postgresql *postgresql.Queries, tracer trace.Tracer, _log *zap.Logger) IPetUsecase {
	return &usecase{postgresql, tracer, _log}
}

func (u *usecase) ListPets(ctx context.Context, page, size int32) ([]dto.Pet, error) {
	ctx, span := u.tracer.Start(ctx, "ListPetsUsecase")
	defer span.End()

	pets, err := u.postgresql.ListPets(ctx, postgresql.ListPetsParams{Offset: page, Limit: paginationSize})
	if err != nil {
		return nil, err
	}
	var _pets []dto.Pet
	for _, pet := range pets {
		_pets = append(_pets, dto.Pet{
			ID:       pet.ID,
			Name:     pet.Name,
			PhotoURL: pet.PhotoUrl,
			Sold:     pet.Sold,
		})
	}
	return _pets, nil
}

func (u *usecase) GetPet(ctx context.Context, id int64) (dto.Pet, error) {
	ctx, span := u.tracer.Start(ctx, "GetPetUsecase")
	defer span.End()

	pet, err := u.postgresql.GetPet(ctx, id)
	if err != nil {
		return dto.Pet{}, err
	}
	_pet := dto.Pet{
		ID:       pet.ID,
		Name:     pet.Name,
		PhotoURL: pet.PhotoUrl,
		Sold:     pet.Sold,
	}
	return _pet, nil
}

func (u *usecase) CreatePet(ctx context.Context, pet *dto.Pet) error {
	ctx, span := u.tracer.Start(ctx, "CreatePetUsecase")
	defer span.End()

	if _, err := u.postgresql.CreatePet(ctx, postgresql.CreatePetParams{Name: pet.Name, PhotoUrl: pet.PhotoURL}); err != nil {
		return err
	}
	u._log.Info("pet created", zap.String("name", pet.Name))
	return nil
}

func (u *usecase) UpdatePet(ctx context.Context, pet *dto.Pet) error {
	ctx, span := u.tracer.Start(ctx, "UpdatePetUsecase")
	defer span.End()

	if err := u.postgresql.UpdatePet(ctx, postgresql.UpdatePetParams{ID: pet.ID, Name: pet.Name}); err != nil {
		return err
	}
	u._log.Info("pet updated", zap.String("name", pet.Name))
	return nil
}

func (u *usecase) DeletePet(ctx context.Context, id int64) error {
	ctx, span := u.tracer.Start(ctx, "DeletePetUsecase")
	defer span.End()

	u._log.Info("pet deleted", zap.Int64("id", id))
	return u.postgresql.DeletePet(ctx, id)
}
