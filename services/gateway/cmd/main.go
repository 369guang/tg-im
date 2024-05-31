package main

import (
	"flag"
	"fmt"
	"github.com/369guang/tg-im/internal/config"
	ml "github.com/369guang/tg-im/internal/middleware/logger"
	"github.com/369guang/tg-im/pkg/logger"
	"github.com/369guang/tg-im/pkg/net/rpc"
	"github.com/369guang/tg-im/pkg/net/web"
	"github.com/369guang/tg-im/proto"
	"github.com/369guang/tg-im/services/gateway/internal/routers"
	"go.uber.org/zap"
)

func main() {

	configFile := flag.String("c", "c", "配置文件名称（不含扩展名）")
	flag.Parse()

	fmt.Println("config file: ", *configFile)
	cfg, err := config.LoadConfig("./config", "gateway")
	if err != nil {
		fmt.Println("Error loading config:", err)
		panic(err)
	}
	fmt.Println("init config: ", cfg)

	// 日志
	err = logs.NewLogger(logs.Config{
		Level:      cfg.Logs.Level,
		Directory:  cfg.Logs.Directory,
		FileName:   cfg.Logs.FileName,
		ToFile:     cfg.Logs.ToFile,
		MaxSize:    cfg.Logs.MaxSize,
		MaxAge:     cfg.Logs.MaxAge,
		MaxBackups: cfg.Logs.MaxBackups,
		Compress:   cfg.Logs.Compress,
	})
	if err != nil {
		fmt.Println("Error loading logger:", err)
		panic(err)
	}
	fmt.Println("init logger: ", logs.Logger)

	logs.Logger.Info("Server started")
	logs.Logger.Info(fmt.Sprintf("config file: %v", cfg))

	// 启动http服务
	httpServer := web.NewHTTPServer()
	httpServer.UseMiddleware(ml.LoggerMiddleware)
	httpServer.RegisterRoute(routers.NewRouter)

	go func() {
		logs.Logger.Info("http server start , listen addr and port: " + fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
		fmt.Println("http server start , listen addr and port: " + fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
		err := httpServer.Start(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port), false, cfg.Tls.CertFile, cfg.Tls.KeyFile)
		if err != nil {
			logs.Logger.Error("http server start failed: ", zap.Error(err))
			fmt.Println("http server start failed: ", err)
			panic(err)
		}
	}()

	// 启动rpc服务
	rpcOptions := rpc.ServerOptions{
		Host:               cfg.Rpc.Host,
		Port:               cfg.Rpc.Port,
		Name:               cfg.Rpc.Name,
		EtcdServers:        cfg.Rpc.Etcd,
		CaPemFile:          cfg.Tls.CertFile,
		CaKeyFile:          cfg.Tls.KeyFile,
		MaxRecvMessageSize: 1024 * 1024 * 10,
		MaxSendMessageSize: 1024 * 1024 * 10,
	}
	logs.Logger.Info("rpc server start , listen addr and port: " + fmt.Sprintf("%s:%d", cfg.Rpc.Host, cfg.Rpc.Port))
	fmt.Println("rpc server start , listen addr and port: " + fmt.Sprintf("%s:%d", cfg.Rpc.Host, cfg.Rpc.Port))
	rpcServer := rpc.NewBaseServer(&rpcOptions)
	rpcServer.Register("auth", new(proto.AuthServiceImpl))
	if err := rpcServer.Start(); err != nil {
		logs.Logger.Error("rpc server start failed: ", zap.Error(err))
		fmt.Println("rpc server start failed: ", err)
		panic(err)
	}
}
