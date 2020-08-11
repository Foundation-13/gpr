package user_test

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/foundation-13/gpr/pkg/api/user"
	"github.com/foundation-13/gpr/pkg/api/user/apimocks"
	"github.com/foundation-13/gpr/pkg/types"
)

func TestReviews(t *testing.T) {
	t.Run("succeeded", func(t *testing.T) {
		a := newApi()
		a.man.On("GetReviews", mock.Anything, mock.Anything).Return(types.ReviewsDTO{}, nil)

		a.req.GET("/profile/reviews").SetDebug(true).
			Run(a.ech, func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, resp.Code)

				a.man.AssertCalled(t, "GetReviews", mock.Anything, "from ctx")
		})
	})
}

// helpers

type apiMocks struct {
	req *gofight.RequestConfig
	ech *echo.Echo
	man *apimocks.Manager
}

func newApi() apiMocks {
	req := gofight.New()

	ech := echo.New()
	man := &apimocks.Manager{}

	user.Assemble(ech, man)

	return apiMocks{
		req: req,
		ech: ech,
		man: man,
	}
}