package rpc

import (
	"context"
	"fmt"
	etcdClient "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

type Cli interface {
	Call(ctx context.Context, fn string, request, reply interface{}) error
	Broadcast(fn string, request, reply interface{}) error
	Run() error
	Close() error
}

type ClientOptions struct {
	client.Option

	Host        string
	Port        int
	Name        string
	EtcdServers []string
	Selector    client.Selector
}

type BaseClient struct {
	cli     client.XClient
	options *ClientOptions
	id      string
}

func NewBaseClient(options *ClientOptions) (*BaseClient, error) {
	ret := &BaseClient{
		options: options,
	}

	var discovery client.ServiceDiscovery
	var err error

	if options.EtcdServers != nil {
		discovery, err = etcdClient.NewEtcdV3Discovery(BaseServicePath, options.Name, options.EtcdServers, false, nil)
		if err != nil {
			fmt.Println("NewEtcdV3Discovery error: ", err)
			return nil, err
		}
	} else {
		srv := fmt.Sprintf("%s@%s:%d", "tcp", options.Host, options.Port)
		discovery, err = client.NewPeer2PeerDiscovery(srv, "")
		if err != nil {
			fmt.Println("NewPeer2PeerDiscovery error: ", err)
			return nil, err
		}
	}

	if options.SerializeType == protocol.SerializeNone {
		options.SerializeType = protocol.ProtoBuffer
	}

	ret.cli = client.NewXClient(options.Name, client.Failtry, client.RoundRobin, discovery, options.Option)

	if options.Selector != nil {
		ret.cli.SetSelector(options.Selector)
	} else {
		ret.cli.SetSelector(NewRoundRobinSelector())
	}

	return ret, nil
}

func (c *BaseClient) Call2(fn string, request, reply interface{}) error {
	return c.Call(context.Background(), fn, request, reply)
}

func (c *BaseClient) Call(ctx context.Context, fn string, request, reply interface{}) error {
	return c.cli.Call(ctx, fn, request, reply)
}

func (c *BaseClient) Broadcast(fn string, request, reply interface{}) error {
	return c.cli.Broadcast(context.Background(), fn, request, reply)
}

func (c *BaseClient) Run() error {
	return nil
}

func (c *BaseClient) Close() error {
	return c.cli.Close()
}
