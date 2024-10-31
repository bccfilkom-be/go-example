package dto

import (
	"context"

	"github.com/go-playground/validator/v10"
)

type CreatePet struct {
	Name     string `json:"name"`
	PhotoURL string `json:"photoURL"`
}

func (p *CreatePet) Validate(ctx context.Context) error {
	validate := validator.New()

	for _, fields := range []struct {
		val  interface{}
		rule string
	}{
		{p.Name, "lowercase"},
		{p.PhotoURL, "url"},
	} {
		if err := validate.VarCtx(ctx, fields.val, fields.rule); err != nil {
			return err
		}
	}

	return nil
}
