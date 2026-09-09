package tests

import "github.com/Nikita-Filonov/axiom"

type axiomCourse struct {
	Title     string
	Published bool
}

type axiomCourseData struct {
	Title string
}

type axiomCourseClient struct{}

func (c *axiomCourseClient) Create(data axiomCourseData) *axiomCourse {
	return &axiomCourse{Title: data.Title, Published: false}
}

func (c *axiomCourseClient) Publish(course *axiomCourse) {
	course.Published = true
}

func (c *axiomCourseClient) Delete(course *axiomCourse) {
	*course = axiomCourse{}
}

func courseClientResource(_ *axiom.Runner) (any, func(), error) {
	return &axiomCourseClient{}, nil, nil
}

func courseDataFixture(cfg *axiom.Config) (any, func(), error) {
	var data axiomCourseData

	cfg.Setup("prepare fixture course data", func() {
		data = axiomCourseData{Title: "Go API Autotests"}
	})

	return data, nil, nil
}

func unpublishedCourseFixture(cfg *axiom.Config) (any, func(), error) {
	client := axiom.MustResource[*axiomCourseClient](cfg.Runner, "course-client")
	data := axiom.GetFixture[axiomCourseData](cfg, "course-data")

	var course *axiomCourse

	cfg.Setup("create fixture unpublished course", func() {
		course = client.Create(data)
	})

	cleanup := func() {
		cfg.Teardown("delete fixture unpublished course", func() {
			client.Delete(course)
		})
	}

	return course, cleanup, nil
}
