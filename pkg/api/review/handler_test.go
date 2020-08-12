package review_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/foundation-13/gpr/pkg/api/middleware"
	"github.com/foundation-13/gpr/pkg/api/middleware/middlewaremocks"
	"github.com/foundation-13/gpr/pkg/api/review"
	"github.com/foundation-13/gpr/pkg/api/review/reviewmocks"
)

func TestCreateReview(t *testing.T) {

	validReviewJson := `{"info":"info","stars":"5"}`
	invalidJson := `{invalid}`

	t.Run("happy path", func(t *testing.T) {
		subj := prepareTest()

		subj.manager.On("CreateReview", mock.Anything, mock.Anything, mock.Anything).Return("1", nil)

		subj.req.POST("/reviews").
			SetDebug(true).
			SetHeader(map[string]string{"Authorization": "Bearer 1"}).
			SetBody(validReviewJson).
			Run(subj.ech, func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
				assert.Equal(t, http.StatusCreated, resp.Code)
				// TODO: check id
				// TODO: check manager call

			})
	})

	t.Run("invalid JSON", func(t *testing.T) {
		subj := prepareTest()

		subj.req.POST("/reviews").
			SetDebug(true).
			SetHeader(map[string]string{"Authorization": "Bearer 1"}).
			SetBody(invalidJson).
			Run(subj.ech, func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
				assert.Equal(t, http.StatusBadRequest, resp.Code)
			})
	})

	t.Run("manager return an error", func(t *testing.T) {
		subj := prepareTest()

		subj.manager.On("CreateReview", mock.Anything, mock.Anything, mock.Anything).Return("", fmt.Errorf(""))

		subj.req.POST("/reviews").
			SetDebug(true).
			SetHeader(map[string]string{"Authorization": "Bearer 1"}).
			SetBody(validReviewJson).
			Run(subj.ech, func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
				assert.Equal(t, 500, resp.Code)
				// TODO: check the error
			})
	})
}

func TestAddImage(t *testing.T) {
	subj := prepareTest()

	subj.manager.On("AddImage", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return( nil)

	subj.req.PUT("/reviews/upload").
		SetDebug(true).
		SetHeader(map[string]string{"Authorization": "Bearer 123"}).
		SetFileFromPath([]gofight.UploadFile{
			{
				Path:    "/images/media.png",
				Name:    "file",
				Content: []byte("123"),
			},
		}).
		Run(subj.ech, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

// helperes

const (
	testUserID = "user-id"
)

type mocks struct {
	req     *gofight.RequestConfig
	ech     *echo.Echo
	manager *reviewmocks.Manager
}

func prepareTest() *mocks {
	req := gofight.New()

	ech := echo.New()
	manager := &reviewmocks.Manager{}

	verifier := &middlewaremocks.TokenVerifier{}
	verifier.On("VerifyToken", mock.Anything).Return(testUserID, nil)
	auth := middleware.NewAuth(verifier)

	review.Assemble(ech, manager, auth.MiddlewareFunc)

	return &mocks{
		req:     req,
		ech:     ech,
		manager: manager,
	}
}
