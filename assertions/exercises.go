package assertions

import (
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

func (a *Assertion) HTTPCreateExerciseResponse(
	actual models.ExerciseResponse,
	expected builders.ExerciseCreate,
) {
	a.cfg.Step("Check create HTTP exercise response", func() {
		a.HTTPExerciseCreated(actual.Exercise, expected)
	})
}

func (a *Assertion) HTTPExerciseCreated(actual models.Exercise, expected builders.ExerciseCreate) {
	a.cfg.Step("Check created HTTP exercise", func() {
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
	a.cfg.Step("Check get HTTP exercise response", func() {
		a.HTTPExercise(actual.Exercise, expected.Exercise)
	})
}

func (a *Assertion) HTTPExercise(actual, expected models.Exercise) {
	a.cfg.Step("Check HTTP exercise", func() {
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
	a.cfg.Step("Check update HTTP exercise response", func() {
		a.HTTPExerciseUpdated(actual.Exercise, expected)
	})
}

func (a *Assertion) HTTPExerciseUpdated(actual models.Exercise, expected builders.ExerciseUpdate) {
	a.cfg.Step("Check updated HTTP exercise", func() {
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
	a.cfg.Step("Check list HTTP exercises response", func() {
		a.HTTPExercisesContain(actual.Exercises, expected.Exercise)
	})
}

func (a *Assertion) HTTPExercisesContain(actual []models.Exercise, expected models.Exercise) {
	a.cfg.Step("Check HTTP exercises contain expected course", func() {
		for _, exercise := range actual {
			if exercise.Id == expected.Id {
				a.HTTPExercise(exercise, expected)
				return
			}
		}

		a.Fail("Expected HTTP exercise was not found")
	})
}
