package login

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/sword/api-backend-challenge/config"
	"github.com/sword/api-backend-challenge/model"
)

type handler struct {
	repository Repository
}

func NewHandler(repository Repository) *handler {
	return &handler{repository: repository}
}

func (h *handler) Login(ctx context.Context, param interface{}) (interface{}, error) {
	request := param.(*model.LoginRequest)

	user, err := h.repository.GetByUsernameAndPassword(ctx, request.Username, request.Password)
	if err != nil {
		return nil, model.ErrorDiscover(model.Unauthorized{})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user.Username
	claims["role"] = user.Role

	secret := config.GetEnv().Authorization.Secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, model.ErrorDiscover(err)
	}

	return model.NewResponse(0, 0, 1, []interface{}{model.LoginResponse{Token: tokenString}}), nil
}
