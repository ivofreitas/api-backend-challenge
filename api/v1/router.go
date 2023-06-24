package v1

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/sword/api-backend-challenge/api/middleware"
	"github.com/sword/api-backend-challenge/api/swagger"
	"github.com/sword/api-backend-challenge/api/v1/health"
	"github.com/sword/api-backend-challenge/api/v1/task"
	"github.com/sword/api-backend-challenge/api/v1/task/mysql"
	"github.com/sword/api-backend-challenge/config"
	"github.com/sword/api-backend-challenge/message_broker/kafka"
	"github.com/sword/api-backend-challenge/model"
	"net/http"
)

type Option struct {
	DB        *sql.DB
	Publisher *kafka.Publisher
}

func Register(g *echo.Group, opts Option) {

	env := config.GetEnv()
	doc := env.Doc
	swagger.Register(swagger.Options{
		Title:       doc.Title,
		Description: doc.Description,
		Version:     doc.Version,
		BasePath:    env.Server.BasePath,
		Group:       g.Group("/swagger"),
	})

	g.GET("/health", health.Handle)

	taskRepo := mysql.NewRepository(opts.DB)
	taskHandler := task.NewHandler(taskRepo)
	taskCreate := middleware.NewController(taskHandler.Create, http.StatusCreated, new(model.Task))
	taskList := middleware.NewController(taskHandler.List, http.StatusOK, nil)

	taskGroup := g.Group("/task")
	taskGroup.POST("", taskCreate.Handle, middleware.CheckRole("admin", "manager", "technician"))
	taskGroup.GET("", taskList.Handle, middleware.CheckRole("admin", "manager"))
}
