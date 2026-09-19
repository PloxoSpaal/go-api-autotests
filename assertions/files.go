package assertions

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

func (a *Assertion) HTTPCreateFileResponse(actual models.FileResponse, expected builders.FileCreate) {
	a.cfg.Step("Check create HTTP file response", func() {
		a.HTTPFileCreated(actual.File, expected)
	})
}

func (a *Assertion) HTTPFileCreated(actual models.File, expected builders.FileCreate) {
	a.cfg.Step("Check created HTTP file", func() {
		a.NotEmpty(actual.ID, "file ID")
		a.Equal(actual.Filename, expected.Filename, "filename")
		a.Equal(actual.Directory, expected.Directory, "file directory")
		a.NotEmpty(actual.URL, "file URL")
	})
}

func (a *Assertion) HTTPGetFileResponse(actual, expected models.FileResponse) {
	a.cfg.Step("Check get HTTP file response", func() {
		a.HTTPFile(actual.File, expected.File)
	})
}

func (a *Assertion) HTTPFile(actual, expected models.File) {
	a.cfg.Step("Check HTTP file", func() {
		a.Equal(actual.ID, expected.ID, "file ID")
		a.Equal(actual.Filename, expected.Filename, "filename")
		a.Equal(actual.Directory, expected.Directory, "file directory")
		a.Equal(actual.URL, expected.URL, "file URL")
	})
}

func (a *Assertion) GRPCCreateFileResponse(actual *v1.GetFileResponse, expected builders.FileCreate) {
	a.cfg.Step("Check create gRPC file response", func() {
		a.NotNil(actual, "create file response")
		a.GRPCFileCreated(actual.GetFile(), expected)
	})
}

func (a *Assertion) GRPCFileCreated(actual *v1.File, expected builders.FileCreate) {
	a.cfg.Step("Check created gRPC file", func() {
		a.NotNil(actual, "file")
		a.NotEmpty(actual.GetId(), "file ID")
		a.Equal(actual.GetFilename(), expected.Filename, "filename")
		a.Equal(actual.GetDirectory(), expected.Directory, "file directory")
		a.NotEmpty(actual.GetUrl(), "file URL")
	})
}

func (a *Assertion) GRPCGetFileResponse(actual, expected *v1.GetFileResponse) {
	a.cfg.Step("Check get gRPC file response", func() {
		a.NotNil(actual, "actual get file response")
		a.NotNil(expected, "expected get file response")
		a.GRPCFile(actual.GetFile(), expected.GetFile())
	})
}

func (a *Assertion) GRPCFile(actual, expected *v1.File) {
	a.cfg.Step("Check gRPC file", func() {
		a.NotNil(actual, "actual file")
		a.NotNil(expected, "expected file")
		a.Equal(actual.GetId(), expected.GetId(), "file ID")
		a.Equal(actual.GetFilename(), expected.GetFilename(), "filename")
		a.Equal(actual.GetDirectory(), expected.GetDirectory(), "file directory")
		a.Equal(actual.GetUrl(), expected.GetUrl(), "file URL")
	})
}
