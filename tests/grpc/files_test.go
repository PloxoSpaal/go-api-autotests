package grpc

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
	"google.golang.org/grpc/codes"
)

func (s *suite) TestDeleteFile() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("GRPC-FILES-001"),
		axiom.WithCaseName("delete file"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory(metadata.StoryDeleteEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		fileFixture := tools.File()

		deleteRequest := &v1.DeleteFileRequest{Id: fileFixture.Response.GetFile().GetId()}
		deleteResponse, err := tools.FilesClient().Delete(cfg.Context.Raw, deleteRequest)
		tools.Assertion.NoError(err)
		tools.Assertion.NotNil(deleteResponse, "delete file response")

		getRequest := &v1.GetFileRequest{Id: fileFixture.Response.GetFile().GetId()}
		getResponse, err := tools.FilesClient().Get(cfg.Context.Raw, getRequest)
		tools.Assertion.GRPCError(err, codes.NotFound, "File not found")
		tools.Assertion.Nil(getResponse, "gRPC response")
	}))
}
