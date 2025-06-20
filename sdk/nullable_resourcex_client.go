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

func (*nullableResourcexClient) ResourcexGetContent(ctx context.Context, in *pb.ResourcexGetContentRequest, opts ...grpc.CallOption) (*pb.ResourcexGetContentResponse, error) {
	log.Logger.Warn("nullableResourcexClient.ResourcexGetContent method")
	return nil, nil
}

// 流式获取文件内容
func (*nullableResourcexClient) ResourcexGetContentStream(ctx context.Context, in *pb.ResourcexGetContentStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pb.ResourcexGetContentStreamChunk], error) {
	log.Logger.Warn("nullableResourcexClient.ResourcexGetContentStream method")
	return nil, nil
}

// 压缩一组文件资源为zip文件
func (*nullableResourcexClient) ResourcexGetFileListContentStream(ctx context.Context, in *pb.ResourcexGetFileListContentStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pb.ResourcexGetFileListContentStreamChunk], error) {
	log.Logger.Warn("nullableResourcexClient.ResourcexGetFileListContentStream method")
	return nil, nil
}
