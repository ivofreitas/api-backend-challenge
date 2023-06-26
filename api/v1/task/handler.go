package task

import (
	gocontext "context"
	"github.com/google/uuid"
	"github.com/sword/api-backend-challenge/context"
	"github.com/sword/api-backend-challenge/model"
)

type Repository interface {
	Create(ctx gocontext.Context, task *model.Task, username string) error
	List(ctx gocontext.Context) ([]*model.Task, error)
	ListByUsername(ctx gocontext.Context, username string) ([]*model.Task, error)
}

type Publisher interface {
	Publish(str string)
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

func (h *handler) Create(ctx gocontext.Context, param interface{}) (interface{}, error) {
	request := param.(*model.Task)
	request.ID = uuid.New().String()
	username := context.Get(ctx, "username").(string)

	if err := h.repository.Create(ctx, request, username); err != nil {
		return nil, model.ErrorDiscover(err)
	}

	b, err := h.marshal(request)
	if err != nil {
		return nil, model.ErrorDiscover(err)
	}

	go h.publisher.Publish(string(b))

	return model.NewResponse(0, 0, 1, []interface{}{request}), nil
}

func (h *handler) List(ctx gocontext.Context, param interface{}) (interface{}, error) {

	var (
		tasks []*model.Task
		err   error
		role  = context.Get(ctx, "role").(string)
	)

	switch role {
	case model.Manager:
		tasks, err = h.repository.List(ctx)
	case model.Tech:
		username := context.Get(ctx, "username").(string)
		tasks, err = h.repository.ListByUsername(ctx, username)
	}
	if err != nil {
		return nil, model.ErrorDiscover(err)
	}

	if len(tasks) == 0 {
		return nil, model.ErrorDiscover(model.NotFound{DeveloperMessage: "no records in the database"})
	}

	return model.NewResponse(0, 0, len(tasks), []interface{}{tasks}), nil
}
