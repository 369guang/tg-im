// Code generated by protoc-gen-rpcx. DO NOT EDIT.
// versions:
// - protoc-gen-rpcx v0.3.0
// - protoc          v3.20.3
// source: auth.proto

package proto

import (
	context "context"
	client "github.com/smallnest/rpcx/client"
	protocol "github.com/smallnest/rpcx/protocol"
	server "github.com/smallnest/rpcx/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = context.TODO
var _ = server.NewServer
var _ = client.NewClient
var _ = protocol.NewMessage

// ================== interface skeleton ===================
type AuthServiceAble interface {
	// AuthServiceAble can be used for interface verification.

	// Login is server rpc method as defined
	Login(ctx context.Context, args *AuthUser, reply *AuthResponse) (err error)
}

// ================== server skeleton ===================
type AuthServiceImpl struct{}

// ServeForAuthService starts a server only registers one service.
// You can register more services and only start one server.
// It blocks until the application exits.
func ServeForAuthService(addr string) error {
	s := server.NewServer()
	s.RegisterName("AuthService", new(AuthServiceImpl), "")
	return s.Serve("tcp", addr)
}

// Login is server rpc method as defined
func (s *AuthServiceImpl) Login(ctx context.Context, args *AuthUser, reply *AuthResponse) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = AuthResponse{}

	return nil
}

// ================== client stub ===================
// AuthService is a client wrapped XClient.
type AuthServiceClient struct {
	xclient client.XClient
}

// NewAuthServiceClient wraps a XClient as AuthServiceClient.
// You can pass a shared XClient object created by NewXClientForAuthService.
func NewAuthServiceClient(xclient client.XClient) *AuthServiceClient {
	return &AuthServiceClient{xclient: xclient}
}

// NewXClientForAuthService creates a XClient.
// You can configure this client with more options such as etcd registry, serialize type, select algorithm and fail mode.
func NewXClientForAuthService(addr string) (client.XClient, error) {
	d, err := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	if err != nil {
		return nil, err
	}

	opt := client.DefaultOption
	opt.SerializeType = protocol.ProtoBuffer

	xclient := client.NewXClient("AuthService", client.Failtry, client.RoundRobin, d, opt)

	return xclient, nil
}

// Login is client rpc method as defined
func (c *AuthServiceClient) Login(ctx context.Context, args *AuthUser) (reply *AuthResponse, err error) {
	reply = &AuthResponse{}
	err = c.xclient.Call(ctx, "Login", args, reply)
	return reply, err
}

// ================== oneclient stub ===================
// AuthServiceOneClient is a client wrapped oneClient.
type AuthServiceOneClient struct {
	serviceName string
	oneclient   *client.OneClient
}

// NewAuthServiceOneClient wraps a OneClient as AuthServiceOneClient.
// You can pass a shared OneClient object created by NewOneClientForAuthService.
func NewAuthServiceOneClient(oneclient *client.OneClient) *AuthServiceOneClient {
	return &AuthServiceOneClient{
		serviceName: "AuthService",
		oneclient:   oneclient,
	}
}

// ======================================================

// Login is client rpc method as defined
func (c *AuthServiceOneClient) Login(ctx context.Context, args *AuthUser) (reply *AuthResponse, err error) {
	reply = &AuthResponse{}
	err = c.oneclient.Call(ctx, c.serviceName, "Login", args, reply)
	return reply, err
}
