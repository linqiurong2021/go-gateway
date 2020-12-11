package main

import (
	"fmt"
	"time"

	"github.com/linqiurong2021/go-gateway/agent"
	"github.com/linqiurong2021/go-gateway/config"
	"github.com/linqiurong2021/go-gateway/model"
	"github.com/linqiurong2021/go-gateway/mysql"
	"go.etcd.io/etcd/clientv3"
)

func main() {
	// //
	// if err := config.Init("./config/config.ini"); err != nil {
	// 	fmt.Printf("load config from file falure !, err:%v\n", err)
	// 	return
	// }
	// if err := mysql.InitMySQL(config.Conf.MySQLConfig); err != nil {
	// 	fmt.Printf("init mysql failed, err:%v\n", err)
	// 	return
	// }
	// // 模型绑定
	// mysql.DB.AutoMigrate(&model.Service{})
	// // // 初始化数据
	// fmt.Printf("%#v", config.Conf.EtcdConfig)
	// // etcd.InitEtcd(config.Conf.EtcdConfig)

	// cli, err := clientv3.New(clientv3.Config{
	// 	Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
	// 	DialTimeout: 5 * time.Second,
	// })
	// if err != nil {
	// 	// handle error!
	// }
	// defer cli.Close()

	// // 实例化agent并启用
	// new(agent.Agent).Start()
}
