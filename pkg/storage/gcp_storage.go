package storage

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"github.com/gavrilaf/errors"

)

func NewGCPBucket(ctx context.Context, bucketName string) (Client, error) {
	stg, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	bucket := stg.Bucket(bucketName)
	if bucket == nil {
		return nil, errors.NotFoundf("couldn't open bucket: %s", bucketName)
	}

	return &client{bucket: bucket, name: bucketName}, nil
}

// impl

type client struct{
	bucket *storage.BucketHandle
	name   string
}

func (c *client) URLByID(id string) string {
	return ""
}

func (c *client) Name() string {
	return ""
}


func (c *client) UploadObject(ctx context.Context, name string, contentType string, r io.Reader) error {
	w := c.bucket.Object(name).NewWriter(ctx)
	// Warning: storage.AllUsers gives public read access to anyone.
	w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	w.ContentType = contentType

	// Entries are immutable, be aggressive about caching (1 day).
	w.CacheControl = "public, max-age=86400"

	if _, err := io.Copy(w, r); err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}

func (c *client) DeleteObject(ctx context.Context, name string) error {
	return nil
}

