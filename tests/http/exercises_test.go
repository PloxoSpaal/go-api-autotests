package http

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
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
