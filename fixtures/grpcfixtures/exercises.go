package grpcfixtures

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/fixtures"
	"github.com/PloxoSpaal/go-api-autotests/resources"
)

const ExerciseFixtureKey = "grpc-exercise"

type ExerciseFixture struct {
	Data     builders.ExerciseCreate
	Request  *v1.CreateExerciseRequest
	Response *v1.GetExerciseResponse
}

func SetExerciseFixture(cfg *axiom.Config) (any, func(), error) {
	assertion := fixtures.GetAssertionFixture(cfg)
	builder := resources.GetBuilderResource(cfg.Runner)
	exercisesClient := GetExercisesClientFixture(cfg)

	courseFixture := GetCourseFixture(cfg)

	var fixture ExerciseFixture

	cfg.Setup("Create fixture exercise", func() {
		fixture.Data = builder.ExerciseCreate(
			builders.WithExerciseCreateCourseId(courseFixture.Response.GetCourse().GetId()),
		)
		fixture.Request = fixture.Data.GRPCRequest()

		var err error
		fixture.Response, err = exercisesClient.Create(cfg.Context.Raw, fixture.Request)
		assertion.NoError(err)
		assertion.GRPCCreateExerciseResponse(fixture.Response, fixture.Data)
	})

	return fixture, nil, nil
}

func GetExerciseFixture(cfg *axiom.Config) ExerciseFixture {
	return axiom.GetFixture[ExerciseFixture](cfg, ExerciseFixtureKey)
}
