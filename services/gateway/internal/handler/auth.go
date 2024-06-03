package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/space-water-bear/tg-im/internal/errno"
	"github.com/space-water-bear/tg-im/pkg/net/web"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Auth 认证处理
func Auth(ctx echo.Context) error {

	return web.JSON(ctx, errno.Success, nil)
}

// Login POST 登录处理
func Login(ctx echo.Context) error {
	// 获取参数
	user := new(LoginData)
	if err := ctx.Bind(user); err != nil {
		fmt.Println("bind failed:", err)
		return web.JSON(ctx, errno.ErrCodeParam, nil)
	}

	fmt.Println("Login 输入:", user)

	return web.JSON(ctx, errno.Success, nil)
}
