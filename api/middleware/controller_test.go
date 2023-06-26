package middleware

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
	"github.com/sword/api-backend-challenge/log"
	"github.com/sword/api-backend-challenge/mock"
	"github.com/sword/api-backend-challenge/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandleCreate(t *testing.T) {
	testCases := []struct {
		Name           string
		Param          interface{}
		HandlerErr     error
		ExpectedError  error
		ExpectedStatus int
	}{
		{
			Name: "Test Case 1",
			Param: &model.Task{
				Summary:     "Hello. This is my summary",
				PerformedAt: time.Now(),
			},
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:           "Test Case 2",
			Param:          "",
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Name: "Test Case 3",
			Param: &model.Task{
				Summary:     "",
				PerformedAt: time.Time{},
			},
			ExpectedStatus: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			b, _ := json.Marshal(&tc.Param)
			req := httptest.NewRequest(http.MethodPost, "/v1/task", strings.NewReader(string(b)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e := echo.New()
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			ctx := log.InitParams(c.Request().Context())
			c.SetRequest(c.Request().WithContext(ctx))

			e.Binder = NewBinder()
			e.Validator = NewValidator()

			taskHandlerMock := &mock.TaskHandlerMock{}
			taskHandlerMock.
				On("Create", ctx, testifymock.Anything).
				Return(tc.Param, tc.HandlerErr)

			ctrl := NewController(taskHandlerMock.Create, http.StatusCreated, new(model.Task))
			err := ctrl.Handle(c)
			if tc.ExpectedError != nil {
				assert.Error(t, err)
				responseError := err.(*model.ResponseError)
				assert.Equal(t, tc.ExpectedError.Error(), responseError.DeveloperMessage)
				return
			}

			assert.Equal(t, tc.ExpectedStatus, rec.Code)
			assert.NotNil(t, rec.Body)
			assert.NoError(t, err)
		})
	}
}

func TestHandleList(t *testing.T) {
	testCases := []struct {
		Name           string
		HandlerResult  interface{}
		HandlerErr     error
		ExpectedError  error
		ExpectedStatus int
	}{
		{
			Name: "Test Case 1",
			HandlerResult: []*model.Task{
				{
					ID:          "1",
					Summary:     "Hello. This the first record!",
					PerformedAt: time.Now(),
				},
				{
					ID:          "2",
					Summary:     "Hello. This the second record!",
					PerformedAt: time.Now(),
				},
				{
					ID:          "3",
					Summary:     "Hello. This the third record!",
					PerformedAt: time.Now(),
				},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "Test Case 2",
			HandlerResult:  []*model.Task{},
			HandlerErr:     model.NewResponseError("no records in the database", "", http.StatusNotFound),
			ExpectedStatus: http.StatusNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, "/v1/task", nil)

			e := echo.New()
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			ctx := log.InitParams(c.Request().Context())
			c.SetRequest(c.Request().WithContext(ctx))

			taskHandlerMock := &mock.TaskHandlerMock{}
			taskHandlerMock.
				On("List", ctx, nil).
				Return(tc.HandlerResult, tc.HandlerErr)

			ctrl := NewController(taskHandlerMock.List, http.StatusOK, nil)
			err := ctrl.Handle(c)
			if tc.ExpectedError != nil {
				assert.Error(t, err)
				responseError := err.(*model.ResponseError)
				assert.Equal(t, tc.ExpectedError.Error(), responseError.DeveloperMessage)
				return
			}

			assert.Equal(t, tc.ExpectedStatus, rec.Code)
			assert.NotNil(t, rec.Body)
			assert.NoError(t, err)
		})
	}
}
