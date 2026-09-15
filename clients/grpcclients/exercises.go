package grpcclients

import (
	"context"

	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"google.golang.org/grpc"
)

type ExercisesClient struct {
	client v1.ExercisesServiceClient
}

func NewExercisesClient(connection grpc.ClientConnInterface) *ExercisesClient {
	return &ExercisesClient{client: v1.NewExercisesServiceClient(connection)}
}

func (c *ExercisesClient) List(
	ctx context.Context,
	request *v1.ListExercisesRequest,
) (*v1.ListExercisesResponse, error) {
	return c.client.ListExercises(ctx, request)
}

func (c *ExercisesClient) Create(
	ctx context.Context,
	request *v1.CreateExerciseRequest,
) (*v1.GetExerciseResponse, error) {
	return c.client.CreateExercise(ctx, request)
}

func (c *ExercisesClient) Get(
	ctx context.Context,
	request *v1.GetExerciseRequest,
) (*v1.GetExerciseResponse, error) {
	return c.client.GetExercise(ctx, request)
}

func (c *ExercisesClient) Update(
	ctx context.Context,
	request *v1.UpdateExerciseRequest,
) (*v1.GetExerciseResponse, error) {
	return c.client.UpdateExercise(ctx, request)
}

func (c *ExercisesClient) Delete(
	ctx context.Context,
	request *v1.DeleteExerciseRequest,
) (*v1.Empty, error) {
	return c.client.DeleteExercise(ctx, request)
}
