package grpcclients

import (
	"context"

	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"google.golang.org/grpc"
)

type CoursesClient struct {
	client v1.CoursesServiceClient
}

func NewCoursesClient(connection grpc.ClientConnInterface) *CoursesClient {
	return &CoursesClient{client: v1.NewCoursesServiceClient(connection)}
}

func (c *CoursesClient) List(
	ctx context.Context,
	request *v1.ListCoursesRequest,
) (*v1.ListCoursesResponse, error) {
	return c.client.ListCourses(ctx, request)
}

func (c *CoursesClient) Create(
	ctx context.Context,
	request *v1.CreateCourseRequest,
) (*v1.GetCourseResponse, error) {
	return c.client.CreateCourse(ctx, request)
}

func (c *CoursesClient) Get(
	ctx context.Context,
	request *v1.GetCourseRequest,
) (*v1.GetCourseResponse, error) {
	return c.client.GetCourse(ctx, request)
}

func (c *CoursesClient) Update(
	ctx context.Context,
	request *v1.UpdateCourseRequest,
) (*v1.GetCourseResponse, error) {
	return c.client.UpdateCourse(ctx, request)
}

func (c *CoursesClient) Delete(
	ctx context.Context,
	request *v1.DeleteCourseRequest,
) (*v1.Empty, error) {
	return c.client.DeleteCourse(ctx, request)
}
