package rpc

import (
	"context"
	"fmt"
	"github.com/369guang/tg-im/proto"
)

type GatewayImpl struct {
}

func (g *GatewayImpl) Login(ctx context.Context, args *proto.AuthRequest, reply *proto.AuthResponse) error {

	fmt.Println("Login 输入:", args, "返回", reply)

	*reply = proto.AuthResponse{
		Token: "token",
	}

	return nil
}
