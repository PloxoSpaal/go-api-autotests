package httpclients

import (
	"bytes"
	"context"
	"net/url"

	"github.com/PloxoSpaal/go-api-autotests/models"
	"github.com/PloxoSpaal/go-api-autotests/transport/httptransport"
	"github.com/go-resty/resty/v2"
)

type FilesClient struct {
	transport *httptransport.Client
}

func NewFilesClient(transport *httptransport.Client) *FilesClient {
	return &FilesClient{transport: transport}
}

func (c *FilesClient) Create(
	ctx context.Context,
	request models.CreateFileRequest,
) (Response[models.FileResponse], error) {
	return execute[models.FileResponse](func() (*resty.Response, error) {
		return c.transport.Post(ctx, httptransport.Request{
			URL: "/files",
			FormData: map[string]string{
				"filename":  request.Filename,
				"directory": request.Directory,
			},
			Files: []httptransport.MultipartFile{{
				FieldName: "upload_file",
				Filename:  request.Filename,
				Content:   bytes.NewReader(request.Content),
			}},
		})
	})
}

func (c *FilesClient) Get(
	ctx context.Context,
	fileID string,
) (Response[models.FileResponse], error) {
	return execute[models.FileResponse](func() (*resty.Response, error) {
		return c.transport.Get(ctx, httptransport.Request{
			URL: "/files/" + url.PathEscape(fileID),
		})
	})
}

func (c *FilesClient) Delete(
	ctx context.Context,
	fileID string,
) (Response[models.Empty], error) {
	return execute[models.Empty](func() (*resty.Response, error) {
		return c.transport.Delete(ctx, httptransport.Request{
			URL: "/files/" + url.PathEscape(fileID),
		})
	})
}
