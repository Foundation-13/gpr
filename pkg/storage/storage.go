package storage

import (
	"context"
	"io"
)

type URLProvider interface {
	URLByID(id string) string
}

type Client interface {
	Name() string

	UploadObject(ctx context.Context, name string, contentType string, r io.Reader) error
	DeleteObject(ctx context.Context, name string) error
}

