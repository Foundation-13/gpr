package storage

import (
	"context"
	"io"
)

func NewGCPBucket(ctx context.Context, bucketName string) Client {
	return &client{}
}

// impl

type client struct{}

func (c *client) URLByID(id string) string {
	return ""
}

func (c *client) Name() string {
	return ""
}

func (c *client) UploadObject(ctx context.Context, name string, contentType string, r io.Reader) error {
	return nil
}

func (c *client) DeleteObject(ctx context.Context, name string) error {
	return nil
}

