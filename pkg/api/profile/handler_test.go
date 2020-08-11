package profile_test

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/foundation-13/gpr/pkg/api/profile"
	"github.com/foundation-13/gpr/pkg/api/profile/profilemocks"
	"github.com/foundation-13/gpr/pkg/types"
)

func TestReviews(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		subj := prepareTest()

		subj.manager.On("GetReviews", mock.Anything, mock.Anything).Return(types.ReviewsDTO{}, nil)

		subj.req.GET("/profile/reviews").SetDebug(true).
			Run(subj.ech, func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, resp.Code)

			subj.manager.AssertCalled(t, "GetReviews", mock.Anything, "from ctx")
		})
	})
}

// helpers

type mocks struct {
	req *gofight.RequestConfig
	ech *echo.Echo
	manager *profilemocks.Manager
}

func prepareTest() *mocks {
	req := gofight.New()

	ech := echo.New()
	manager := &profilemocks.Manager{}

	profile.Assemble(ech, manager)

	return &mocks{
		req: req,
		ech: ech,
		manager: manager,
	}
}