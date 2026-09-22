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

func (s *suite) TestCreateCourse() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("GRPC-COURSES-003"),
		axiom.WithCaseName("create course"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryCreateEntity),
			axiom.WithMetaSeverity(axiom.SeverityBlocker),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		userFixture := tools.User()
		fileFixture := tools.File()
		create := tools.Builder.CourseCreate(
			builders.WithCourseCreatePreviewFileID(fileFixture.Response.GetFile().GetId()),
			builders.WithCourseCreateCreatedByUserID(userFixture.Response.GetUser().GetId()),
		)
		response, err := tools.CoursesClient().Create(cfg.Context.Raw, create.GRPCRequest())

		tools.Assertion.NoError(err)
		tools.Assertion.GRPCCreateCourseResponse(response, create)
	}))
}

func (s *suite) TestUpdateCourseWithEmptyDescription() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("GRPC-COURSES-004"),
		axiom.WithCaseName("update course with empty description"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagNegative),
			axiom.WithMetaStory(metadata.StoryValidateEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		update := tools.Builder.CourseUpdate(builders.WithCourseUpdateDescription(""))
		response, err := tools.CoursesClient().Update(
			cfg.Context.Raw, update.GRPCRequest(tools.Course().Response.GetCourse().GetId()),
		)

		tools.Assertion.GRPCError(err, codes.InvalidArgument, "description is required")
		tools.Assertion.Nil(response, "update course with empty description response")
	}))
}

func (s *suite) TestGetCourse() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("GRPC-COURSES-005"),
		axiom.WithCaseName("get course"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryGetEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		course := tools.Course()
		request := &v1.GetCourseRequest{Id: course.Response.GetCourse().GetId()}
		response, err := tools.CoursesClient().Get(cfg.Context.Raw, request)

		tools.Assertion.NoError(err)
		tools.Assertion.GRPCGetCourseResponse(response, course.Response)
	}))
}
