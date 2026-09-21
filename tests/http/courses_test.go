package http

import (
	"net/http"

	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

func (s *suite) TestListCourses() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("HTTP-COURSES-001"),
		axiom.WithCaseName("list courses"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryGetEntities),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		courseFixture := tools.Course()
		query := models.ListCoursesQuery{
			UserID: courseFixture.Response.Data.Course.CreatedByUser.ID,
		}
		response, err := tools.CoursesClient().List(cfg.Context.Raw, query)

		tools.Assertion.HTTPOK(response.StatusCode, err)
		tools.Assertion.HTTPListCoursesResponse(response.Data, courseFixture.Response.Data)
	}))
}

func (s *suite) TestCreateCourseWithEmptyTitle() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("HTTP-COURSES-002"),
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
			builders.WithCourseCreatePreviewFileID(fileFixture.Response.Data.File.ID),
			builders.WithCourseCreateCreatedByUserID(userFixture.Response.Data.User.ID),
		)
		response, err := tools.CoursesClient().Create(cfg.Context.Raw, create.HTTPRequest())

		tools.Assertion.NoError(err)
		tools.Assertion.HTTPStatus(response.StatusCode, http.StatusUnprocessableEntity)
		tools.Assertion.HTTPError(response.APIError, "title and description are required")
	}))
}
