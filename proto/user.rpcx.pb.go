// Code generated by protoc-gen-rpcx. DO NOT EDIT.
// versions:
// - protoc-gen-rpcx v0.3.0
// - protoc          v3.20.3
// source: user.proto

package proto

import (
	context "context"
	client "github.com/smallnest/rpcx/client"
	protocol "github.com/smallnest/rpcx/protocol"
	server "github.com/smallnest/rpcx/server"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = context.TODO
var _ = server.NewServer
var _ = client.NewClient
var _ = protocol.NewMessage

// ================== interface skeleton ===================
type UserServiceAble interface {
	// UserServiceAble can be used for interface verification.

	// GetUser is server rpc method as defined
	GetUser(ctx context.Context, args *GetUserRequest, reply *GetUserResponse) (err error)

	// SearchUser is server rpc method as defined
	SearchUser(ctx context.Context, args *SearchUserRequest, reply *SearchUserResponse) (err error)

	// UpdateUser is server rpc method as defined
	UpdateUser(ctx context.Context, args *UpdateUserRequest, reply *UpdateUserResponse) (err error)

	// AddFriend is server rpc method as defined
	AddFriend(ctx context.Context, args *FriendRequest, reply *FriendResponse) (err error)

	// RemoveFriend is server rpc method as defined
	RemoveFriend(ctx context.Context, args *FriendRequest, reply *emptypb.Empty) (err error)

	// GetFriendList is server rpc method as defined
	GetFriendList(ctx context.Context, args *FriendListRequest, reply *FriendListResponse) (err error)

	// GetFriendInfo is server rpc method as defined
	GetFriendInfo(ctx context.Context, args *FriendRequest, reply *FriendResponse) (err error)
}

// ================== server skeleton ===================
type UserServiceImpl struct{}

// ServeForUserService starts a server only registers one service.
// You can register more services and only start one server.
// It blocks until the application exits.
func ServeForUserService(addr string) error {
	s := server.NewServer()
	s.RegisterName("UserService", new(UserServiceImpl), "")
	return s.Serve("tcp", addr)
}

// GetUser is server rpc method as defined
func (s *UserServiceImpl) GetUser(ctx context.Context, args *GetUserRequest, reply *GetUserResponse) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = GetUserResponse{}

	return nil
}

// SearchUser is server rpc method as defined
func (s *UserServiceImpl) SearchUser(ctx context.Context, args *SearchUserRequest, reply *SearchUserResponse) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = SearchUserResponse{}

	return nil
}

// UpdateUser is server rpc method as defined
func (s *UserServiceImpl) UpdateUser(ctx context.Context, args *UpdateUserRequest, reply *UpdateUserResponse) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = UpdateUserResponse{}

	return nil
}

// AddFriend is server rpc method as defined
func (s *UserServiceImpl) AddFriend(ctx context.Context, args *FriendRequest, reply *FriendResponse) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = FriendResponse{}

	return nil
}

// RemoveFriend is server rpc method as defined
func (s *UserServiceImpl) RemoveFriend(ctx context.Context, args *FriendRequest, reply *emptypb.Empty) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = emptypb.Empty{}

	return nil
}

// GetFriendList is server rpc method as defined
func (s *UserServiceImpl) GetFriendList(ctx context.Context, args *FriendListRequest, reply *FriendListResponse) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = FriendListResponse{}

	return nil
}

// GetFriendInfo is server rpc method as defined
func (s *UserServiceImpl) GetFriendInfo(ctx context.Context, args *FriendRequest, reply *FriendResponse) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = FriendResponse{}

	return nil
}

// ================== client stub ===================
// UserService is a client wrapped XClient.
type UserServiceClient struct {
	xclient client.XClient
}

// NewUserServiceClient wraps a XClient as UserServiceClient.
// You can pass a shared XClient object created by NewXClientForUserService.
func NewUserServiceClient(xclient client.XClient) *UserServiceClient {
	return &UserServiceClient{xclient: xclient}
}

// NewXClientForUserService creates a XClient.
// You can configure this client with more options such as etcd registry, serialize type, select algorithm and fail mode.
func NewXClientForUserService(addr string) (client.XClient, error) {
	d, err := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	if err != nil {
		return nil, err
	}

	opt := client.DefaultOption
	opt.SerializeType = protocol.ProtoBuffer

	xclient := client.NewXClient("UserService", client.Failtry, client.RoundRobin, d, opt)

	return xclient, nil
}

// GetUser is client rpc method as defined
func (c *UserServiceClient) GetUser(ctx context.Context, args *GetUserRequest) (reply *GetUserResponse, err error) {
	reply = &GetUserResponse{}
	err = c.xclient.Call(ctx, "GetUser", args, reply)
	return reply, err
}

