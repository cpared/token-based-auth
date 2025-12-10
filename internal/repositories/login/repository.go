package login

import (
	"context"
	model "token-based-auth/internal/domain"

)

type Repository struct {
	Users map[string]User
}

func New() *Repository {
	return &Repository{
		Users: map[string]User{
			"1": {
				ID:       "1",
				Name:     "Cris",
				Password: "12345",
				Role:     "admin",
			},
		},
	}
}

func (r *Repository) Validate(ctx context.Context, ID, pass string) *model.Credentials {
	usr, found := r.Users[ID]
	if !found {
		return &model.Credentials{}
	}
	return &model.Credentials{
		User: usr.Name,
		Role: usr.Role,
	}
}
