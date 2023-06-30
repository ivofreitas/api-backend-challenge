package api

import (
	gocontext "context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/sword/api-backend-challenge/api/middleware"
	"github.com/sword/api-backend-challenge/api/v1"
	"github.com/sword/api-backend-challenge/config"
	"github.com/sword/api-backend-challenge/context"
	"github.com/sword/api-backend-challenge/cypher"
	"github.com/sword/api-backend-challenge/db/mysql"
	"github.com/sword/api-backend-challenge/log"
	"github.com/sword/api-backend-challenge/message_broker/kafka"
	"github.com/sword/api-backend-challenge/model"
)

type Server struct {
	http      *echo.Echo
	db        *sql.DB
	publisher *kafka.Publisher
	crypto    *cypher.Crypto
	logger    *logrus.Entry
	signal    chan struct{}
}

func NewServer() *Server {
	log.Init()

	return &Server{
		crypto: cypher.NewCrypto(),
		logger: log.NewEntry(),
		signal: make(chan struct{}),
	}
}

func (s *Server) Run() {
	s.start()
	s.logger.Println("Server started and waiting for the graceful signal...")
	<-s.signal
}

func (s *Server) start() {
	go s.watchStop()

	serverConfig := config.GetEnv().Server

	s.logger.Infof("Server is starting in port %s.", serverConfig.Port)

	s.initHttp()

	s.initDB()

	s.initPublisher()

	v1.Register(s.http.Group("/v1"), v1.Option{DB: s.db, Publisher: s.publisher, Crypto: s.crypto})

	addr := fmt.Sprintf(":%s", serverConfig.Port)
	go func() {
		if err := s.http.Start(addr); err != nil {
			s.logger.WithError(err).Fatal("Shutting down the server now")
		}
	}()
}

func (s *Server) watchStop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	s.stop()
}

func (s *Server) stop() {
	ctx, cancel := gocontext.WithTimeout(gocontext.Background(), time.Second)
	defer cancel()

	s.logger.Info("Server is stopping...")

	err := s.http.Shutdown(ctx)
	if err != nil {
		s.logger.Errorln(err)
	}

	err = s.db.Close()
	if err != nil {
		s.logger.Errorln(err)
	}

	s.publisher.Close()

	close(s.signal)
}

func (s *Server) initHttp() {
	s.http = echo.New()
	s.http.Validator = middleware.NewValidator()
	s.http.Binder = middleware.NewBinder()
	s.http.Use(middleware.Logger)
	s.http.Use(middleware.Authorize)
	s.http.Use(echomiddleware.Recover())
	s.http.Pre(echomiddleware.RemoveTrailingSlash())
	s.http.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		httpLog := context.Get(c.Request().Context(), log.HTTPKey).(*log.HTTP)
		httpLog.Level = logrus.WarnLevel

		responseErr := model.ErrorDiscover(err)
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(responseErr.StatusCode)
		} else {
			err = c.JSON(responseErr.StatusCode, responseErr)
		}
		if err != nil {
			s.http.Logger.Error(err)
		}
	}
}

func (s *Server) initDB() {
	s.db = mysql.GetConn()
}

func (s *Server) initPublisher() {
	s.publisher = kafka.NewPublisher()
}
