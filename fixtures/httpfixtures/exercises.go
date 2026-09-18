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

const ExerciseFixtureKey = "http-exercise"

type ExerciseFixture struct {
	Data     builders.ExerciseCreate
	Request  models.CreateExerciseRequest
	Response httpclients.Response[models.ExerciseResponse]
}

func SetExerciseFixture(cfg *axiom.Config) (any, func(), error) {
	builder := resources.GetBuilderResource(cfg.Runner)
	exercisesClient := GetExercisesClientFixture(cfg)
	courseFixture := GetCourseFixture(cfg)

	var fixture ExerciseFixture

	cfg.Setup("Create exercise course", func() {
		fixture.Data = builder.ExerciseCreate(
			builders.WithExerciseCreateCourseId(courseFixture.Response.Data.Course.ID),
		)
		fixture.Request = fixture.Data.HTTPRequest()

		var err error
		fixture.Response, err = exercisesClient.Create(cfg.Context.Raw, fixture.Request)
		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, fixture.Response.StatusCode)
		require.NotEmpty(cfg.T(), fixture.Response.Data.Exercise.Id)
	})

	return fixture, nil, nil
}

func GetExerciseFixture(cfg *axiom.Config) ExerciseFixture {
	return axiom.GetFixture[ExerciseFixture](cfg, ExerciseFixtureKey)
}
