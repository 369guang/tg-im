package rpc

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/server"
	"os"
	"time"
)

const (
	BaseServicePath = "/tg_im_service"
)

type ServerOptions struct {
	Host               string
	Port               int
	MaxRecvMessageSize int
	MaxSendMessageSize int
	Name               string
	EtcdServers        []string
	CaPemFile          string
	CaKeyFile          string
}

type BaseServer struct {
	Srv          *server.Server
	Option       *ServerOptions
	etcdRegistry *serverplugin.EtcdV3RegisterPlugin
	reg          []func(src *BaseServer) error
	id           string
}

func (s *BaseServer) GetServerId() string {
	if s.id == "" {
		s.id = fmt.Sprintf("%s@%s:%d", s.Option.Name, s.Option.Host, s.Option.Port)
	}
	return s.id
}

func (s *BaseServer) Register(name string, sv interface{}) {
	s.reg = append(s.reg, func(src *BaseServer) error {
		return src.Srv.RegisterName(name, sv, "")
	})
}

func (s *BaseServer) Start() error {
	uri := fmt.Sprintf("%s:%d", s.Option.Host, s.Option.Port)

	if s.etcdRegistry != nil {
		s.etcdRegistry.ServiceAddress = fmt.Sprintf("tcp@%s", uri)

		err := s.etcdRegistry.Start()
		if err != nil {
			return err
		}

		s.Srv.Plugins.Add(s.etcdRegistry)
	}

	for _, f := range s.reg {
		if err := f(s); err != nil {
			return err
		}
	}

	fmt.Println("quic uri: ", uri)
	return s.Srv.Serve("tcp", uri)
}

func NewBaseServer(option *ServerOptions) *BaseServer {

	caCertPEM, err := os.ReadFile(option.CaPemFile)
	if err != nil {
		panic(err)
	}

	caCertPool := x509.NewCertPool()
	ok := caCertPool.AppendCertsFromPEM(caCertPEM)
	if !ok {
		panic("failed to parse root certificate")
	}

	cert, err := tls.LoadX509KeyPair(option.CaPemFile, option.CaKeyFile)
	if err != nil {
		panic(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
	}

	baseServer := &BaseServer{
		Srv: server.NewServer(server.WithTLSConfig(config)),
		id:  fmt.Sprintf("%s@%s:%d", option.Name, option.Host, option.Port),
	}
	baseServer.Option = option
	if len(option.EtcdServers) > 0 {
		baseServer.etcdRegistry = &serverplugin.EtcdV3RegisterPlugin{
			EtcdServers:    option.EtcdServers,
			BasePath:       BaseServicePath,
			Metrics:        metrics.NewRegistry(),
			UpdateInterval: time.Minute,
		}
	}

	return baseServer
}
