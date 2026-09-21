package http

import (
	"net/http"

	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
)

func (s *suite) TestDeleteFile() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("HTTP-FILES-001"),
		axiom.WithCaseName("delete file"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory(metadata.StoryDeleteEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		fileFixture := tools.File()
		deleteResponse, err := tools.FilesClient().Delete(
			cfg.Context.Raw,
			fileFixture.Response.Data.File.ID,
		)
		tools.Assertion.HTTPOK(deleteResponse.StatusCode, err)

		getResponse, err := tools.FilesClient().Get(
			cfg.Context.Raw,
			fileFixture.Response.Data.File.ID,
		)
		tools.Assertion.NoError(err)
		tools.Assertion.HTTPStatus(getResponse.StatusCode, http.StatusNotFound)
		tools.Assertion.HTTPError(getResponse.APIError, "File not found")
	}))
}

func (s *suite) TestCreateFile() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("HTTP-FILES-002"),
		axiom.WithCaseName("create file"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory(metadata.StoryCreateEntity),
			axiom.WithMetaSeverity(axiom.SeverityBlocker),
			axiom.WithMetaTag(metadata.TagSmoke),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		fileData := tools.Builder.FileCreate()
		createResponse, err := tools.FilesClient().Create(
			cfg.Context.Raw,
			fileData.HTTPRequest(),
		)

		tools.Assertion.HTTPOK(createResponse.StatusCode, err)
		tools.Assertion.HTTPCreateFileResponse(createResponse.Data, fileData)
	}))
}

func (s *suite) TestGetFile() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("HTTP-FILES-003"),
		axiom.WithCaseName("get file"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory(metadata.StoryGetEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
			axiom.WithMetaTag(metadata.TagSmoke),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		fileData := tools.File()
		getResponse, err := tools.FilesClient().Get(
			cfg.Context.Raw,
			fileData.Response.Data.File.ID,
		)

		tools.Assertion.HTTPOK(getResponse.StatusCode, err)
		tools.Assertion.HTTPGetFileResponse(getResponse.Data, fileData.Response.Data)
	}))
}
