package sdk

import (
	"fmt"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/resourcex-sdk/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// client interface
type IClient interface {
	pb.ResourcexClient
}

type Client struct {
	option *Options

	conn *grpc.ClientConn

	pb.ResourcexClient
}

var _ IClient = (*Client)(nil)

func NewClient(opts ...Option) *Client {
	client := &Client{
		option: newDefaultOptions(),
	}
	for _, eachOpt := range opts {
		eachOpt(client.option)
	}
	return client
}

func (c *Client) GetOption() *Options {
	return c.option
}

// 初始化client
func (c *Client) InitConnnection(opts ...grpc.DialOption) error {
	mergedOpts := make([]grpc.DialOption, 0)
	mergedOpts = append(mergedOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	mergedOpts = append(mergedOpts, grpc.WithKeepaliveParams(*c.option.keepaliveParam))
	mergedOpts = append(mergedOpts, opts...)
	conn, err := grpc.NewClient(c.option.getHostTarget(), mergedOpts...)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("occur error when create hi grpc server connection , host:%s,error: %s",
			c.option.getHostTarget(), err.Error()))
		return err
	}
	log.Logger.Info(fmt.Sprintf("initialize hi grpc connection finished,host:%s", c.option.getHostTarget()))
	c.conn = conn
	//保存客户端
	c.ResourcexClient = pb.NewResourcexClient(conn)

	return nil
}
