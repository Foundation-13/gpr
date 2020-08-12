package review

import (
	"context"
	"io"
	"path"

	"github.com/foundation-13/gpr/pkg/log"
	"github.com/foundation-13/gpr/pkg/storage"
	"github.com/foundation-13/gpr/pkg/types"
	"github.com/foundation-13/gpr/pkg/utils"
)

//go:generate mockery -name Manager -outpkg reviewmocks -output ./reviewmocks -dir .
type Manager interface {
	CreateReview(ctx context.Context, userID string, review types.ReviewDTO) (string, error)
	AddImage(ctx context.Context, r io.Reader, fileName string, contentType string) error
}


type Config struct {
	Storage        storage.Client
	IDGen          utils.IDGen
}

func NewManager(config Config) Manager {
	return &manager{
		stg:   config.Storage,
		idGen: config.IDGen,
	}
}

// impl

type manager struct {
	stg   storage.Client
	idGen utils.IDGen
}

func (man *manager) CreateReview(ctx context.Context, userID string, review types.ReviewDTO) (string, error) {
	return "", nil
}

func (man *manager) AddImage(ctx context.Context, r io.Reader, fileName string, contentType string) error {
	log.InitLog(true)
	var imageName = "image-" + path.Base(fileName)
	if err := man.stg.UploadObject(ctx, imageName, contentType, r); err != nil {
		log.L.WithError(err).Error("failed to upload image to the storage")
		return err
	}
	return nil
}