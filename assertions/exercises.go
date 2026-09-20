package assertions

import (
	v1 "github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

func (a *Assertion) HTTPCreateExerciseResponse(
	actual models.ExerciseResponse,
	expected builders.ExerciseCreate,
) {
	message := "Check create HTTP exercise response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.HTTPExerciseCreated(actual.Exercise, expected)
	})
}

func (a *Assertion) HTTPExerciseCreated(actual models.Exercise, expected builders.ExerciseCreate) {
	message := "Check created HTTP exercise"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotEmpty(actual.Id, "exercise ID")
		a.NotNil(actual.MaxScore, "exercise max score")
		a.NotNil(actual.MinScore, "exercise min score")
		a.NotNil(actual.EstimatedTime, "exercise estimated time")
		a.Equal(actual.Title, expected.Title, "exercise title")
		a.Equal(actual.CourseId, expected.CourseId, "exercise course id")
		a.Equal(*actual.MaxScore, expected.MaxScore, "exercise max score")
		a.Equal(*actual.MinScore, expected.MinScore, "exercise min score")
		a.Equal(actual.OrderIndex, expected.OrderIndex, "exercise order index")
		a.Equal(actual.Description, expected.Description, "exercise description")
		a.Equal(*actual.EstimatedTime, expected.EstimatedTime, "exercise estimated time")
	})
}

func (a *Assertion) HTTPGetExerciseResponse(
	actual models.ExerciseResponse,
	expected models.ExerciseResponse,
) {
	message := "Check get HTTP exercise response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.HTTPExercise(actual.Exercise, expected.Exercise)
	})
}

func (a *Assertion) HTTPExercise(actual, expected models.Exercise) {
	message := "Check HTTP exercise"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(expected.MaxScore, "expected exercise max score")
		a.NotNil(actual.MaxScore, "actual exercise max score")
		a.NotNil(expected.MinScore, "expected exercise min score")
		a.NotNil(actual.MinScore, "actual exercise min score")
		a.NotNil(expected.EstimatedTime, "expected exercise estimated time")
		a.NotNil(actual.EstimatedTime, "actual exercise estimated time")
		a.Equal(actual.Id, expected.Id, "exercise Id")
		a.Equal(actual.Title, expected.Title, "exercise title")
		a.Equal(actual.CourseId, expected.CourseId, "exercise course id")
		a.Equal(*actual.MaxScore, *expected.MaxScore, "exercise max score")
		a.Equal(*actual.MinScore, *expected.MinScore, "exercise min score")
		a.Equal(actual.OrderIndex, expected.OrderIndex, "exercise min score")
		a.Equal(actual.Description, expected.Description, "exercise description")
		a.Equal(*actual.EstimatedTime, *expected.EstimatedTime, "exercise estimated time")
	})
}

func (a *Assertion) HTTPUpdateExerciseResponse(
	actual models.ExerciseResponse,
	expected builders.ExerciseUpdate,
) {
	message := "Check update HTTP exercise response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.HTTPExerciseUpdated(actual.Exercise, expected)
	})
}

func (a *Assertion) HTTPExerciseUpdated(actual models.Exercise, expected builders.ExerciseUpdate) {
	message := "Check updated HTTP exercise"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual.MaxScore, "exercise max score")
		a.NotNil(actual.MinScore, "exercise min score")
		a.NotNil(actual.EstimatedTime, "exercise estimated time")
		a.Equal(actual.Title, expected.Title, "exercise title")
		a.Equal(*actual.MaxScore, expected.MaxScore, "exercise max score")
		a.Equal(*actual.MinScore, expected.MinScore, "exercise min score")
		a.Equal(actual.OrderIndex, expected.OrderIndex, "exercise order index")
		a.Equal(actual.Description, expected.Description, "exercise description")
		a.Equal(*actual.EstimatedTime, expected.EstimatedTime, "exercise estimated time")
	})
}

func (a *Assertion) HTTPListExercisesResponse(
	actual models.ExercisesResponse,
	expected models.ExerciseResponse,
) {
	message := "Check list HTTP exercises response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.HTTPExercisesContain(actual.Exercises, expected.Exercise)
	})
}

func (a *Assertion) HTTPExercisesContain(actual []models.Exercise, expected models.Exercise) {
	message := "Check HTTP exercises contain expected course"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		for _, exercise := range actual {
			if exercise.Id == expected.Id {
				a.HTTPExercise(exercise, expected)
				return
			}
		}

		a.Fail("Expected HTTP exercise was not found")
	})
}

func (a *Assertion) GRPCCreateExerciseResponse(
	actual *v1.GetExerciseResponse,
	expected builders.ExerciseCreate,
) {
	message := "Check create gRPC exercise response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "create exercise response")
		a.GRPCExerciseCreated(actual.GetExercise(), expected)
	})
}

