package plugins

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/Nikita-Filonov/axiom"
	"github.com/Nikita-Filonov/axiom/plugins/testallure"
	"github.com/PloxoSpaal/go-api-autotests/config"
	"github.com/PloxoSpaal/go-api-autotests/resources"
	allure "github.com/allure-framework/allure-go/commons/gotest"
	allurewriter "github.com/allure-framework/allure-go/commons/writer"
)

func Allure() axiom.Plugin {
	var (
		once     sync.Once
		delegate axiom.Plugin
		setupErr error
	)

	return func(cfg *axiom.Config) {
		once.Do(func() {
			delegate, setupErr = newAllurePlugin(cfg.Runner)
		})
		if setupErr != nil {
			panic(fmt.Errorf("setup allure plugin: %w", setupErr))
		}

		delegate(cfg)
	}
}

func newAllurePlugin(runner *axiom.Runner) (axiom.Plugin, error) {
	resultsDir, err := allureResultsDir()
	if err != nil {
		return nil, err
	}

	writer := allurewriter.NewFileSystemWriter(resultsDir)
	settings := resources.GetConfigResource(runner)

	if err = writer.WriteEnvironmentInfo(context.Background(), buildEnvironment(settings)); err != nil {
		return nil, fmt.Errorf("write allure environment: %w", err)
	}

	return testallure.Plugin(allure.WithWriter(writer)), nil
}

func allureResultsDir() (string, error) {
	if resultsDir := os.Getenv("ALLURE_RESULTS_DIR"); resultsDir != "" {
		return filepath.Abs(resultsDir)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get working directory: %w", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return filepath.Join(currentDir, "allure-results"), nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return "", fmt.Errorf("go.mod not found")
		}

		currentDir = parentDir
	}
}

func buildEnvironment(settings *config.Config) map[string]string {
	return map[string]string{
		"http.url":     settings.HTTP.URL,
		"http.timeout": settings.HTTP.Timeout.String(),
		"grpc.address": settings.GRPC.Address,
		"grpc.timeout": settings.GRPC.Timeout.String(),
		"go.version":   runtime.Version(),
		"os":           runtime.GOOS,
		"architecture": runtime.GOARCH,
	}
}
