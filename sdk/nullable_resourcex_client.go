package sdk

import (
	"context"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/resourcex-sdk/proto"
	"google.golang.org/grpc"
)

type nullableResourcexClient struct {
}

var _ pb.ResourcexClient = (*nullableResourcexClient)(nil)

func (*nullableResourcexClient) ResourcexHealthCheck(ctx context.Context, in *pb.ResourcexHealthCheckRequest, opts ...grpc.CallOption) (*pb.ResourcexHealthCheckResponse, error) {
	return &pb.ResourcexHealthCheckResponse{
		Status: pb.ResourcexHealthCheckResponse_NOT_SERVING,
	}, nil
}

func (*nullableResourcexClient) ResourcexUploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[pb.ResourcexUploadFileChunk, pb.ResourcexUploadFileResponse], error) {
	log.Logger.Warn("nullableResourcexClient.ResourcexUploadFile method")
	return nil, nil
}
