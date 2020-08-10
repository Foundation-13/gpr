package content

import (
	"context"
	"github.com/foundation-13/gpr/pkg/types"
)

type Manager interface {
	Create(ctx context.Context, dto types.ReviewDTO)(ID string, err error)
}


func NewManager() Manager{
	return &manager{}
}

type manager struct{
}
func(man *manager) Create(ctx context.Context, dto types.ReviewDTO)(ID string, err error){
	return
}
