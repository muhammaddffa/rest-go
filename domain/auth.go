package domain

import (
	"context"

	"shellrean.id/belajar-golang-rest-api/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}
