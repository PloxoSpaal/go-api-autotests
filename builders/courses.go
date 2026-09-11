package builders

import (
	coursev1 "github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

type CourseCreate struct {
	Title         string
	MaxScore      *int
	MinScore      *int
	Description   string
	EstimatedTime *string
}

type CourseCreateOption func(*CourseCreate)

func (b *Builder) CourseCreate(options ...CourseCreateOption) CourseCreate {
	maxScore := b.generator.MaxScore()
	minScore := b.generator.MinScore()
	estimatedTime := b.generator.EstimatedTime()

	course := CourseCreate{
		Title:         b.generator.Title(),
		MaxScore:      &maxScore,
		MinScore:      &minScore,
		Description:   b.generator.Description(),
		EstimatedTime: &estimatedTime,
	}

	for _, option := range options {
		option(&course)
	}

	return course
}

func WithCourseCreateTitle(title string) CourseCreateOption {
	return func(course *CourseCreate) { course.Title = title }
}

func WithCourseCreateDescription(description string) CourseCreateOption {
	return func(course *CourseCreate) { course.Description = description }
}

func WithCourseCreateMaxScore(maxScore int) CourseCreateOption {
	return func(course *CourseCreate) { course.MaxScore = &maxScore }
}

func WithCourseCreateMinScore(minScore int) CourseCreateOption {
	return func(course *CourseCreate) { course.MinScore = &minScore }
}

func WithCourseCreateEstimatedTime(estimatedTime string) CourseCreateOption {
	return func(course *CourseCreate) { course.EstimatedTime = &estimatedTime }
}

func (c CourseCreate) HTTPRequest(FileID, UserID string) models.CreateCourseRequest {
	return models.CreateCourseRequest{
		Title:           c.Title,
		MaxScore:        c.MaxScore,
		MinScore:        c.MinScore,
		Description:     c.Description,
		EstimatedTime:   c.EstimatedTime,
		PreviewFileID:   FileID,
		CreatedByUserID: UserID,
	}
}

func (c CourseCreate) GRPCRequest(FileID, UserID string) *coursev1.CreateCourseRequest {
	return &coursev1.CreateCourseRequest{
		Title:           c.Title,
		MaxScore:        new(int32(*c.MaxScore)),
		MinScore:        new(int32(*c.MinScore)),
		Description:     c.Description,
		EstimatedTime:   c.EstimatedTime,
		PreviewFileId:   FileID,
		CreatedByUserId: UserID,
	}
}

type CourseUpdate struct {
	Title         *string
	MaxScore      *int
	MinScore      *int
	Description   *string
	EstimatedTime *string
}

type CourseUpdateOption func(*CourseUpdate)

func (b *Builder) CourseUpdate(options ...CourseUpdateOption) CourseUpdate {
	title := b.generator.Title()
	minScore := b.generator.MinScore()
	maxScore := b.generator.MaxScore()
	description := b.generator.Description()
	estimatedTime := b.generator.EstimatedTime()

	course := CourseUpdate{
		Title:         &title,
		MaxScore:      &maxScore,
		MinScore:      &minScore,
		Description:   &description,
		EstimatedTime: &estimatedTime,
	}

	for _, option := range options {
		option(&course)
	}

	return course
}

func WithCourseUpdateTitle(title string) CourseUpdateOption {
	return func(course *CourseUpdate) { course.Title = &title }
}

func WithCourseUpdateDescription(description string) CourseUpdateOption {
	return func(course *CourseUpdate) { course.Description = &description }
}

func WithCourseUpdateMaxScore(maxScore int) CourseUpdateOption {
	return func(course *CourseUpdate) { course.MaxScore = &maxScore }
}

func WithCourseUpdateMinScore(minScore int) CourseUpdateOption {
	return func(course *CourseUpdate) { course.MinScore = &minScore }
}

func WithCourseUpdateEstimatedTime(estimatedTime string) CourseUpdateOption {
	return func(course *CourseUpdate) { course.EstimatedTime = &estimatedTime }
}

func (c CourseUpdate) HTTPRequest() models.UpdateCourseRequest {
	return models.UpdateCourseRequest{
		Title:         c.Title,
		MaxScore:      c.MaxScore,
		MinScore:      c.MinScore,
		Description:   c.Description,
		EstimatedTime: c.EstimatedTime,
	}
}

func (c CourseUpdate) GRPCRequest(CourseId string) *coursev1.UpdateCourseRequest {
	return &coursev1.UpdateCourseRequest{
		Id:            CourseId,
		Title:         c.Title,
		MaxScore:      new(int32(*c.MaxScore)),
		MinScore:      new(int32(*c.MinScore)),
		Description:   c.Description,
		EstimatedTime: c.EstimatedTime,
	}
}
