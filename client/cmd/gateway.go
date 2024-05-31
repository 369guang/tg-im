package main

import (
	"github.com/369guang/tg-im/pkg/net/rpc"
)

func main() {

	etcd := []string{"127.0.0.1:2379"}

	options := &rpc.ClientOptions{
		Host:        "127.0.0.1",
		Port:        9912,
		Name:        "gateway",
		EtcdServers: etcd,
	}

	cli, err := rpc.NewBaseClient(options)
	if err != nil {
		panic(err)
	}

	defer cli.Close()

	cli.Call2("Auth", nil, nil)
}