// SearchUser is client rpc method as defined
func (c *UserServiceClient) SearchUser(ctx context.Context, args *SearchUserRequest) (reply *SearchUserResponse, err error) {
	reply = &SearchUserResponse{}
	err = c.xclient.Call(ctx, "SearchUser", args, reply)
	return reply, err
}

// UpdateUser is client rpc method as defined
func (c *UserServiceClient) UpdateUser(ctx context.Context, args *UpdateUserRequest) (reply *UpdateUserResponse, err error) {
	reply = &UpdateUserResponse{}
	err = c.xclient.Call(ctx, "UpdateUser", args, reply)
	return reply, err
}

// AddFriend is client rpc method as defined
func (c *UserServiceClient) AddFriend(ctx context.Context, args *FriendRequest) (reply *FriendResponse, err error) {
	reply = &FriendResponse{}
	err = c.xclient.Call(ctx, "AddFriend", args, reply)
	return reply, err
}

// RemoveFriend is client rpc method as defined
func (c *UserServiceClient) RemoveFriend(ctx context.Context, args *FriendRequest) (reply *emptypb.Empty, err error) {
	reply = &emptypb.Empty{}
	err = c.xclient.Call(ctx, "RemoveFriend", args, reply)
	return reply, err
}

// GetFriendList is client rpc method as defined
func (c *UserServiceClient) GetFriendList(ctx context.Context, args *FriendListRequest) (reply *FriendListResponse, err error) {
	reply = &FriendListResponse{}
	err = c.xclient.Call(ctx, "GetFriendList", args, reply)
	return reply, err
}

// GetFriendInfo is client rpc method as defined
func (c *UserServiceClient) GetFriendInfo(ctx context.Context, args *FriendRequest) (reply *FriendResponse, err error) {
	reply = &FriendResponse{}
	err = c.xclient.Call(ctx, "GetFriendInfo", args, reply)
	return reply, err
}

// ================== oneclient stub ===================
// UserServiceOneClient is a client wrapped oneClient.
type UserServiceOneClient struct {
	serviceName string
	oneclient   *client.OneClient
}

// NewUserServiceOneClient wraps a OneClient as UserServiceOneClient.
// You can pass a shared OneClient object created by NewOneClientForUserService.
func NewUserServiceOneClient(oneclient *client.OneClient) *UserServiceOneClient {
	return &UserServiceOneClient{
		serviceName: "UserService",
		oneclient:   oneclient,
	}
}

// ======================================================

// GetUser is client rpc method as defined
func (c *UserServiceOneClient) GetUser(ctx context.Context, args *GetUserRequest) (reply *GetUserResponse, err error) {
	reply = &GetUserResponse{}
	err = c.oneclient.Call(ctx, c.serviceName, "GetUser", args, reply)
	return reply, err
}

// SearchUser is client rpc method as defined
func (c *UserServiceOneClient) SearchUser(ctx context.Context, args *SearchUserRequest) (reply *SearchUserResponse, err error) {
	reply = &SearchUserResponse{}
	err = c.oneclient.Call(ctx, c.serviceName, "SearchUser", args, reply)
	return reply, err
}

// UpdateUser is client rpc method as defined
func (c *UserServiceOneClient) UpdateUser(ctx context.Context, args *UpdateUserRequest) (reply *UpdateUserResponse, err error) {
	reply = &UpdateUserResponse{}
	err = c.oneclient.Call(ctx, c.serviceName, "UpdateUser", args, reply)
	return reply, err
}

// AddFriend is client rpc method as defined
func (c *UserServiceOneClient) AddFriend(ctx context.Context, args *FriendRequest) (reply *FriendResponse, err error) {
	reply = &FriendResponse{}
	err = c.oneclient.Call(ctx, c.serviceName, "AddFriend", args, reply)
	return reply, err
}

// RemoveFriend is client rpc method as defined
func (c *UserServiceOneClient) RemoveFriend(ctx context.Context, args *FriendRequest) (reply *emptypb.Empty, err error) {
	reply = &emptypb.Empty{}
	err = c.oneclient.Call(ctx, c.serviceName, "RemoveFriend", args, reply)
	return reply, err
}

// GetFriendList is client rpc method as defined
func (c *UserServiceOneClient) GetFriendList(ctx context.Context, args *FriendListRequest) (reply *FriendListResponse, err error) {
	reply = &FriendListResponse{}
	err = c.oneclient.Call(ctx, c.serviceName, "GetFriendList", args, reply)
	return reply, err
}

// GetFriendInfo is client rpc method as defined
func (c *UserServiceOneClient) GetFriendInfo(ctx context.Context, args *FriendRequest) (reply *FriendResponse, err error) {
	reply = &FriendResponse{}
	err = c.oneclient.Call(ctx, c.serviceName, "GetFriendInfo", args, reply)
	return reply, err
}
