package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Cfg Config

func init() {
	loadConfig()
}

const defaultConfigFile = "config.yaml"

// LoadConfig 目前只支持本地配置文件的方式
// 1. 配置中心方案：在镜像中设置配置中心地址常量，使用os.Getenv来获取地址信息
// 2. k8s方案：读取config资源
func loadConfig() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置文件加载失败: %+v", err))
	}

	// 动态监测配置文件变更
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(&Cfg); err != nil {
			panic(fmt.Errorf("配置文件解析错误: %+v", err))
		}
	})

	if err := v.Unmarshal(&Cfg); err != nil {
		panic(fmt.Errorf("配置文件解析错误: %+v", err))
	}
}
