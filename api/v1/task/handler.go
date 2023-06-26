package task

import (
	"context"
	"github.com/google/uuid"
	"github.com/sword/api-backend-challenge/model"
)

type Repository interface {
	Create(ctx context.Context, task *model.Task) error
	List(ctx context.Context) ([]*model.Task, error)
}

type Publisher interface {
	Publish(str string) error
}

type Marshal func(v any) ([]byte, error)

type handler struct {
	repository Repository
	publisher  Publisher
	marshal    Marshal
}

func NewHandler(repository Repository, publisher Publisher, marshal Marshal) *handler {
	return &handler{repository: repository, publisher: publisher, marshal: marshal}
}

func (h *handler) Create(ctx context.Context, param interface{}) (interface{}, error) {
	request := param.(*model.Task)
	request.ID = uuid.New().String()

	if err := h.repository.Create(ctx, request); err != nil {
		return nil, model.ErrorDiscover(err)
	}

	b, err := h.marshal(request)
	if err != nil {
		return nil, model.ErrorDiscover(err)
	}

	err = h.publisher.Publish(string(b))
	if err != nil {
		return nil, model.ErrorDiscover(err)
	}

	return model.NewResponse(0, 0, 1, []interface{}{request}), nil
}

func (h *handler) List(ctx context.Context, param interface{}) (interface{}, error) {

	tasks, err := h.repository.List(ctx)
	if err != nil {
		return nil, model.ErrorDiscover(err)
	}

	if len(tasks) == 0 {
		return nil, model.ErrorDiscover(model.NotFound{DeveloperMessage: "no records in the database"})
	}

	return model.NewResponse(0, 0, len(tasks), []interface{}{tasks}), nil
}
