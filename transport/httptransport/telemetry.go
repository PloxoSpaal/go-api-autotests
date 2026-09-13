package httptransport

import (
	"fmt"

	"github.com/Nikita-Filonov/axiom"
	"github.com/go-resty/resty/v2"
)

func attachJSON(cfg *axiom.Config, name string, value any) {
	artefact, err := axiom.NewJSONArtefact(name, value)
	if err == nil {
		cfg.Artefact(artefact)
	}
}

func onBeforeRequestHook(cfg *axiom.Config) resty.RequestMiddleware {
	return func(client *resty.Client, request *resty.Request) error {
		cfg.Step(fmt.Sprintf("Send %s request to %s", request.Method, request.URL), func() {
			attachJSON(cfg, "Request headers", request.Header)

			if request.Body == nil {
				return
			}

			data, err := client.JSONMarshal(request.Body)
			if err == nil {
				cfg.Artefact(axiom.NewBytesArtefact("Request body", data))
			}
		})
		return nil
	}
}

func onAfterResponseHook(cfg *axiom.Config) resty.ResponseMiddleware {
	return func(_ *resty.Client, response *resty.Response) error {
		cfg.Step(fmt.Sprintf("Response for %s %s", response.Request.Method, response.Request.URL), func() {
			attachJSON(cfg, "Response headers", response.Header())

			cfg.Artefact(axiom.NewTextArtefact("Response status", response.Status()))
			cfg.Artefact(axiom.NewBytesArtefact("Response body", response.Body()))
			cfg.Log(axiom.NewInfoLog(fmt.Sprintf(
				"%s %s -> %d",
				response.Request.Method,
				response.Request.URL,
				response.StatusCode(),
			)))
		})
		return nil
	}
}
