package httpclients

import (
	"context"

	"github.com/PloxoSpaal/go-api-autotests/models"
	"github.com/PloxoSpaal/go-api-autotests/transport/httptransport"
	"github.com/go-resty/resty/v2"
)

type AuthenticationClient struct {
	transport *httptransport.Client
}

func NewAuthenticationClient(transport *httptransport.Client) *AuthenticationClient {
	return &AuthenticationClient{transport: transport}
}

func (c *AuthenticationClient) Login(
	ctx context.Context,
	request models.LoginRequest,
) (Response[models.LoginResponse], error) {
	return execute[models.LoginResponse](func() (*resty.Response, error) {
		return c.transport.Post(ctx, httptransport.Request{
			URL:  "/authentication/login",
			Body: request,
		})
	})
}

func (c *AuthenticationClient) Refresh(
	ctx context.Context,
	request models.RefreshRequest,
) (Response[models.LoginResponse], error) {
	return execute[models.LoginResponse](func() (*resty.Response, error) {
		return c.transport.Post(ctx, httptransport.Request{
			URL:  "/authentication/refresh",
			Body: request,
		})
	})
}
