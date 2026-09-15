package grpctransport

import (
	"context"
	"fmt"
	"time"

	"github.com/Nikita-Filonov/axiom"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func telemetryInterceptor(
	cfg *axiom.Config,
	timeout time.Duration,
) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		request any,
		response any,
		connection *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		options ...grpc.CallOption,
	) error {
		callContext, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		cfg.Step("gRPC request "+method, func() {
			attachProto(cfg, "gRPC request", request)

			if outgoingMetadata, ok := metadata.FromOutgoingContext(callContext); ok {
				attachJSON(cfg, "gRPC metadata", outgoingMetadata)
			}
		})

		err := invoker(
			callContext,
			method,
			request,
			response,
			connection,
			options...,
		)
		if err != nil {
			grpcStatus := status.Convert(err)

			cfg.Step("gRPC error "+grpcStatus.Code().String(), func() {
				cfg.Artefact(
					axiom.NewTextArtefact("gRPC status", grpcStatus.Message()),
				)
				cfg.Log(
					axiom.NewErrorLog(fmt.Sprintf("%s -> %s", method, grpcStatus.Code())),
				)
			})

			return err
		}

		cfg.Step("gRPC response "+method, func() {
			attachProto(cfg, "gRPC response", response)
			cfg.Log(axiom.NewInfoLog(method + " -> OK"))
		})

		return nil
	}
}

func attachProto(cfg *axiom.Config, name string, message any) {
	protobufMessage, ok := message.(proto.Message)
	if !ok || protobufMessage == nil {
		return
	}

	data, err := protojson.MarshalOptions{Indent: "  "}.Marshal(protobufMessage)
	if err == nil {
		cfg.Artefact(axiom.NewBytesArtefact(name, data))
	}
}

func attachJSON(cfg *axiom.Config, name string, value any) {
	artefact, err := axiom.NewJSONArtefact(name, value)
	if err == nil {
		cfg.Artefact(artefact)
	}
}
