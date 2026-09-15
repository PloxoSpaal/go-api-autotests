package httpclients

import (
	"context"
	"net/url"

	"github.com/PloxoSpaal/go-api-autotests/models"
	"github.com/PloxoSpaal/go-api-autotests/transport/httptransport"
	"github.com/go-resty/resty/v2"
)

type CoursesClient struct {
	transport *httptransport.Client
}

func NewCoursesClient(transport *httptransport.Client) *CoursesClient {
	return &CoursesClient{transport: transport}
}

func (c *CoursesClient) List(
	ctx context.Context,
	request models.ListCoursesQuery,
) (Response[models.CoursesResponse], error) {
	return execute[models.CoursesResponse](func() (*resty.Response, error) {
		return c.transport.Get(ctx, httptransport.Request{
			URL:         "/courses",
			QueryParams: request.ToQueryParams(),
		})
	})
}

func (c *CoursesClient) Create(
	ctx context.Context,
	request models.CreateCourseRequest,
) (Response[models.CourseResponse], error) {
	return execute[models.CourseResponse](func() (*resty.Response, error) {
		return c.transport.Post(ctx, httptransport.Request{
			URL:  "/courses",
			Body: request,
		})
	})
}

func (c *CoursesClient) Get(
	ctx context.Context,
	courseID string,
) (Response[models.CourseResponse], error) {
	return execute[models.CourseResponse](func() (*resty.Response, error) {
		return c.transport.Get(ctx, httptransport.Request{
			URL: "/courses/" + url.PathEscape(courseID),
		})
	})
}

func (c *CoursesClient) Update(
	ctx context.Context,
	courseID string,
	request models.UpdateCourseRequest,
) (Response[models.CourseResponse], error) {
	return execute[models.CourseResponse](func() (*resty.Response, error) {
		return c.transport.Patch(ctx, httptransport.Request{
			URL:  "/courses/" + url.PathEscape(courseID),
			Body: request,
		})
	})
}

func (c *CoursesClient) Delete(
	ctx context.Context,
	courseID string,
) (Response[models.Empty], error) {
	return execute[models.Empty](func() (*resty.Response, error) {
		return c.transport.Delete(ctx, httptransport.Request{
			URL: "/courses/" + url.PathEscape(courseID),
		})
	})
}
