package main

import (
	"fmt"
	"github.com/space-water-bear/tg-im/pkg/net/rpc"
	"github.com/space-water-bear/tg-im/proto"
)

func main() {

	etcd := []string{"127.0.0.1:2379"}

	options := &rpc.ClientOptions{
		Host:        "10.3.21.120",
		Port:        9912,
		Name:        "gateway",
		EtcdServers: etcd,
		CaKeyFile:   "tls/server.key",
		CaPemFile:   "tls/server.pem",
	}

	cli, err := rpc.NewBaseClient(options)
	if err != nil {
		fmt.Println("new client failed:", err)
		panic(err)
	}

	defer cli.Close()

	req := &proto.AuthUser{
		Username: "zhangsan",
		Password: "lishi123",
	}
	resp := &proto.AuthResponse{}

	if err := cli.Call2("Login", req, resp); err != nil {
		fmt.Println("call failed:", err)
		panic(err)
	}

	fmt.Println(resp)
}
