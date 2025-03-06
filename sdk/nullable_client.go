package sdk

import (
	"context"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/resourcex-sdk/proto"
	"google.golang.org/grpc"
)

type NullableClient struct {
}

var _ IClient = (*NullableClient)(nil)
var _ pb.ResourcexClient = (*NullableClient)(nil)

func (*NullableClient) ResourcexHealthCheck(ctx context.Context, in *pb.ResourcexHealthCheckRequest, opts ...grpc.CallOption) (*pb.ResourcexHealthCheckResponse, error) {
	return &pb.ResourcexHealthCheckResponse{
		Status: pb.ResourcexHealthCheckResponse_NOT_SERVING,
	}, nil
}

func (*NullableClient) ResourcexUploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[pb.ResourcexUploadFileChunk, pb.ResourcexUploadFileResponse], error) {
	log.Logger.Warn("NullableClient.ResourcexUploadFile method")
	return nil, nil
}
