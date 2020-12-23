package logic

import (
	"encoding/json"
	"fmt"

	"github.com/linqiurong2021/go-gateway/config"
	"github.com/linqiurong2021/go-gateway/etcd"
	"github.com/linqiurong2021/go-gateway/service"
)

// proxyConfList 代理配置列表
var proxyConfList etcd.EtcdProxyConf

func init() {
	proxyConfList = getEtcdConfList(config.Conf.EtcdConfig.Key)
}

// AddConf 新增
func AddConf(key string, conf *etcd.EtcdProxyConfItem) {
	//
	list := service.AddConf(conf, proxyConfList)
	//
	result, err := marshal(list)
	if err != nil {
		fmt.Printf("\n json marshal failure, err: %s \n", err.Error())
		return
	}
	SaveToEtcd(result)
}

// DelConf 删除
func DelConf(ID uint) {
	list := service.DelConf(ID, proxyConfList)
	result, err := marshal(list)
	if err != nil {
		fmt.Printf("\n json marshal failure, err: %s \n", err.Error())
		return
	}
	SaveToEtcd(result)
}

// UpdateConf 更新
func UpdateConf(updateConf *etcd.EtcdProxyConfItem) {
	list := service.UpdateConf(updateConf, proxyConfList)
	result, err := marshal(list)
	if err != nil {
		fmt.Printf("\n json marshal failure, err: %s \n", err.Error())
		return
	}
	SaveToEtcd(result)
}

// getEtcdConfList 获取配置信息
func getEtcdConfList(key string) etcd.EtcdProxyConf {
	return etcd.GetProxyConfList(key)
}

// json 序列化
func marshal(list etcd.EtcdProxyConf) (result []byte, err error) {
	result, err = json.Marshal(list)
	if err != nil {
		return nil, err
	}
	return
}

// SaveToEtcd 保存到etcd
func SaveToEtcd(json []byte) {
	resp, err := etcd.Set(config.Conf.EtcdConfig.Key, string(json))
	if err != nil {
		fmt.Printf(" save to etcd failure, err: %s", err.Error())
		return
	}
	fmt.Printf("\n%v\n", resp)
}
