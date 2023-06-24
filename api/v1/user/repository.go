package user

import (
	"context"
	"github.com/sword/api-backend-challenge/model"
)

type Repository interface {
	GetByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error)
}
