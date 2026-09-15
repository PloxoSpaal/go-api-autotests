package grpcclients

import (
	"context"

	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"google.golang.org/grpc"
)

type FilesClient struct {
	client v1.FilesServiceClient
}

func NewFilesClient(connection grpc.ClientConnInterface) *FilesClient {
	return &FilesClient{client: v1.NewFilesServiceClient(connection)}
}

func (c *FilesClient) Create(
	ctx context.Context,
	request *v1.CreateFileRequest,
) (*v1.GetFileResponse, error) {
	return c.client.CreateFile(ctx, request)
}

func (c *FilesClient) Get(
	ctx context.Context,
	request *v1.GetFileRequest,
) (*v1.GetFileResponse, error) {
	return c.client.GetFile(ctx, request)
}

func (c *FilesClient) Delete(
	ctx context.Context,
	request *v1.DeleteFileRequest,
) (*v1.Empty, error) {
	return c.client.DeleteFile(ctx, request)
}
