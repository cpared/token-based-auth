package login

import (
	"context"
	"fmt"
	"time"
	model "token-based-auth/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

type Repository interface {
	Validate(ctx context.Context, ID, pass string) *model.Credentials
}

type Service struct {
	repo Repository
}

func New(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Login(ctx context.Context, secret, ID, pass string) (*model.Credentials, error) {
	cred := s.repo.Validate(ctx, ID, pass)
	if cred.User == "" {
		return cred, fmt.Errorf("invalid user or password")
	}

	exp := time.Now().Add(3600 * time.Second)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": ID,
		"iat": time.Now().Unix(),
		"exp": exp.Unix(),
	})

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return cred, fmt.Errorf("cannot create token, err: %v", err)
	}

	cred.Token = tokenStr
	cred.ExpirationTime = exp
	return cred, nil
}