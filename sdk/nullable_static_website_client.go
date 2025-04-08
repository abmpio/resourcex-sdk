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

// 获取静态站点的页面文件列表
func (*nullableStaticWebsiteClient) StaticWebsitePageList(ctx context.Context, in *pb.StaticWebsitePageListRequest, opts ...grpc.CallOption) (*pb.StaticWebsitePageListResponse, error) {
	log.Logger.Warn("nullableStaticWebsiteClient.StaticWebsitePageList method")
	return nil, nil
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
