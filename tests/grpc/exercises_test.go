package grpc

import (
	v1 "github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
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

func (s *suite) TestGetExercise() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("GRPC-EXERCISES-002"),
		axiom.WithCaseName("get exercise"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryGetEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		exercise := tools.Exercise()
		request := &v1.GetExerciseRequest{Id: exercise.Response.GetExercise().GetId()}
		response, err := tools.ExercisesClient().Get(cfg.Context.Raw, request)

		tools.Assertion.NoError(err)
		tools.Assertion.GRPCGetExerciseResponse(response, exercise.Response)
	}))
}

func (s *suite) TestUpdateExercise() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("GRPC-EXERCISES-003"),
		axiom.WithCaseName("update exercise"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory(metadata.StoryUpdateEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		exercise := tools.Exercise()
		request := tools.Builder.ExerciseUpdate()
		response, err := tools.ExercisesClient().Update(
			cfg.Context.Raw, request.GRPCRequest(exercise.Response.GetExercise().GetId()),
		)

		tools.Assertion.NoError(err)
		tools.Assertion.GRPCUpdateExerciseResponse(response, request)
	}))
}
