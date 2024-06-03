package rpc

import (
	"context"
	"fmt"
	"github.com/space-water-bear/tg-im/proto"
)

type GatewayImpl struct {
}

func (g *GatewayImpl) Login(ctx context.Context, requests *proto.AuthUser, response *proto.AuthResponse) error {

	fmt.Println("Login 输入:", requests, "返回", response)

	response.Token = "123456"

	return nil
}
