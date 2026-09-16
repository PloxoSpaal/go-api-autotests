package httpfixtures

import (
	"net/http"

	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/clients/httpclients"
	"github.com/PloxoSpaal/go-api-autotests/models"
	"github.com/PloxoSpaal/go-api-autotests/resources"
	"github.com/stretchr/testify/require"
)

const FileFixtureKey = "http-file"

type FileFixture struct {
	Data     builders.FileCreate
	Request  models.CreateFileRequest
	Response httpclients.Response[models.FileResponse]
}

func SetFileFixture(cfg *axiom.Config) (any, func(), error) {
	builder := resources.GetBuilderResource(cfg.Runner)
	filesClient := GetFilesClientFixture(cfg)

	var fixture FileFixture

	cfg.Setup("Create fixture file", func() {
		fixture.Data = builder.FileCreate()
		fixture.Request = fixture.Data.HTTPRequest()

		var err error
		fixture.Response, err = filesClient.Create(cfg.Context.Raw, fixture.Request)
		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, fixture.Response.StatusCode)
		require.NotEmpty(cfg.T(), fixture.Response.Data.File.ID)
	})

	return fixture, nil, nil
}

func GetFileFixture(cfg *axiom.Config) FileFixture {
	return axiom.GetFixture[FileFixture](cfg, FileFixtureKey)
}
