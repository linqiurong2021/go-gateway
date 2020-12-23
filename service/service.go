package service

import "github.com/linqiurong2021/go-gateway/etcd"

// AddConf 新增
func AddConf(conf *etcd.EtcdProxyConfItem, confList etcd.EtcdProxyConf) etcd.EtcdProxyConf {
	//
	newConfList := append(confList[0:], conf)
	return newConfList
}

// DelConf 删除
func DelConf(ID uint, confList etcd.EtcdProxyConf) etcd.EtcdProxyConf {
	for index, conf := range confList {
		if ID == conf.ID {
			confList = append(confList[:index], confList[index+1:]...)
		}
	}
	return confList
}

// UpdateConf 更新
func UpdateConf(updateConf *etcd.EtcdProxyConfItem, confList etcd.EtcdProxyConf) etcd.EtcdProxyConf {
	for index, confItem := range confList {
		if confItem.ID == updateConf.ID {
			//
			confList = append(confList[:index], confList[index+1:]...)
		}
	}
	return AddConf(updateConf, confList)
}
