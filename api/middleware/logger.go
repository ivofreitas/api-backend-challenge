package middleware

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/sword/api-backend-challenge/context"
	"github.com/sword/api-backend-challenge/log"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	module = "api"
)

// Logger - Generates a JSON with information of request
func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasSuffix(c.Request().URL.String(), "health") {
			return next(c)
		}

		start := time.Now()

		ctx := log.InitParams(c.Request().Context())
		c.SetRequest(c.Request().WithContext(ctx))

		httpLog := context.Get(ctx, log.HTTPKey).(*log.HTTP)
		req := c.Request()
		httpLog.Module = module
		httpLog.Level = logrus.ErrorLevel
		httpLog.Request.Host = req.Host
		httpLog.Request.Route = fmt.Sprintf("[%s] %s", req.Method, req.URL.Path)
		httpLog.Request.Header = req.Header

		defer func() {
			res := c.Response()

			httpLog.Latency = float64(time.Since(start)/time.Millisecond) / 1000

			httpLog.Response.Header = res.Header()
			httpLog.Response.Status = res.Status
			httpLog.Response.RemoteIP = c.RealIP()

			entry := log.NewEntry()
			entry = entry.WithField("http", httpLog)
			entry.Log(httpLog.Level)
		}()

		return next(c)
	}
}
