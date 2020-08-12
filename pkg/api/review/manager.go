package review

import (
	"context"

	"github.com/foundation-13/gpr/pkg/types"
)

//go:generate mockery -name Manager -outpkg reviewmocks -output ./reviewmocks -dir .
type Manager interface {
	CreateReview(ctx context.Context, userID string, review types.ReviewDTO) (string, error)
}

func NewManager() Manager {
	return &manager{}
}

// impl

type manager struct {
}

func (man *manager) CreateReview(ctx context.Context, userID string, review types.ReviewDTO) (string, error) {
	return "", nil
}
