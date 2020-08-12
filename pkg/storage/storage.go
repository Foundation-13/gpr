package storage

import (
	"context"
	"io"
)

type URLProvider interface {
	URLByID(id string) string
}

//go:generate mockery -name Client -outpkg storagemocks -output ./storagemocks -dir .
type Client interface {
	Name() string

	UploadObject(ctx context.Context, name string, contentType string, r io.Reader) error
	DeleteObject(ctx context.Context, name string) error
}

