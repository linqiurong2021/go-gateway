package config

import (
	"time"

	"gopkg.in/ini.v1"
)

// AppConfig App配置项
type AppConfig struct {
	Release         bool `ini:"release"`
	Port            uint `ini:"port"`
	*MySQLConfig    `ini:"mysql"`
	*RegisterServer `ini:"register"`
	*EtcdConfig     `ini:"etcd"`
}

// MySQLConfig 数据库配置项
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Charset  string `ini:"charset"`
}

// RegisterServer 注册服务地址
type RegisterServer struct {
	Host string `ini:"host"`
	Port uint   `ini:"port"`
	URL  string `ini:"url"`
}

// EtcdConfig Etcd集群配置文件
type EtcdConfig struct {
	Endpoints   []string      `ini:"endpoints"`
	DialTimeout time.Duration `ini:"timeout"`
}

// Conf 配置
var Conf = new(AppConfig)

// Init 初始化
func Init(file string) error {
	return ini.MapTo(Conf, file)
}
