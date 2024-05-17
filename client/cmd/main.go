package main

import (
	"fmt"
	"tg-im/core/config"
)

func main() {

	cfg, err := config.LoadConfig("./config", "server")
	if err != nil {
		fmt.Println("Error loading config:", err)
		panic(err)
	}

	fmt.Println("Server:", cfg.Server)
	fmt.Println("Database:", cfg.Database)
	fmt.Println("JWT:", cfg.JWT)
	fmt.Println("Logs:", cfg.Logs)
}
