package grpcfixtures

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/fixtures"
	"github.com/PloxoSpaal/go-api-autotests/resources"
)

const FileFixtureKey = "grpc-file"

type FileFixture struct {
	Data     builders.FileCreate
	Request  *v1.CreateFileRequest
	Response *v1.GetFileResponse
}

func SetFileFixture(cfg *axiom.Config) (any, func(), error) {
	assertion := fixtures.GetAssertionFixture(cfg)
	builder := resources.GetBuilderResource(cfg.Runner)
	filesClient := GetFilesClientFixture(cfg)

	var fixture FileFixture

	cfg.Setup("Create fixture file", func() {
		fixture.Data = builder.FileCreate()
		fixture.Request = fixture.Data.GRPCRequest()

		var err error
		fixture.Response, err = filesClient.Create(cfg.Context.Raw, fixture.Request)
		assertion.NoError(err)
		assertion.GRPCCreateFileResponse(fixture.Response, fixture.Data)
	})

	return fixture, nil, nil
}

func GetFileFixture(cfg *axiom.Config) FileFixture {
	return axiom.GetFixture[FileFixture](cfg, FileFixtureKey)
}
