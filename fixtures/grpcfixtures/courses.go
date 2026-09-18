package grpcfixtures

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/resources"
	"github.com/stretchr/testify/require"
)

const CourseFixtureKey = "grpc-course"

type CourseFixture struct {
	Data     builders.CourseCreate
	Request  *v1.CreateCourseRequest
	Response *v1.GetCourseResponse
}

func SetCourseFixture(cfg *axiom.Config) (any, func(), error) {
	builder := resources.GetBuilderResource(cfg.Runner)
	coursesClient := GetCoursesClientFixture(cfg)

	userFixture := GetUserFixture(cfg)
	fileFixture := GetFileFixture(cfg)

	var fixture CourseFixture

	cfg.Setup("Create fixture course", func() {
		fixture.Data = builder.CourseCreate(
			builders.WithCourseCreatePreviewFileID(fileFixture.Response.GetFile().GetId()),
			builders.WithCourseCreateCreatedByUserID(userFixture.Response.GetUser().GetId()),
		)
		fixture.Request = fixture.Data.GRPCRequest()

		var err error
		fixture.Response, err = coursesClient.Create(cfg.Context.Raw, fixture.Request)
		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), fixture.Response)
		require.NotNil(cfg.T(), fixture.Response.GetCourse())
		require.NotEmpty(cfg.T(), fixture.Response.GetCourse().GetId())
	})

	return fixture, nil, nil
}

func GetCourseFixture(cfg *axiom.Config) CourseFixture {
	return axiom.GetFixture[CourseFixture](cfg, CourseFixtureKey)
}
