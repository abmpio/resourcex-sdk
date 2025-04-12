package options

import (
	"fmt"
	"sync"

	"github.com/abmpio/abmp/pkg/log"
	"github.com/abmpio/app/host"
	"github.com/abmpio/configurationx"
	"github.com/go-viper/mapstructure/v2"
)

const (
	ConfigurationKey string = "resourcex-sdk"
)

var (
	_options ResourcexSdkOptions
	_once    sync.Once
)

type ResourcexSdkOptions struct {
	DefaultApp string `json:"defaultApp"`
	Host       string `json:"host"`
	Port       int32  `json:"port"`
	Disabled   bool   `json:"disabled"`
}

func (o *ResourcexSdkOptions) normalize() {
	if len(o.DefaultApp) <= 0 {
		appName := host.GetHostEnvironment().GetEnvString(host.ENV_AppName)
		o.DefaultApp = appName
	}
	if len(o.Host) <= 0 {
		o.Host = "127.0.0.1"
	}
	if o.Port <= 0 {
		o.Port = 9029
	}
}

func (o *ResourcexSdkOptions) String() string {
	return fmt.Sprintf("%s:%d,defaultApp:%s", o.Host, o.Port, o.DefaultApp)
}

func GetOptions() *ResourcexSdkOptions {
	_once.Do(func() {
		if err := configurationx.GetInstance().UnmarshFromKey(ConfigurationKey, &_options, func(dc *mapstructure.DecoderConfig) {
			dc.TagName = "json"
		}); err != nil {
			err = fmt.Errorf("无效的配置文件,%s", err)
			log.Logger.Error(err.Error())
			panic(err)
		}
		_options.normalize()
	})
	return &_options
}
