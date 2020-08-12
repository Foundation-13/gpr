package review

import (
	"context"
	"github.com/foundation-13/gpr/pkg/storage/storagemocks"
	"github.com/stretchr/testify/mock"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/foundation-13/gpr/pkg/types"
	"github.com/foundation-13/gpr/pkg/utils/utilsmocks"
)

func TestManager_CreateReview(t *testing.T) {
	validReviewDto := types.ReviewDTO{
		Info:  "info",
		Stars: "5",
	}
	t.Run("succeeded", func(t *testing.T) {
		subject,_ := newManagerWithMocks()
		review, err := subject.CreateReview(context.Background(), "from ctx", validReviewDto)

		assert.NoError(t, err)
		assert.Equal(t, "", review)
	})
}

func TestManager_AddImage(t *testing.T) {
	subject, mocks := newManagerWithMocks()
	mocks.stg.On("UploadObject",mock.Anything, mock.Anything, "image/png", mock.Anything).Return(nil)
	err := subject.AddImage(context.Background(),strings.NewReader(""),"image.png", "image/png")

	assert.NoError(t, err)

}

type mocks_ struct {
	stg       *storagemocks.Client
	idGen     *utilmocks.IDGen
}

func newManagerWithMocks() (Manager, mocks_) {
	m := mocks_{
		stg:       &storagemocks.Client{},
		idGen:     &utilmocks.IDGen{},
	}

	config := Config{
		Storage:        m.stg,
		IDGen:          m.idGen,
	}

	subject := NewManager(config)

	return subject, m
}
