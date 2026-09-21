package http

import (
	"net/http"

	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

func (s *suite) TestCreateExercise() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("HTTP-EXERCISES-001"),
		axiom.WithCaseName("create exercise"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryCreateEntity),
			axiom.WithMetaSeverity(axiom.SeverityBlocker),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		course := tools.Course()
		create := tools.Builder.ExerciseCreate(
			builders.WithExerciseCreateCourseId(course.Response.Data.Course.ID),
		)
		response, err := tools.ExercisesClient().Create(cfg.Context.Raw, create.HTTPRequest())

		tools.Assertion.HTTPOK(response.StatusCode, err)
		tools.Assertion.HTTPCreateExerciseResponse(response.Data, create)
	}))
}

func (s *suite) TestGetExercise() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("HTTP-EXERCISES-002"),
		axiom.WithCaseName("get exercise"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryGetEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		exercise := tools.Exercise()
		response, err := tools.ExercisesClient().Get(cfg.Context.Raw, exercise.Response.Data.Exercise.Id)

		tools.Assertion.HTTPOK(response.StatusCode, err)
		tools.Assertion.HTTPGetExerciseResponse(response.Data, exercise.Response.Data)
	}))
}

func (s *suite) TestUpdateExercise() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("HTTP-EXERCISES-003"),
		axiom.WithCaseName("update exercise"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryUpdateEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		update := tools.Builder.ExerciseUpdate()
		response, err := tools.ExercisesClient().Update(
			cfg.Context.Raw, tools.Exercise().Response.Data.Exercise.Id, update.HTTPRequest(),
		)

		tools.Assertion.HTTPOK(response.StatusCode, err)
		tools.Assertion.HTTPUpdateExerciseResponse(response.Data, update)
	}))
}

func (s *suite) TestDeleteExercise() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("HTTP-EXERCISES-004"),
		axiom.WithCaseName("delete exercise"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory(metadata.StoryDeleteEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		exerciseId := tools.Exercise().Response.Data.Exercise.Id
		response, err := tools.ExercisesClient().Delete(cfg.Context.Raw, exerciseId)

		tools.Assertion.HTTPOK(response.StatusCode, err)

		getResponse, err := tools.ExercisesClient().Get(cfg.Context.Raw, exerciseId)

		tools.Assertion.NoError(err)
		tools.Assertion.HTTPStatus(getResponse.StatusCode, http.StatusNotFound)
		tools.Assertion.HTTPError(getResponse.APIError, "Exercise not found")
	}))
}

func (s *suite) TestListExercises() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("HTTP-EXERCISES-005"),
		axiom.WithCaseName("list exercises"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryGetEntities),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		exercise := tools.Exercise()
		query := models.ListExercisesQuery{
			CourseId: exercise.Response.Data.Exercise.CourseId,
		}
		response, err := tools.ExercisesClient().List(cfg.Context.Raw, query)

		tools.Assertion.HTTPOK(response.StatusCode, err)
		tools.Assertion.HTTPListExercisesResponse(response.Data, exercise.Response.Data)
	}))
}

func (s *suite) TestCreateExerciseWithInvalidScores() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("HTTP-EXERCISES-006"),
		axiom.WithCaseName("create exercise with invalid scores"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagNegative),
			axiom.WithMetaStory(metadata.StoryValidateEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		course := tools.Course()
		create := tools.Builder.ExerciseCreate(
			builders.WithExerciseCreateCourseId(course.Response.Data.Course.ID),
			builders.WithExerciseCreateMaxScore(10),
			builders.WithExerciseCreateMinScore(100),
		)
		response, err := tools.ExercisesClient().Create(cfg.Context.Raw, create.HTTPRequest())

		tools.Assertion.NoError(err)
		tools.Assertion.HTTPStatus(response.StatusCode, http.StatusUnprocessableEntity)
		tools.Assertion.HTTPError(response.APIError, "max score should not be less than min score")
	}))
}

func (s *suite) TestGetExerciseWithUnknownID() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("HTTP-EXERCISES-007"),
		axiom.WithCaseName("get exercise with unknown id"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagNegative),
			axiom.WithMetaStory(metadata.StoryGetEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		response, err := tools.ExercisesClient().Get(cfg.Context.Raw, tools.Fake.UUID())

		tools.Assertion.NoError(err)
		tools.Assertion.HTTPStatus(response.StatusCode, http.StatusNotFound)
		tools.Assertion.HTTPError(response.APIError, "Exercise not found")
	}))
}
