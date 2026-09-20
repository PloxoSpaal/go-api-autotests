package assertions

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

func (a *Assertion) HTTPCreateCourseResponse(
	actual models.CourseResponse,
	expected builders.CourseCreate,
) {
	message := "Check create HTTP course response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.HTTPCourseCreated(actual.Course, expected)
	})
}

func (a *Assertion) HTTPCourseCreated(actual models.Course, expected builders.CourseCreate) {
	message := "Check created HTTP course"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
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
	message := "Check get HTTP course response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.HTTPCourse(actual.Course, expected.Course)
	})
}

func (a *Assertion) HTTPCourse(actual, expected models.Course) {
	message := "Check HTTP course"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual.MaxScore, "actual course max score")
		a.NotNil(expected.MaxScore, "expected course max score")
		a.NotNil(actual.MinScore, "actual course min score")
		a.NotNil(expected.MinScore, "expected course min score")
		a.NotNil(actual.EstimatedTime, "actual course estimated time")
		a.NotNil(expected.EstimatedTime, "expected course estimated time")
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
	message := "Check update HTTP course response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.HTTPCourseUpdated(actual.Course, expected)
	})
}

func (a *Assertion) HTTPCourseUpdated(actual models.Course, expected builders.CourseUpdate) {
	message := "Check updated HTTP course"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
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
	message := "Check list HTTP courses response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.HTTPCoursesContain(actual.Courses, expected.Course)
	})
}

func (a *Assertion) HTTPCoursesContain(actual []models.Course, expected models.Course) {
	message := "Check HTTP courses contain expected course"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		for _, course := range actual {
			if course.ID == expected.ID {
				a.HTTPCourse(course, expected)
				return
			}
		}

		a.Fail("Expected HTTP course was not found")
	})
}

func (a *Assertion) GRPCCreateCourseResponse(actual *v1.GetCourseResponse, expected builders.CourseCreate) {
	message := "Check create gRPC course response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "create course response")
		a.GRPCCourseCreated(actual.GetCourse(), expected)
	})
}

func (a *Assertion) GRPCCourseCreated(actual *v1.Course, expected builders.CourseCreate) {
	message := "Check created gRPC course"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "course")
		a.NotEmpty(actual.GetId(), "course ID")
		a.NotNil(actual.MaxScore, "course max score")
		a.NotNil(actual.MinScore, "course min score")
		a.NotNil(actual.EstimatedTime, "course estimated time")
		a.Equal(actual.GetTitle(), expected.Title, "course title")
		a.Equal(actual.GetMaxScore(), int32(expected.MaxScore), "course max score")
		a.Equal(actual.GetMinScore(), int32(expected.MinScore), "course min score")
		a.Equal(actual.GetDescription(), expected.Description, "course description")
		a.Equal(actual.GetEstimatedTime(), expected.EstimatedTime, "course estimated time")
		a.Equal(actual.GetPreviewFile().GetId(), expected.PreviewFileID, "preview file ID")
		a.Equal(actual.GetCreatedByUser().GetId(), expected.CreatedByUserID, "creator user ID")
	})
}

func (a *Assertion) GRPCGetCourseResponse(actual, expected *v1.GetCourseResponse) {
	message := "Check get gRPC course response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "actual get course response")
		a.NotNil(expected, "expected get course response")
		a.GRPCCourse(actual.GetCourse(), expected.GetCourse())
	})
}

func (a *Assertion) GRPCCourse(actual, expected *v1.Course) {
	message := "Check gRPC course"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "actual course")
		a.NotNil(expected, "expected course")
		a.NotNil(actual.MaxScore, "actual course max score")
		a.NotNil(expected.MaxScore, "expected course max score")
		a.NotNil(actual.MinScore, "actual course min score")
		a.NotNil(expected.MinScore, "expected course min score")
		a.NotNil(actual.EstimatedTime, "actual course estimated time")
		a.NotNil(expected.EstimatedTime, "expected course estimated time")
		a.Equal(actual.GetId(), expected.GetId(), "course ID")
		a.Equal(actual.GetTitle(), expected.GetTitle(), "course title")
		a.Equal(actual.GetMaxScore(), expected.GetMaxScore(), "course max score")
		a.Equal(actual.GetMinScore(), expected.GetMinScore(), "course min score")
		a.Equal(actual.GetDescription(), expected.GetDescription(), "course description")
		a.Equal(actual.GetEstimatedTime(), expected.GetEstimatedTime(), "course estimated time")
		a.GRPCFile(actual.GetPreviewFile(), expected.GetPreviewFile())
		a.GRPCUser(actual.GetCreatedByUser(), expected.GetCreatedByUser())
	})
}

func (a *Assertion) GRPCUpdateCourseResponse(actual *v1.GetCourseResponse, expected builders.CourseUpdate) {
	message := "Check update gRPC course response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "update course response")
		a.GRPCCourseUpdated(actual.GetCourse(), expected)
	})
}

func (a *Assertion) GRPCCourseUpdated(actual *v1.Course, expected builders.CourseUpdate) {
	message := "Check updated gRPC course"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "course")
		a.NotNil(actual.MaxScore, "course max score")
		a.NotNil(actual.MinScore, "course min score")
		a.NotNil(actual.EstimatedTime, "course estimated time")
		a.Equal(actual.GetTitle(), expected.Title, "course title")
		a.Equal(actual.GetMaxScore(), int32(expected.MaxScore), "course max score")
		a.Equal(actual.GetMinScore(), int32(expected.MinScore), "course min score")
		a.Equal(actual.GetDescription(), expected.Description, "course description")
		a.Equal(actual.GetEstimatedTime(), expected.EstimatedTime, "course estimated time")
	})
}

func (a *Assertion) GRPCListCoursesResponse(actual *v1.ListCoursesResponse, expected *v1.GetCourseResponse) {
	message := "Check list gRPC courses response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "list courses response")
		a.NotNil(expected, "expected course response")
		a.GRPCCoursesContain(actual.GetCourses(), expected.GetCourse())
	})
}

func (a *Assertion) GRPCCoursesContain(actual []*v1.Course, expected *v1.Course) {
	message := "Check gRPC courses contain expected course"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		for _, course := range actual {
			if course.GetId() == expected.GetId() {
				a.GRPCCourse(course, expected)
				return
			}
		}

		a.Fail("Expected gRPC course was not found")
	})
}
