package main

import (
	"flag"
	"fmt"
	"github.com/369guang/tg-im/internal/cache"
	"github.com/369guang/tg-im/internal/config"
	"github.com/369guang/tg-im/internal/database"
	"github.com/369guang/tg-im/internal/logs"
	"github.com/369guang/tg-im/internal/network"
	"github.com/369guang/tg-im/server/internal/db/migrations"
	"go.uber.org/zap"
)

func main() {

	// 配置
	configFile := flag.String("c", "c", "配置文件名称（不含扩展名）")
	flag.Parse()

	fmt.Println("config file: ", *configFile)
	cfg, err := config.LoadConfig("./config", "server")
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

	// 数据库
	err = database.Init(cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DbName, cfg.Database.Port)
	if err != nil {
		logs.Logger.Error("Error loading database:", zap.Error(err))
		panic(err)

	}
	fmt.Println("init database: ", database.DB)

	// 缓存
	cache.InitCache(cfg.Cache.Host, cfg.Cache.Password, cfg.Cache.Port, cfg.Cache.DB)
	fmt.Println("init cache: success ")

	if cfg.DEBUG { // debug模式下才执行迁移
		//database.DB = database.DB.Debug()
		migrations.Migrate()
	}

	// websocket 启动服务
	//server := network.NewWebSocketServer()
	//http.HandleFunc("/ws", server.ServerHttpHandler)
	//logs.Logger.Info("Server listening on", zap.String("host", cfg.Server.Host), zap.Int("port", cfg.Server.Port))
	//fmt.Println("Server listening on", cfg.Server.Host, cfg.Server.Port)
	//err = http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port), nil)
	//if err != nil {
	//	logs.Logger.Error("Error starting server:", zap.Error(err))
	//	panic(err)
	//}

	// quic
	server, err := network.NewQuicServer(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port), cfg.Tls.CertFile, cfg.Tls.KeyFile)
	if err != nil {
		logs.Logger.Error("Error starting server:", zap.Error(err))
		panic(err)
	}

	err = server.Serve()
	select {}

	// webtransport
	//server, err := network.NewWebTransportServer(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port), cfg.Tls.CertFile, cfg.Tls.KeyFile)
	//if err != nil {
	//	logs.Logger.Error("Error starting server:", zap.Error(err))
	//	panic(err)
	//}
	//fmt.Println("init server: ", fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
	//server.Serve()
	//select {}

}
