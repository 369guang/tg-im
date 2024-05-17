package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"tg-im/core/config"
	"tg-im/core/database"
	"tg-im/core/logs"
	"tg-im/server/internal/db/migrations"
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

	logs.Logger.Info("Server started")
	logs.Logger.Info(fmt.Sprintf("config file: %v", cfg))

	// 数据库
	err = database.Init(cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DbName, cfg.Database.Port)
	if err != nil {
		logs.Logger.Error("Error loading database:", zap.Error(err))
		panic(err)

	}

	if cfg.DEBUG {
		//database.DB = database.DB.Debug()
		migrations.Migrate()
	}

}
