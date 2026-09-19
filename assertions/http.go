package assertions

import (
	"net/http"

	"github.com/PloxoSpaal/go-api-autotests/models"
)

func (a *Assertion) HTTPStatus(actual, expected int) {
	a.Equal(actual, expected, "HTTP response status code")
}

func (a *Assertion) HTTPError(actual *models.APIError, expected string) {
	a.NotNil(actual, "HTTP API error")
	a.Equal(actual.Detail, expected, "HTTP API error detail")
}

func (a *Assertion) HTTPOK(statusCode int, err error) {
	a.NoError(err)
	a.HTTPStatus(statusCode, http.StatusOK)
}
