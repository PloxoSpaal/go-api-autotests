package grpcfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/clients/grpcclients"
)

const FilesClientFixtureKey = "grpc-files-client"

func SetFilesClientFixture(cfg *axiom.Config) (any, func(), error) {
	connection := GetPrivateConnectionFixture(cfg)
	return grpcclients.NewFilesClient(connection), nil, nil
}

func GetFilesClientFixture(cfg *axiom.Config) *grpcclients.FilesClient {
	return axiom.GetFixture[*grpcclients.FilesClient](cfg, FilesClientFixtureKey)
}
