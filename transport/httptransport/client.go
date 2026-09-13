package httptransport

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/config"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	resty *resty.Client
}

func New(client *resty.Client) *Client {
	return &Client{resty: client}
}

type Request struct {
	URL         string
	Headers     map[string]string
	QueryParams map[string]string
	Body        any
	FormData    map[string]string
	Files       []MultipartFile
}

type MultipartFile struct {
	FieldName string
	Filename  string
	Content   io.Reader
}

func newRestyClient(cfg *axiom.Config, settings config.HTTP) *resty.Client {
	return resty.New().
		SetBaseURL(settings.URL).
		SetTimeout(settings.Timeout).
		SetHeader("Accept", "application/json").
		OnBeforeRequest(onBeforeRequestHook(cfg)).
		OnAfterResponse(onAfterResponseHook(cfg)).
		OnError(onErrorHook(cfg))
}

func NewPublic(cfg *axiom.Config, settings config.HTTP) *Client {
	return New(newRestyClient(cfg, settings))
}

func NewPrivate(cfg *axiom.Config, settings config.HTTP, token string) *Client {
	return New(newRestyClient(cfg, settings).SetAuthToken(token))
}

func (c *Client) Do(
	ctx context.Context,
	method string,
	request Request,
) (*resty.Response, error) {
	if err := request.validate(); err != nil {
		return nil, err
	}

	restyRequest := c.resty.R().
		SetContext(ctx).
		SetHeaders(request.Headers).
		SetQueryParams(request.QueryParams)

	if request.Body != nil {
		restyRequest.SetBody(request.Body)
	}
	if len(request.FormData) > 0 {
		restyRequest.SetFormData(request.FormData)
	}
	for _, file := range request.Files {
		restyRequest.SetFileReader(file.FieldName, file.Filename, file.Content)
	}

	return restyRequest.Execute(method, request.URL)
}

func (c *Client) Get(ctx context.Context, request Request) (*resty.Response, error) {
	return c.Do(ctx, http.MethodGet, request)
}

func (c *Client) Post(ctx context.Context, request Request) (*resty.Response, error) {
	return c.Do(ctx, http.MethodPost, request)
}

func (c *Client) Patch(ctx context.Context, request Request) (*resty.Response, error) {
	return c.Do(ctx, http.MethodPatch, request)
}

func (c *Client) Put(ctx context.Context, request Request) (*resty.Response, error) {
	return c.Do(ctx, http.MethodPut, request)
}

func (c *Client) Delete(ctx context.Context, request Request) (*resty.Response, error) {
	return c.Do(ctx, http.MethodDelete, request)
}

func (r Request) validate() error {
	body := r.Body
	file := r.Files
	formData := r.FormData

	if body != nil && (formData != nil || file != nil) {
		return errors.New("httptransport: body cannot be combined with form data or files")
	}

	for i, file := range r.Files {
		if file.FieldName == "" {
			return fmt.Errorf("httptransport: file %d has an empty field name", i)
		}
		if file.Filename == "" {
			return fmt.Errorf("httptransport: file %d has an empty filename", i)
		}
		if file.Content == nil {
			return fmt.Errorf("httptransport: file %d has no content", i)
		}
	}

	return nil
}
