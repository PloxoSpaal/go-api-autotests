package grpcclients

import (
	"context"

	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"google.golang.org/grpc"
)

type UsersClient struct {
	client v1.UsersServiceClient
}

func NewUsersClient(connection grpc.ClientConnInterface) *UsersClient {
	return &UsersClient{client: v1.NewUsersServiceClient(connection)}
}

func (c *UsersClient) Create(
	ctx context.Context,
	request *v1.CreateUserRequest,
) (*v1.GetUserResponse, error) {
	return c.client.CreateUser(ctx, request)
}

func (c *UsersClient) GetMe(ctx context.Context) (*v1.GetUserResponse, error) {
	return c.client.GetMe(ctx, &v1.Empty{})
}

func (c *UsersClient) Get(
	ctx context.Context,
	request *v1.GetUserRequest,
) (*v1.GetUserResponse, error) {
	return c.client.GetUser(ctx, request)
}

func (c *UsersClient) Update(
	ctx context.Context,
	request *v1.UpdateUserRequest,
) (*v1.GetUserResponse, error) {
	return c.client.UpdateUser(ctx, request)
}

func (c *UsersClient) Delete(
	ctx context.Context,
	request *v1.GetUserRequest,
) (*v1.Empty, error) {
	return c.client.DeleteUser(ctx, request)
}
