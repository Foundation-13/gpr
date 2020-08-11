package review_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/foundation-13/gpr/pkg/api/review"
	"github.com/foundation-13/gpr/pkg/api/review/reviewmocks"
)

func TestCreateReview(t *testing.T) {

	validReviewJson := `{"info":"info","stars":"5"}`
	invalidJson := `{invalid}`

	t.Run("happy path", func(t *testing.T) {
		subj := prepareTest()

		subj.manager.On("CreateReview", mock.Anything, mock.Anything).Return("1", nil)

		subj.req.POST("/reviews/").
			SetDebug(true).
			SetBody(validReviewJson).
			Run(subj.ech, func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
				assert.Equal(t, http.StatusCreated, resp.Code)
				// TODO: check id
				// TODO: check manager call

			})
	})

	t.Run("invalid JSON", func(t *testing.T) {
		subj := prepareTest()

		subj.req.POST("/reviews/").
			SetDebug(true).
			SetBody(invalidJson).
			Run(subj.ech, func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
				assert.Equal(t, http.StatusBadRequest, resp.Code)
			})
	})

	t.Run("manager return an error", func(t *testing.T) {
		subj := prepareTest()

		subj.manager.On("CreateReview", mock.Anything, mock.Anything).Return("", fmt.Errorf(""))

		subj.req.POST("/reviews/").
			SetDebug(true).
			SetBody(validReviewJson).
			Run(subj.ech, func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
				assert.Equal(t, 500, resp.Code)
				// TODO: check the error
			})
	})
}

// helperes

type mocks struct {
	req     *gofight.RequestConfig
	ech     *echo.Echo
	manager *reviewmocks.Manager
}

func prepareTest() *mocks {
	req := gofight.New()

	ech := echo.New()
	manager := &reviewmocks.Manager{}

	review.Assemble(ech, manager)

	return &mocks{
		req:     req,
		ech:     ech,
		manager: manager,
	}
}
