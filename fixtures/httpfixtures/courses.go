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

const CourseFixtureKey = "http-course"

type CourseFixture struct {
	Data     builders.CourseCreate
	Request  models.CreateCourseRequest
	Response httpclients.Response[models.CourseResponse]
}

func SetCourseFixture(cfg *axiom.Config) (any, func(), error) {
	builder := resources.GetBuilderResource(cfg.Runner)
	coursesClient := GetCoursesClientFixture(cfg)

	userFixture := GetUserFixture(cfg)
	fileFixture := GetFileFixture(cfg)

	var fixture CourseFixture

	cfg.Setup("Create fixture course", func() {
		fixture.Data = builder.CourseCreate(
			builders.WithCourseCreatePreviewFileID(fileFixture.Response.Data.File.ID),
			builders.WithCourseCreateCreatedByUserID(userFixture.Response.Data.User.ID),
		)
		fixture.Request = fixture.Data.HTTPRequest()

		var err error
		fixture.Response, err = coursesClient.Create(cfg.Context.Raw, fixture.Request)
		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, fixture.Response.StatusCode)
		require.NotEmpty(cfg.T(), fixture.Response.Data.Course.ID)
	})

	return fixture, nil, nil
}

func GetCourseFixture(cfg *axiom.Config) CourseFixture {
	return axiom.GetFixture[CourseFixture](cfg, CourseFixtureKey)
}
