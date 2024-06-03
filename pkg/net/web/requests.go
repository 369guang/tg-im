package web

import (
	"github.com/labstack/echo/v4"
	"github.com/space-water-bear/tg-im/internal/errno"
	"net/http"
)

// JSON 返回值JSON处理
func JSON(c echo.Context, code int, data interface{}) error {

	resp := map[string]interface{}{
		"code": code,
		"msg":  errno.ErrMsg[code],
		"data": data,
	}

	return c.JSON(http.StatusOK, resp)
}
