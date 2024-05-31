package routers

import (
	"github.com/369guang/tg-im/services/gateway/internal/handler"
	"github.com/labstack/echo/v4"
)

// NewRouter 创建路由
func NewRouter(e *echo.Group) {
	e.GET("/auth", handler.Auth)
}
