package task

import (
	gocontext "context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
	"github.com/sword/api-backend-challenge/context"
	"github.com/sword/api-backend-challenge/mock"
	"github.com/sword/api-backend-challenge/model"
)

var (
	repositoryErr   = errors.New("repository failed")
	publisherErr    = errors.New("publisher failed")
	marshalErr      = errors.New("marshal failed")
	cryptoErr       = errors.New("crypto failed")
	encryptedResult = "encrypted"
	decryptedResult = "decrypted"
)

const (
	username = "joe.doe"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		Name          string
		Request       *model.Task
		Role          string
		RepositoryErr error
		MarshalErr    error
		PublisherErr  error
		CryptoResult  string
		CryptoErr     error
		ExpectedError error
	}{
		{
			Name: "Test Case 1",
			Request: &model.Task{
				Summary:     "Hello manager. This is my summary",
				PerformedAt: time.Now(),
			},
			CryptoResult: encryptedResult,
		},
		{
			Name: "Test Case 2",
			Request: &model.Task{
				Summary:     "Hello technician. This is my summary",
				PerformedAt: time.Now(),
			},
			CryptoResult: encryptedResult,
		},
		{
			Name: "Test Case 3",
			Request: &model.Task{
				Summary:     "So sad! Repository is going to fail!",
				PerformedAt: time.Now(),
			},
			RepositoryErr: repositoryErr,
			ExpectedError: repositoryErr,
		},
		{
			Name: "Test Case 5",
			Request: &model.Task{
				Summary:     "So sad! Crypto is going to fail!",
				PerformedAt: time.Now(),
			},
			CryptoErr:     cryptoErr,
			ExpectedError: cryptoErr,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := gocontext.Background()
			ctx = context.Set(ctx, "username", username)
			ctx = context.Set(ctx, "role", tc.Role)

			repositoryMock := &mock.TaskRepositoryMock{}
			repositoryMock.
				On("Create", ctx, tc.Request, username).
				Return(tc.RepositoryErr)

			jsonMock := &mock.JsonMock{}
			jsonMock.
				On("Marshal", tc.Request).
				Return([]byte{}, tc.MarshalErr)

			publisherMock := &mock.PublisherMock{}
			publisherMock.
				On("Publish", testifymock.Anything).
				Return(tc.PublisherErr)

			cryptoMock := &mock.CryptoMock{}
			cryptoMock.
				On("Encrypt", testifymock.Anything).
				Return(tc.CryptoResult, tc.CryptoErr)

			hdl := NewHandler(repositoryMock, publisherMock, jsonMock.Marshal, cryptoMock)
			response, err := hdl.Create(ctx, tc.Request)
			if tc.ExpectedError != nil {
				assert.Error(t, err)
				responseError := err.(*model.ResponseError)
				assert.Equal(t, tc.ExpectedError.Error(), responseError.DeveloperMessage)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, response)
		})
	}
}

func TestList(t *testing.T) {
	testCases := []struct {
		Name             string
		Role             string
		RepositoryResult []*model.Task
		RepositoryErr    error
		PublisherErr     error
		CryptoResult     string
		CryptoErr        error
		ExpectedError    error
	}{
		{
			Name: "Test Case 1",
			RepositoryResult: []*model.Task{
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
			Role:         model.Manager,
			CryptoResult: decryptedResult,
		},
		{
			Name: "Test Case 2",
			RepositoryResult: []*model.Task{
				{
					ID:          "1",
					Summary:     "Hello. This the first record!",
					PerformedAt: time.Now(),
				},
			},
			Role:         model.Tech,
			CryptoResult: decryptedResult,
		},
		{
			Name:          "Test Case 2",
			RepositoryErr: repositoryErr,
			ExpectedError: repositoryErr,
			Role:          model.Manager,
		},
		{
			Name:             "Test Case 3",
			RepositoryResult: []*model.Task{},
			ExpectedError:    errors.New("no records in the database"),
		},
		{
			Name: "Test Case 2",
			RepositoryResult: []*model.Task{
				{
					ID:          "1",
					Summary:     "Hello. This the first record!",
					PerformedAt: time.Now(),
				},
			},
			Role:          model.Tech,
			CryptoErr:     cryptoErr,
			ExpectedError: cryptoErr,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := gocontext.Background()
			ctx = context.Set(ctx, "username", username)
			ctx = context.Set(ctx, "role", tc.Role)

			repositoryMock := &mock.TaskRepositoryMock{}
			switch tc.Role {
			case model.Manager:
				repositoryMock.
					On("List", ctx).
					Return(tc.RepositoryResult, tc.RepositoryErr)
			case model.Tech:
				repositoryMock.
					On("ListByUsername", ctx, username).
					Return(tc.RepositoryResult, tc.RepositoryErr)
			}

			jsonMock := mock.JsonMock{}

			cryptoMock := &mock.CryptoMock{}
			cryptoMock.
				On("Decrypt", testifymock.Anything).
				Return(tc.CryptoResult, tc.CryptoErr)

			hdl := NewHandler(repositoryMock, &mock.PublisherMock{}, jsonMock.Marshal, cryptoMock)
			response, err := hdl.List(ctx, nil)
			if tc.ExpectedError != nil {
				assert.Error(t, err)
				responseError := err.(*model.ResponseError)
				assert.Equal(t, tc.ExpectedError.Error(), responseError.DeveloperMessage)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, response)
		})
	}
}
