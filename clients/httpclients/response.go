package httpclients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PloxoSpaal/go-api-autotests/models"
	"github.com/go-resty/resty/v2"
)

type Response[T any] struct {
	StatusCode int
	Headers    http.Header
	Data       T
	APIError   *models.APIError
	RawBody    []byte
}

func execute[T any](operation func() (*resty.Response, error)) (Response[T], error) {
	raw, err := operation()
	if err != nil {
		return Response[T]{}, err
	}
	return decode[T](raw)
}

func decode[T any](raw *resty.Response) (Response[T], error) {
	result := Response[T]{
		StatusCode: raw.StatusCode(),
		Headers:    raw.Header().Clone(),
		RawBody:    append([]byte(nil), raw.Body()...),
	}

	if len(raw.Body()) == 0 {
		return result, nil
	}
	if raw.IsError() {
		var apiError models.APIError
		if err := json.Unmarshal(raw.Body(), &apiError); err != nil {
			return result, fmt.Errorf("decode API error: %w", err)
		}
		result.APIError = &apiError
		return result, nil
	}
	if err := json.Unmarshal(raw.Body(), &result.Data); err != nil {
		return result, fmt.Errorf("decode successful response: %w", err)
	}
	return result, nil
}
