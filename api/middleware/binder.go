package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sword/api-backend-challenge/model"
)

type Binder struct{}

func NewBinder() *Binder {
	return &Binder{}
}

func (cb *Binder) Bind(i interface{}, c echo.Context) error {
	db := new(echo.DefaultBinder)

	if err := db.Bind(i, c); err != nil {
		return model.ErrorDiscover(model.BadRequest{DeveloperMessage: err.Error()})
	}
	return nil
}
