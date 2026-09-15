package httpclients

import (
	"context"
	"net/url"

	"github.com/PloxoSpaal/go-api-autotests/models"
	"github.com/PloxoSpaal/go-api-autotests/transport/httptransport"
	"github.com/go-resty/resty/v2"
)

type ExercisesClient struct {
	transport *httptransport.Client
}

func NewExercisesClient(transport *httptransport.Client) *ExercisesClient {
	return &ExercisesClient{transport: transport}
}

func (c *ExercisesClient) List(
	ctx context.Context,
	request models.ListExercisesQuery,
) (Response[models.ExercisesResponse], error) {
	return execute[models.ExercisesResponse](func() (*resty.Response, error) {
		return c.transport.Get(ctx, httptransport.Request{
			URL:         "/exercises",
			QueryParams: request.ToQueryParams(),
		})
	})
}

func (c *ExercisesClient) Create(
	ctx context.Context,
	request models.CreateExerciseRequest,
) (Response[models.ExerciseResponse], error) {
	return execute[models.ExerciseResponse](func() (*resty.Response, error) {
		return c.transport.Post(ctx, httptransport.Request{
			URL:  "/exercises",
			Body: request,
		})
	})
}

func (c *ExercisesClient) Get(
	ctx context.Context,
	exerciseID string,
) (Response[models.ExerciseResponse], error) {
	return execute[models.ExerciseResponse](func() (*resty.Response, error) {
		return c.transport.Get(ctx, httptransport.Request{
			URL: "/exercises/" + url.PathEscape(exerciseID),
		})
	})
}

func (c *ExercisesClient) Update(
	ctx context.Context,
	exerciseID string,
	request models.UpdateExerciseRequest,
) (Response[models.ExerciseResponse], error) {
	return execute[models.ExerciseResponse](func() (*resty.Response, error) {
		return c.transport.Patch(ctx, httptransport.Request{
			URL:  "/exercises/" + url.PathEscape(exerciseID),
			Body: request,
		})
	})
}

func (c *ExercisesClient) Delete(
	ctx context.Context,
	exerciseID string,
) (Response[models.Empty], error) {
	return execute[models.Empty](func() (*resty.Response, error) {
		return c.transport.Delete(ctx, httptransport.Request{
			URL: "/exercises/" + url.PathEscape(exerciseID),
		})
	})
}
