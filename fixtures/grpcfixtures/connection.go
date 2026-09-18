package grpcfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/resources"
	"github.com/PloxoSpaal/go-api-autotests/transport/grpctransport"
	"google.golang.org/grpc"
)

const PublicConnectionFixtureKey = "grpc-public-connection"

func SetPublicConnectionFixture(cfg *axiom.Config) (any, func(), error) {
	settings := resources.GetConfigResource(cfg.Runner)

	connection, err := grpctransport.NewPublic(cfg, settings.GRPC)
	if err != nil {
		return nil, nil, err
	}

	return connection, func() {
		_ = connection.Close()
	}, nil
}

func GetPublicConnectionFixture(cfg *axiom.Config) *grpc.ClientConn {
	return axiom.GetFixture[*grpc.ClientConn](cfg, PublicConnectionFixtureKey)
}
