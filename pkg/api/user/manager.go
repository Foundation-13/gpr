package user

import (
	"context"
	"github.com/foundation-13/gpr/pkg/types"
)

//go:generate mockery -name Manager -outpkg apimocks -output ./apimocks -dir .
type Manager interface {
	GetReviews(ctx context.Context, userID string) (types.ReviewsDTO, error)
}

func NewManager() Manager{
	return &manager{}
}

type manager struct{
}

func (m *manager) GetReviews(ctx context.Context, userID string) (types.ReviewsDTO, error) {
	return types.ReviewsDTO{}, nil
}