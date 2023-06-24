package task

import (
	"context"
	"github.com/sword/api-backend-challenge/model"
)

type Repository interface {
	Create(ctx context.Context, task *model.Task) error
	List(ctx context.Context) ([]*model.Task, error)
}
