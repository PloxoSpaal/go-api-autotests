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
