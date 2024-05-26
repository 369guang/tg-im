package main

import (
	"flag"
	"fmt"
	"github.com/369guang/tg-im/internal/config"
	"github.com/369guang/tg-im/internal/network"
	"log"
	"time"
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
	client, err := network.NewQuicClient("localhost:9991", "localhost")
	//client, err := network.NewWebTransportClient(fmt.Sprintf("https://%s:%d/", "127.0.0.1", cfg.Server.Port))
	if err != nil {
		fmt.Println("Error loading client:", err)
		panic(err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatalf("Failed to close client: %v", err)
		}
	}()

	fmt.Println("init client: ", client)

	message := "Hello, world!"
	err = client.Send(message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		panic(err)
	}

	fmt.Printf("Sent message: %s\n", message)

	response, err := client.Receive()
	if err != nil {
		fmt.Println("Error receiving message:", err)
		panic(err)
	}

	fmt.Println("Received message:", response)

	time.Sleep(5 * time.Second)
}
