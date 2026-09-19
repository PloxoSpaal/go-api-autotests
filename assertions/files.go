package assertions

import (
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
