package httpclients

import (
	"context"
	"net/url"

	"github.com/PloxoSpaal/go-api-autotests/models"
	"github.com/PloxoSpaal/go-api-autotests/transport/httptransport"
	"github.com/go-resty/resty/v2"
)

type UsersClient struct {
	transport *httptransport.Client
}

func NewUsersClient(transport *httptransport.Client) *UsersClient {
	return &UsersClient{transport: transport}
}

func (c *UsersClient) Create(
	ctx context.Context,
	request models.CreateUserRequest,
) (Response[models.UserResponse], error) {
	return execute[models.UserResponse](func() (*resty.Response, error) {
		return c.transport.Post(ctx, httptransport.Request{
			URL:  "/users",
			Body: request,
		})
	})
}

func (c *UsersClient) GetMe(ctx context.Context) (Response[models.UserResponse], error) {
	return execute[models.UserResponse](func() (*resty.Response, error) {
		return c.transport.Get(ctx, httptransport.Request{
			URL: "/users/me",
		})
	})
}

func (c *UsersClient) Get(
	ctx context.Context,
	userID string,
) (Response[models.UserResponse], error) {
	return execute[models.UserResponse](func() (*resty.Response, error) {
		return c.transport.Get(ctx, httptransport.Request{
			URL: "/users/" + url.PathEscape(userID),
		})
	})
}

func (c *UsersClient) Update(
	ctx context.Context,
	userID string,
	request models.UpdateUserRequest,
) (Response[models.UserResponse], error) {
	return execute[models.UserResponse](func() (*resty.Response, error) {
		return c.transport.Patch(ctx, httptransport.Request{
			URL:  "/users/" + url.PathEscape(userID),
			Body: request,
		})
	})
}

func (c *UsersClient) Delete(
	ctx context.Context,
	userID string,
) (Response[models.Empty], error) {
	return execute[models.Empty](func() (*resty.Response, error) {
		return c.transport.Delete(ctx, httptransport.Request{
			URL: "/users/" + url.PathEscape(userID),
		})
	})
}
