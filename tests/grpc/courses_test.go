package grpc

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
	"google.golang.org/grpc/codes"
)

func (s *suite) TestListCourses() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("GRPC-COURSES-001"),
		axiom.WithCaseName("list courses"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryGetEntities),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		courseFixture := tools.Course()
		request := &v1.ListCoursesRequest{
			UserId: courseFixture.Response.GetCourse().GetCreatedByUser().GetId(),
		}
		response, err := tools.CoursesClient().List(cfg.Context.Raw, request)

		tools.Assertion.NoError(err)
		tools.Assertion.GRPCListCoursesResponse(response, courseFixture.Response)
	}))
}

func (s *suite) TestCreateCourseWithEmptyTitle() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("GRPC-COURSES-002"),
		axiom.WithCaseName("create course with empty title"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagNegative),
			axiom.WithMetaStory(metadata.StoryValidateEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		userFixture := tools.User()
		fileFixture := tools.File()
		create := tools.Builder.CourseCreate(
			builders.WithCourseCreateTitle(""),
			builders.WithCourseCreatePreviewFileID(fileFixture.Response.GetFile().GetId()),
			builders.WithCourseCreateCreatedByUserID(userFixture.Response.GetUser().GetId()),
		)
		response, err := tools.CoursesClient().Create(cfg.Context.Raw, create.GRPCRequest())

		tools.Assertion.GRPCError(err, codes.InvalidArgument, "title and description are required")
		tools.Assertion.Nil(response, "gRPC response")
	}))
}
