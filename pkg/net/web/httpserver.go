package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HTTPServer struct {
	e      *echo.Echo
	router *echo.Group
}

// NewHTTPServer 创建实例
// eg:
// httpServer := web.NewHTTPServer()
// httpServer.RegisterRoute(router.RegisterRoute)
// httpServer.UseMiddleware(middleware.CORS())
// httpServer.Start(":8080")
func NewHTTPServer() *HTTPServer {
	e := echo.New()
	// 配置中间件
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	return &HTTPServer{
		e:      e,
		router: e.Group(""),
	}
}

// Start 启动服务
func (h *HTTPServer) Start(address string, isTls bool, caPem, caKey string) error {
	// 如果是tls服务器
	if isTls {
		return h.e.StartTLS(address, caPem, caKey)
	}
	return h.e.Start(address)
}

// Stop 停止服务
func (h *HTTPServer) Stop() {
	// 接受关闭信号 kill -9
	err := h.e.Close()

	if err != nil {
		return
	}
}

// RegisterRoute 注册路由
func (h *HTTPServer) RegisterRoute(register func(r *echo.Group)) {
	register(h.router)
}

// UseMiddleware 使用中间件
func (h *HTTPServer) UseMiddleware(middleware ...echo.MiddlewareFunc) {
	h.router.Use(middleware...)
}
