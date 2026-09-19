package assertions

import (
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

func (a *Assertion) HTTPCreateCourseResponse(
	actual models.CourseResponse,
	expected builders.CourseCreate,
) {
	a.cfg.Step("Check create HTTP course response", func() {
		a.HTTPCourseCreated(actual.Course, expected)
	})
}

func (a *Assertion) HTTPCourseCreated(actual models.Course, expected builders.CourseCreate) {
	a.cfg.Step("Check created HTTP course", func() {
		a.NotEmpty(actual.ID, "course ID")
		a.NotNil(actual.MaxScore, "course max score")
		a.NotNil(actual.MinScore, "course min score")
		a.NotNil(actual.EstimatedTime, "course estimated time")
		a.Equal(actual.Title, expected.Title, "course title")
		a.Equal(*actual.MaxScore, expected.MaxScore, "course max score")
		a.Equal(*actual.MinScore, expected.MinScore, "course min score")
		a.Equal(actual.Description, expected.Description, "course description")
		a.Equal(*actual.EstimatedTime, expected.EstimatedTime, "course estimated time")
		a.Equal(actual.PreviewFile.ID, expected.PreviewFileID, "preview file ID")
		a.Equal(actual.CreatedByUser.ID, expected.CreatedByUserID, "creator user ID")
	})
}

func (a *Assertion) HTTPGetCourseResponse(
	actual models.CourseResponse,
	expected models.CourseResponse,
) {
	a.cfg.Step("Check get HTTP course response", func() {
		a.HTTPCourse(actual.Course, expected.Course)
	})
}

func (a *Assertion) HTTPCourse(actual, expected models.Course) {
	a.cfg.Step("Check HTTP course", func() {
		a.NotNil(expected.MaxScore, "expected course max score")
		a.NotNil(actual.MaxScore, "actual course max score")
		a.NotNil(expected.MinScore, "expected course min score")
		a.NotNil(actual.MinScore, "actual course min score")
		a.NotNil(expected.EstimatedTime, "expected course estimated time")
		a.NotNil(actual.EstimatedTime, "actual course estimated time")
		a.Equal(actual.ID, expected.ID, "course ID")
		a.Equal(actual.Title, expected.Title, "course title")
		a.Equal(*actual.MaxScore, *expected.MaxScore, "course max score")
		a.Equal(*actual.MinScore, *expected.MinScore, "course min score")
		a.Equal(actual.Description, expected.Description, "course description")
		a.Equal(*actual.EstimatedTime, *expected.EstimatedTime, "course estimated time")
		a.HTTPFile(actual.PreviewFile, expected.PreviewFile)
		a.HTTPUser(actual.CreatedByUser, expected.CreatedByUser)
	})
}

func (a *Assertion) HTTPUpdateCourseResponse(
	actual models.CourseResponse,
	expected builders.CourseUpdate,
) {
	a.cfg.Step("Check update HTTP course response", func() {
		a.HTTPCourseUpdated(actual.Course, expected)
	})
}

func (a *Assertion) HTTPCourseUpdated(actual models.Course, expected builders.CourseUpdate) {
	a.cfg.Step("Check updated HTTP course", func() {
		a.NotNil(actual.MaxScore, "course max score")
		a.NotNil(actual.MinScore, "course min score")
		a.NotNil(actual.EstimatedTime, "course estimated time")
		a.Equal(actual.Title, expected.Title, "course title")
		a.Equal(*actual.MaxScore, expected.MaxScore, "course max score")
		a.Equal(*actual.MinScore, expected.MinScore, "course min score")
		a.Equal(actual.Description, expected.Description, "course description")
		a.Equal(*actual.EstimatedTime, expected.EstimatedTime, "course estimated time")
	})
}

func (a *Assertion) HTTPListCoursesResponse(
	actual models.CoursesResponse,
	expected models.CourseResponse,
) {
	a.cfg.Step("Check list HTTP courses response", func() {
		a.HTTPCoursesContain(actual.Courses, expected.Course)
	})
}

func (a *Assertion) HTTPCoursesContain(actual []models.Course, expected models.Course) {
	a.cfg.Step("Check HTTP courses contain expected course", func() {
		for _, course := range actual {
			if course.ID == expected.ID {
				a.HTTPCourse(course, expected)
				return
			}
		}

		a.Fail("Expected HTTP course was not found")
	})
}