func (a *Assertion) GRPCExerciseCreated(
	actual *v1.Exercise,
	expected builders.ExerciseCreate,
) {
	message := "Check created gRPC exercise"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "exercise")
		a.NotEmpty(actual.GetId(), "exercise Id")
		a.NotNil(actual.MaxScore, "exercise max score")
		a.NotNil(actual.MinScore, "exercise min score")
		a.NotNil(actual.EstimatedTime, "exercise estimated time")
		a.Equal(actual.GetTitle(), expected.Title, "exercise title")
		a.Equal(actual.GetCourseId(), expected.CourseId, "exercise course id")
		a.Equal(actual.GetMaxScore(), int32(expected.MaxScore), "exercise max score")
		a.Equal(actual.GetMinScore(), int32(expected.MinScore), "exercise min score")
		a.Equal(actual.GetOrderIndex(), int32(expected.OrderIndex), "exercise order index")
		a.Equal(actual.GetDescription(), expected.Description, "exercise description")
		a.Equal(actual.GetEstimatedTime(), expected.EstimatedTime, "exercise estimated time")
	})
}

func (a *Assertion) GRPCGetExerciseResponse(
	actual *v1.GetExerciseResponse,
	expected *v1.GetExerciseResponse,
) {
	message := "Check get gRPC exercise response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "actual get exercise response")
		a.NotNil(expected, "expected get exercise response")
		a.GRPCExercise(actual.GetExercise(), expected.GetExercise())
	})
}

func (a *Assertion) GRPCExercise(
	actual *v1.Exercise,
	expected *v1.Exercise,
) {
	message := "Check gRPC exercise"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "actual exercise")
		a.NotNil(expected, "expected exercise")
		a.NotNil(actual.MaxScore, "actual exercise max score")
		a.NotNil(expected.MaxScore, "expected exercise max score")
		a.NotNil(actual.MinScore, "actual exercise min score")
		a.NotNil(expected.MinScore, "expected exercise min score")
		a.NotNil(actual.EstimatedTime, "actual exercise estimated time")
		a.NotNil(expected.EstimatedTime, "expected exercise estimated time")
		a.Equal(actual.GetId(), expected.GetId(), "exercise ID")
		a.Equal(actual.GetTitle(), expected.GetTitle(), "exercise title")
		a.Equal(actual.GetCourseId(), expected.GetCourseId(), "exercise course id")
		a.Equal(actual.GetMaxScore(), expected.GetMaxScore(), "exercise max score")
		a.Equal(actual.GetMinScore(), expected.GetMinScore(), "exercise min score")
		a.Equal(actual.GetOrderIndex(), expected.GetOrderIndex(), "exercise order index")
		a.Equal(actual.GetDescription(), expected.GetDescription(), "exercise description")
		a.Equal(actual.GetEstimatedTime(), expected.GetEstimatedTime(), "exercise estimated time")
	})
}

func (a *Assertion) GRPCUpdateExerciseResponse(
	actual *v1.GetExerciseResponse,
	expected builders.ExerciseUpdate,
) {
	message := "Check update gRPC exercise response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "update exercise response")
		a.NotNil(actual.GetExercise(), "exercise")
		a.GRPCExerciseUpdated(actual.GetExercise(), expected)
	})
}

func (a *Assertion) GRPCExerciseUpdated(
	actual *v1.Exercise,
	expected builders.ExerciseUpdate,
) {
	message := "Check updated gRPC exercise"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "exercise")
		a.NotNil(actual.MaxScore, "exercise max score")
		a.NotNil(actual.MinScore, "exercise min score")
		a.NotNil(actual.EstimatedTime, "exercise estimated time")
		a.Equal(actual.GetTitle(), expected.Title, "exercise title")
		a.Equal(actual.GetMaxScore(), int32(expected.MaxScore), "exercise max score")
		a.Equal(actual.GetMinScore(), int32(expected.MinScore), "exercise min score")
		a.Equal(actual.GetOrderIndex(), int32(expected.OrderIndex), "exercise order index")
		a.Equal(actual.GetDescription(), expected.Description, "exercise description")
		a.Equal(actual.GetEstimatedTime(), expected.EstimatedTime, "exercise estimated time")
	})
}

func (a *Assertion) GRPCListExercisesResponse(
	actual *v1.ListExercisesResponse,
	expected *v1.GetExerciseResponse,
) {
	message := "Check list gRPC exercises response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "list exercises response")
		a.NotNil(expected, "expected exercise response")
		a.GRPCExercisesContain(actual.GetExercises(), expected.GetExercise())
	})
}

func (a *Assertion) GRPCExercisesContain(
	actual []*v1.Exercise,
	expected *v1.Exercise,
) {
	message := "Check gRPC exercises contain expected exercise"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		for _, exercise := range actual {
			if exercise.GetId() == expected.GetId() {
				a.GRPCExercise(exercise, expected)
				return
			}
		}

		a.Fail("Expected gRPC exercise was not found")
	})
}
