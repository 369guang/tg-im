package handler

import (
	"github.com/369guang/tg-im/internal/errno"
	"github.com/369guang/tg-im/pkg/net/web"
	"github.com/labstack/echo/v4"
)

// Auth 认证处理
func Auth(ctx echo.Context) error {

	return web.JSON(ctx, errno.Success, nil)
}
