package user

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/sword/api-backend-challenge/model"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type Repository interface {
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
}

type GenerateFromPassword func(password []byte, cost int) ([]byte, error)

type CompareHashAndPassword func(hashedPassword, password []byte) error

type SignedString func(key interface{}) (string, error)

type handler struct {
	repository             Repository
	generateFromPassword   GenerateFromPassword
	compareHashAndPassword CompareHashAndPassword
	secret                 string
}

func NewHandler(repository Repository,
	generateFromPassword GenerateFromPassword,
	compareHashAndPassword CompareHashAndPassword,
	secret string) *handler {
	return &handler{
		repository:             repository,
		generateFromPassword:   generateFromPassword,
		compareHashAndPassword: compareHashAndPassword,
		secret:                 secret}
}

func (h *handler) Create(ctx context.Context, param interface{}) (interface{}, error) {
	request := param.(*model.User)
	request.ID = uuid.New().String()

	user, err := h.repository.GetByUsername(ctx, request.Username)
	if err != nil && !strings.Contains(err.Error(), "sql: no rows in result set") {
		return nil, model.ErrorDiscover(err)
	}
	if user != nil {
		return nil, model.ErrorDiscover(model.Conflict{DeveloperMessage: "Username already present in the database"})
	}

	hashedPassword, err := h.generateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, model.ErrorDiscover(err)
	}
	request.Password = string(hashedPassword)

	if err := h.repository.Create(ctx, request); err != nil {
		return nil, model.ErrorDiscover(err)
	}

	response := new(model.CreateResponse)
	response.ID = request.ID
	response.Username = request.Username
	response.Role = request.Role

	return model.NewResponse(0, 0, 1, []interface{}{response}), nil
}

func (h *handler) Login(ctx context.Context, param interface{}) (interface{}, error) {
	request := param.(*model.LoginRequest)

	user, err := h.repository.GetByUsername(ctx, request.Username)
	if err != nil {
		return nil, model.ErrorDiscover(err)
	}

	if err := h.compareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, model.ErrorDiscover(model.Unauthorized{DeveloperMessage: "Invalid username or password"})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user.Username
	claims["role"] = user.Role

	tokenString, err := token.SignedString([]byte(h.secret))
	if err != nil {
		return nil, model.ErrorDiscover(err)
	}

	return model.NewResponse(0, 0, 1, []interface{}{model.LoginResponse{Token: tokenString}}), nil
}
