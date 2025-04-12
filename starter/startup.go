package starter

import (
	"context"
	"fmt"
	"time"

	"github.com/abmpio/abmp/pkg/log"
	"github.com/abmpio/app"
	"github.com/abmpio/app/cli"
	"github.com/abmpio/resourcex-sdk/options"
	pb "github.com/abmpio/resourcex-sdk/proto"
	"github.com/abmpio/resourcex-sdk/sdk"
)

func init() {
	fmt.Println("abmpio.resource-sdk.starter init")

	cli.ConfigureService(serviceConfigurator)
}

func serviceConfigurator(wa cli.CliApplication) {
	if app.HostApplication.SystemConfig().App.IsRunInCli {
		return
	}
	var _client sdk.IClient

	opt := options.GetOptions()
	if !opt.Disabled {
		resourcexClient := sdk.NewClient(sdk.WithHost(opt.Host), sdk.WithPort(opt.Port))
		//测试ping
		for {
			err := resourcexClient.InitConnnection()
			if err != nil {
				log.Logger.Warn(fmt.Sprintf("初始化resourcex grpc连接时出现异常,option:%s, err:%s",
					opt.String(),
					err.Error()))
			} else {
				res, err := resourcexClient.ResourcexHealthCheck(context.TODO(), &pb.ResourcexHealthCheckRequest{})
				if err != nil {
					log.Logger.Warn(fmt.Sprintf("检测resourcex grpc 服务健康是否正常时出现异常,option:%s, err:%s",
						opt.String(),
						err.Error()))
				} else {
					if res != nil {
						log.Logger.Info(fmt.Sprintf("resourcex grpc status:%s", res.Status.String()))
					}
					// set client
					_client = resourcexClient
					break
				}
			}
			log.Logger.Warn("2s后重新测试...")
			time.Sleep(2 * time.Second)
		}
	} else {
		log.Logger.Warn("resourcex grpc client disabled")
		_client = &sdk.NullableClient{}
	}
	sdk.SetGlobalClient(_client)
	app.Context.RegistInstanceAs(_client, new(pb.ResourcexClient))
	app.Context.RegistInstanceAs(_client, new(sdk.Client))
}
