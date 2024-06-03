package routers

import (
	"fmt"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/space-water-bear/tg-im/services/gateway/internal/handler"
	"github.com/spf13/viper"
)

// NewRouter 创建路由
func NewRouter(e *echo.Group) {
	fmt.Println("init router")

	baseGroup := e.Group("/api/")
	baseGroup.GET("/auth", handler.Auth)
	baseGroup.POST("/login", handler.Login)

	jwtGroup := baseGroup.Group("/api/v1")
	jwtGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(viper.GetString("jwt.secret")),
	}))

	// 用户管理
	jwtGroup.GET("/user/list", handler.UserList)
	jwtGroup.POST("/user/add", handler.UserAdd)
	jwtGroup.PUT("/user/edit", handler.UserEdit)
	jwtGroup.DELETE("/user/del", handler.UserDel)

	// 组管理
	jwtGroup.GET("/group/list", handler.GroupList)
	jwtGroup.PUT("/group/edit", handler.GroupEdit)
	jwtGroup.DELETE("/group/del", handler.GroupDel)
}
