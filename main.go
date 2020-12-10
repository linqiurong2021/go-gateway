package main

import (
	"fmt"

	"github.com/linqiurong2021/go-gateway/agent"
	"github.com/linqiurong2021/go-gateway/config"
	"github.com/linqiurong2021/go-gateway/model"
	"github.com/linqiurong2021/go-gateway/mysql"
)

func main() {
	//
	if err := config.Init("./config/config.ini"); err != nil {
		fmt.Printf("load config from file falure !, err:%v\n", err)
		return
	}
	if err := mysql.InitMySQL(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	// 模型绑定
	mysql.DB.AutoMigrate(&model.Service{})

	// 实例化agent并启用
	new(agent.Agent).Start()
}
