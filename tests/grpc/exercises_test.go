package grpc

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
)

func (s *suite) TestCreateExercise() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("GRPC-EXERCISES-001"),
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
			builders.WithExerciseCreateCourseId(course.Response.GetCourse().GetId()),
		)
		response, err := tools.ExercisesClient().Create(cfg.Context.Raw, create.GRPCRequest())

		tools.Assertion.NoError(err)
		tools.Assertion.GRPCCreateExerciseResponse(response, create)
	}))
}
