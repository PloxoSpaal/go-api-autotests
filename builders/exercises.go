package builders

import (
	coursev1 "github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

type ExerciseCreate struct {
	Title         string
	CourseId      string
	MaxScore      int
	MinScore      int
	OrderIndex    int
	Description   string
	EstimatedTime string
}

type ExerciseCreateOption func(*ExerciseCreate)

func (b *Builder) ExerciseCreate(options ...ExerciseCreateOption) ExerciseCreate {
	exercise := ExerciseCreate{
		Title:         b.generator.Title(),
		MaxScore:      b.generator.MaxScore(),
		MinScore:      b.generator.MinScore(),
		OrderIndex:    b.generator.OrderIndex(),
		Description:   b.generator.Description(),
		EstimatedTime: b.generator.EstimatedTime(),
	}

	for _, option := range options {
		option(&exercise)
	}

	return exercise
}

func WithExerciseCreateTitle(title string) ExerciseCreateOption {
	return func(exercise *ExerciseCreate) { exercise.Title = title }
}

func WithExerciseCreateCourseId(courseId string) ExerciseCreateOption {
	return func(exercise *ExerciseCreate) { exercise.CourseId = courseId }
}

func WithExerciseCreateMaxScore(maxScore int) ExerciseCreateOption {
	return func(exercise *ExerciseCreate) { exercise.MaxScore = maxScore }
}

func WithExerciseCreateMinScore(minScore int) ExerciseCreateOption {
	return func(exercise *ExerciseCreate) { exercise.MinScore = minScore }
}

func WithExerciseCreateOrderIndex(orderIndex int) ExerciseCreateOption {
	return func(exercise *ExerciseCreate) { exercise.OrderIndex = orderIndex }
}

func WithExerciseCreateDescription(description string) ExerciseCreateOption {
	return func(exercise *ExerciseCreate) { exercise.Description = description }
}

func WithExerciseCreateEstimatedTime(estimatedTime string) ExerciseCreateOption {
	return func(exercise *ExerciseCreate) { exercise.EstimatedTime = estimatedTime }
}

func (c ExerciseCreate) HTTPRequest(courseId string) models.CreateExerciseRequest {
	return models.CreateExerciseRequest{
		Title:         c.Title,
		CourseId:      courseId,
		MaxScore:      &c.MaxScore,
		MinScore:      &c.MinScore,
		OrderIndex:    &c.OrderIndex,
		Description:   c.Description,
		EstimatedTime: &c.EstimatedTime,
	}
}

func (c ExerciseCreate) GRPCRequest(courseId string) *coursev1.CreateExerciseRequest {
	maxScore := int32(c.MaxScore)
	minScore := int32(c.MinScore)

	return &coursev1.CreateExerciseRequest{
		Title:         c.Title,
		CourseId:      courseId,
		MaxScore:      &maxScore,
		MinScore:      &minScore,
		OrderIndex:    int32(c.OrderIndex),
		Description:   c.Description,
		EstimatedTime: &c.EstimatedTime,
	}
}

type ExerciseUpdate struct {
	Title         string
	MaxScore      int
	MinScore      int
	OrderIndex    int
	Description   string
	EstimatedTime string
}

type ExerciseUpdateOption func(*ExerciseUpdate)

func (b *Builder) ExerciseUpdate(options ...ExerciseUpdateOption) ExerciseUpdate {
	update := ExerciseUpdate{
		Title:         b.generator.Title(),
		MaxScore:      b.generator.MaxScore(),
		MinScore:      b.generator.MinScore(),
		OrderIndex:    b.generator.OrderIndex(),
		Description:   b.generator.Description(),
		EstimatedTime: b.generator.EstimatedTime(),
	}

	for _, option := range options {
		option(&update)
	}

	return update
}

func WithExerciseUpdateTitle(title string) ExerciseUpdateOption {
	return func(update *ExerciseUpdate) { update.Title = title }
}

func WithExerciseUpdateMaxScore(maxScore int) ExerciseUpdateOption {
	return func(update *ExerciseUpdate) { update.MaxScore = maxScore }
}

func WithExerciseUpdateMinScore(minScore int) ExerciseUpdateOption {
	return func(update *ExerciseUpdate) { update.MinScore = minScore }
}

func WithExerciseUpdateOrderIndex(orderIndex int) ExerciseUpdateOption {
	return func(exercise *ExerciseUpdate) { exercise.OrderIndex = orderIndex }
}

func WithExerciseUpdateDescription(description string) ExerciseUpdateOption {
	return func(update *ExerciseUpdate) { update.Description = description }
}

func WithExerciseUpdateEstimatedTime(estimatedTime string) ExerciseUpdateOption {
	return func(update *ExerciseUpdate) { update.EstimatedTime = estimatedTime }
}

func (u ExerciseUpdate) HTTPRequest() models.UpdateExerciseRequest {
	return models.UpdateExerciseRequest{
		Title:         &u.Title,
		MaxScore:      &u.MaxScore,
		MinScore:      &u.MinScore,
		OrderIndex:    &u.OrderIndex,
		Description:   &u.Description,
		EstimatedTime: &u.EstimatedTime,
	}
}

func (u ExerciseUpdate) GRPCRequest(exerciseID string) *coursev1.UpdateExerciseRequest {
	maxScore := int32(u.MaxScore)
	minScore := int32(u.MinScore)
	orderIndex := int32(u.OrderIndex)

	return &coursev1.UpdateExerciseRequest{
		Id:            exerciseID,
		Title:         &u.Title,
		MaxScore:      &maxScore,
		MinScore:      &minScore,
		OrderIndex:    &orderIndex,
		Description:   &u.Description,
		EstimatedTime: &u.EstimatedTime,
	}
}
