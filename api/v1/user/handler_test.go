package user

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
	"github.com/sword/api-backend-challenge/mock"
	"github.com/sword/api-backend-challenge/model"
	"testing"
)

var (
	repoErr        = errors.New("repository failed")
	userErr        = errors.New("username already present in the database")
	passHashErr    = errors.New("passHash failed")
	invalidPassErr = errors.New("Invalid username or password")
)

const (
	secret = "f22597e87a7ccd96e732eb0e2abbf90e"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		Name          string
		Request       *model.User
		PassHashErr   error
		GetRepoUser   *model.User
		GetRepoErr    error
		CreateRepoErr error
		ExpectedError error
	}{
		{
			Name: "Test Case 1",
			Request: &model.User{
				Username: "john.doe",
				Password: "123456",
				Role:     "admin",
			},
		},
		{
			Name: "Test Case 2",
			Request: &model.User{
				Username: "john.doe",
				Password: "123456",
				Role:     "admin",
			},
			GetRepoErr:    repoErr,
			ExpectedError: repoErr,
		},
		{
			Name: "Test Case 3",
			Request: &model.User{
				Username: "john.doe",
				Password: "123456",
				Role:     "admin",
			},
			GetRepoUser: &model.User{
				Username: "john.doe",
			},
			GetRepoErr:    userErr,
			ExpectedError: userErr,
		},
		{
			Name: "Test Case 4",
			Request: &model.User{
				Username: "john.doe",
				Password: "123456",
				Role:     "admin",
			},
			PassHashErr:   passHashErr,
			ExpectedError: passHashErr,
		},
		{
			Name: "Test Case 5",
			Request: &model.User{
				Username: "john.doe",
				Password: "123456",
				Role:     "admin",
			},
			CreateRepoErr: repoErr,
			ExpectedError: repoErr,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()

			repositoryMock := &mock.UserRepositoryMock{}

			repositoryMock.
				On("GetByUsername", ctx, tc.Request.Username).
				Return(tc.GetRepoUser, tc.GetRepoErr)

			repositoryMock.
				On("Create", ctx, tc.Request).
				Return(tc.CreateRepoErr)

			passHashMock := mock.PassHashMock{}
			passHashMock.
				On("GenerateFromPassword", testifymock.Anything, testifymock.Anything).
				Return([]byte{}, tc.PassHashErr)

			hdl := NewHandler(repositoryMock, passHashMock.GenerateFromPassword, passHashMock.CompareHashAndPassword, secret)
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

func TestLogin(t *testing.T) {
	testCases := []struct {
		Name          string
		Request       *model.LoginRequest
		PassHashErr   error
		GetRepoUser   *model.User
		GetRepoErr    error
		CreateRepoErr error
		ExpectedError error
	}{
		{
			Name: "Test Case 1",
			Request: &model.LoginRequest{
				Username: "john.doe",
				Password: "123456",
			},
			GetRepoUser: &model.User{
				Username: "john.doe",
				Password: "123456",
				Role:     "admin",
			},
		},
		{
			Name: "Test Case 2",
			Request: &model.LoginRequest{
				Username: "john.doe",
				Password: "123456",
			},
			GetRepoErr:    repoErr,
			ExpectedError: repoErr,
		},
		{
			Name: "Test Case 3",
			Request: &model.LoginRequest{
				Username: "john.doe",
				Password: "123456",
			},
			GetRepoUser: &model.User{
				Username: "john.doe",
				Password: "123456",
				Role:     "admin",
			},
			PassHashErr:   invalidPassErr,
			ExpectedError: invalidPassErr,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()

			repositoryMock := &mock.UserRepositoryMock{}

			repositoryMock.
				On("GetByUsername", ctx, tc.Request.Username).
				Return(tc.GetRepoUser, tc.GetRepoErr)

			passHashMock := mock.PassHashMock{}
			passHashMock.
				On("CompareHashAndPassword", testifymock.Anything, testifymock.Anything).
				Return(tc.PassHashErr)

			hdl := NewHandler(repositoryMock, passHashMock.GenerateFromPassword, passHashMock.CompareHashAndPassword, secret)
			response, err := hdl.Login(ctx, tc.Request)
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
