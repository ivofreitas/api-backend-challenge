package task

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sword/api-backend-challenge/message_broker/kafka"
	"github.com/sword/api-backend-challenge/model"
)

type handler struct {
	repository Repository
	publisher  *kafka.Publisher
}

func NewHandler(repository Repository) *handler {
	return &handler{repository: repository}
}

func (h *handler) Create(ctx context.Context, param interface{}) (interface{}, error) {
	request := param.(*model.Task)
	request.ID = uuid.New().String()

	if err := h.repository.Create(ctx, request); err != nil {
		return nil, model.ErrorDiscover(err)
	}

	b, err := json.Marshal(request)
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

	return model.NewResponse(0, 0, len(tasks), []interface{}{tasks}), nil
}
