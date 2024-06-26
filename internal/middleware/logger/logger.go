package logger

import (
	"github.com/labstack/echo/v4"
	logs "github.com/space-water-bear/tg-im/pkg/logger"
	"go.uber.org/zap"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 记录请求日志
		req := c.Request()
		res := c.Response()

		logs.Logger.Info("Incoming request",
			zap.String("method", req.Method),
			zap.String("uri", req.RequestURI),
			zap.String("IP", c.RealIP()),
		)

		if err := next(c); err != nil {
			c.Error(err)
		}

		logs.Logger.Info("Outgoing response",
			zap.Int("status", res.Status),
		)

		return nil
	}
}
