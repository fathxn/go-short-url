package domain

import (
	"context"
	"go-shorten/internal/model/dto"
	"time"
)

type User struct {
	Id                         string    `db:"id"`
	Name                       string    `db:"name"`
	Email                      string    `db:"email"`
	PasswordHash               string    `db:"password_hash"`
	IsVerified                 bool      `db:"is_verified"`
	VerificationToken          string    `db:"verification_token"`
	VerificationTokenExpiresAt time.Time `db:"verification_token_expires_at"`
	CreatedAt                  time.Time `db:"created_at"`
	UpdatedAt                  time.Time `db:"updated_at"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetById(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByVerificationToken(ctx context.Context, token string) (*User, error)
	UpdateVerificationStatus(ctx context.Context, userId string, isVerified bool) error
	Delete(ctx context.Context, id string) error
}

type UserUsecase interface {
	RegisterUser(ctx context.Context, registerInput *dto.UserRegisterInput) error
	VerifyEmail(ctx context.Context, token string) error
	GetById(ctx context.Context, id string) (*User, error)
	GetURLsByUserId(ctx context.Context, userId string) (*[]URL, error)
	Delete(ctx context.Context, id string) error
}
