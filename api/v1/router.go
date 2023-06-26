package v1

import (
	"database/sql"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/sword/api-backend-challenge/api/middleware"
	"github.com/sword/api-backend-challenge/api/swagger"
	"github.com/sword/api-backend-challenge/api/v1/health"
	"github.com/sword/api-backend-challenge/api/v1/task"
	taskmysql "github.com/sword/api-backend-challenge/api/v1/task/mysql"
	"github.com/sword/api-backend-challenge/api/v1/user"
	usermysql "github.com/sword/api-backend-challenge/api/v1/user/mysql"
	"github.com/sword/api-backend-challenge/config"
	"github.com/sword/api-backend-challenge/message_broker/kafka"
	"github.com/sword/api-backend-challenge/model"
	"golang.org/x/crypto/bcrypt"
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

	taskRepo := taskmysql.NewRepository(opts.DB)
	taskHandler := task.NewHandler(taskRepo, opts.Publisher, json.Marshal)
	taskCreate := middleware.NewController(taskHandler.Create, http.StatusCreated, new(model.Task))
	taskList := middleware.NewController(taskHandler.List, http.StatusOK, nil)

	taskGroup := g.Group("/task")
	taskGroup.POST("", taskCreate.Handle, middleware.CheckRole("admin", "manager", "technician"))
	taskGroup.GET("", taskList.Handle, middleware.CheckRole("admin", "manager"))

	userRepo := usermysql.NewRepository(opts.DB)
	userHandler := user.NewHandler(userRepo, bcrypt.GenerateFromPassword, bcrypt.CompareHashAndPassword, env.Authorization.Secret)
	userCreate := middleware.NewController(userHandler.Create, http.StatusCreated, new(model.User))
	userLogin := middleware.NewController(userHandler.Login, http.StatusCreated, new(model.LoginRequest))

	userGroup := g.Group("/user")
	userGroup.POST("", userCreate.Handle, middleware.CheckRole("admin"))
	userGroup.POST("/login", userLogin.Handle)
}
