package grpcclients

import (
	"context"

	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"google.golang.org/grpc"
)

type AuthenticationClient struct {
	client v1.AuthenticationServiceClient
}

func NewAuthenticationClient(connection grpc.ClientConnInterface) *AuthenticationClient {
	return &AuthenticationClient{client: v1.NewAuthenticationServiceClient(connection)}
}

func (c *AuthenticationClient) Login(
	ctx context.Context,
	request *v1.LoginRequest,
) (*v1.LoginResponse, error) {
	return c.client.Login(ctx, request)
}

func (c *AuthenticationClient) Refresh(
	ctx context.Context,
	request *v1.RefreshRequest,
) (*v1.LoginResponse, error) {
	return c.client.Refresh(ctx, request)
}
