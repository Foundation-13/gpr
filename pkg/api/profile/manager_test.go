package profile_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/foundation-13/gpr/pkg/api/profile"
	"github.com/foundation-13/gpr/pkg/types"
)

func TestManager_GetReviews(t *testing.T) {
	t.Run("succeeded", func(t *testing.T) {
		man := profile.NewManager()

		reviews, err := man.GetReviews(context.Background(), "from ctx")

		assert.NoError(t, err)
		assert.Equal(t, types.ReviewsDTO{}, reviews)
	})
}
