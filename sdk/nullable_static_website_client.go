package sdk

import (
	"context"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/resourcex-sdk/proto"
	"google.golang.org/grpc"
)

type nullableStaticWebsiteClient struct {
}

var _ pb.StaticWebsiteClient = (*nullableStaticWebsiteClient)(nil)

func (*nullableStaticWebsiteClient) StaticWebsiteHealthCheck(ctx context.Context, in *pb.StaticWebsiteHealthCheckRequest, opts ...grpc.CallOption) (*pb.StaticWebsiteHealthCheckResponse, error) {
	return &pb.StaticWebsiteHealthCheckResponse{
		Status: pb.StaticWebsiteHealthCheckResponse_NOT_SERVING,
	}, nil
}

func (*nullableStaticWebsiteClient) StaticWebsiteUploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[pb.StaticWebsiteUploadFileChunk, pb.StaticWebsiteUploadFileResponse], error) {
	log.Logger.Warn("nullableStaticWebsiteClient.StaticWebsiteUploadFile method")
	return nil, nil
}

// 删除静态站点文件
func (*nullableStaticWebsiteClient) StaticWebsiteDeleteFile(ctx context.Context, in *pb.StaticWebsiteDeleteFileRequest, opts ...grpc.CallOption) (*pb.StaticWebsiteDeleteFileResponse, error) {
	log.Logger.Warn("nullableStaticWebsiteClient.StaticWebsiteDeleteFile method")
	return nil, nil
}

// 发布站点
func (*nullableStaticWebsiteClient) StaticWebsitePublish(ctx context.Context, in *pb.StaticWebsitePublishRequest, opts ...grpc.CallOption) (*pb.StaticWebsitePublishResponse, error) {
	log.Logger.Warn("nullableStaticWebsiteClient.StaticWebsitePublish method")
	return nil, nil
}
