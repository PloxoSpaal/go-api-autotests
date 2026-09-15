package grpctransport

import (
	"context"

	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type bearerTokenCredentials struct {
	token string
}

func (c bearerTokenCredentials) GetRequestMetadata(
	_ context.Context,
	_ ...string,
) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + c.token,
	}, nil
}

func (bearerTokenCredentials) RequireTransportSecurity() bool {
	return false
}

func NewPublic(
	cfg *axiom.Config,
	settings config.GRPC,
) (*grpc.ClientConn, error) {
	return newConnection(cfg, settings)
}

func NewPrivate(
	cfg *axiom.Config,
	settings config.GRPC,
	token string,
) (*grpc.ClientConn, error) {
	return newConnection(
		cfg,
		settings,
		grpc.WithPerRPCCredentials(bearerTokenCredentials{token: token}),
	)
}

func newConnection(
	cfg *axiom.Config,
	settings config.GRPC,
	options ...grpc.DialOption,
) (*grpc.ClientConn, error) {
	baseOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(telemetryInterceptor(cfg, settings.Timeout)),
	}

	return grpc.NewClient(
		settings.Address,
		append(baseOptions, options...)...,
	)
}
