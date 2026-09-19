package assertions

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Assertion) GRPCError(actual error, expectedCode codes.Code, expectedMessage string) {
	a.Error(actual, "gRPC response")

	grpcStatus := status.Convert(actual)

	a.Equal(grpcStatus.Code(), expectedCode, "gRPC status code")
	a.Equal(grpcStatus.Message(), expectedMessage, "gRPC status message")
}
