package httpfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/clients/httpclients"
)

const FilesClientFixtureKey = "http-files-client"

func SetFilesClientFixture(cfg *axiom.Config) (any, func(), error) {
	transport := GetPrivateTransportFixture(cfg)
	return httpclients.NewFilesClient(transport), nil, nil
}

func GetFilesClientFixture(cfg *axiom.Config) *httpclients.FilesClient {
	return axiom.GetFixture[*httpclients.FilesClient](cfg, FilesClientFixtureKey)
}
