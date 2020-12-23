package config

import (
	"time"

	"gopkg.in/ini.v1"
)

// AppConfig App配置项
type AppConfig struct {
	Release     bool `ini:"release"`
	Port        uint `ini:"port"`
	*EtcdConfig `ini:"etcd"`
}

// EtcdConfig Etcd集群配置文件
type EtcdConfig struct {
	Endpoints   []string      `ini:"endpoints"`
	DialTimeout time.Duration `ini:"timeout"`
	Key         string        `ini:"key"`
}

// Conf 配置
var Conf = new(AppConfig)

// Init 初始化
func Init(file string) error {
	return ini.MapTo(Conf, file)
}
