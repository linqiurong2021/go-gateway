package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/linqiurong2021/go-gateway/config"
	"go.etcd.io/etcd/clientv3"
)

// EtcdProxyConf Etcd代理配置项
type EtcdProxyConf []*EtcdProxyConfItem

// EtcdProxyConfItem 配置项
type EtcdProxyConfItem struct {
	ID    string `gorm:"id;primaryKey"`
	Host  string `gorm:"host"`
	Port  uint   `gorm:"port"`
	URL   string `gorm:"url"`
	Name  string `gorm:"name"`
	Alive bool   `gorm:"alive"`
}

// EtcdClient EtcdClient
var EtcdClient *clientv3.Client

// Init Init
func Init(cfg *config.EtcdConfig) (client *clientv3.Client, err error) {
	//
	EtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: cfg.DialTimeout * time.Second,
	})
	if err != nil {
		fmt.Printf("\n init etcd failure, err: %s\n", err.Error())
		return
	}
	return EtcdClient, err
}

// WatchConf 监听配置修改
func WatchConf(key string, newProxyConfList chan<- []*EtcdProxyConfItem) {
	//
	for true {
		rch := EtcdClient.Watch(context.Background(), key)
		for wresp := range rch {
			for _, ev := range wresp.Events {
				// fmt.Printf("Type:%s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				var newConfList []*EtcdProxyConfItem
				if ev.Type != clientv3.EventTypeDelete {
					err := json.Unmarshal(ev.Kv.Value, &newConfList)
					if err != nil {
						fmt.Println("unmarshal new configuration failed,err:", err)
						return
					}
				}
				newProxyConfList <- newConfList
			}
		}
	}
}

// Set 设置
func Set(key string, value string) (*clientv3.PutResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := EtcdClient.Put(ctx, key, value)
	cancel()
	if err != nil {
		// handle error!
	}
	return resp, err
}

// Delete 清除某个值
func Delete(key string) (*clientv3.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := EtcdClient.Delete(ctx, key)
	cancel()
	if err != nil {
		// handle error!
	}
	return resp, err
}

// Get 设置
func Get(key string) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := EtcdClient.Get(ctx, key)
	cancel()
	if err != nil {
		// handle error!
	}
	return resp, err
}

// GetProxyConfList 获取代理配置列表
func GetProxyConfList(key string) (etcdConfigConf EtcdProxyConf) {
	//
	resp, err := Get(key)
	if err != nil {
		fmt.Printf("\n etcd get failure, err: %s \n", err.Error())
		return
	}
	if len(resp.Kvs) > 0 {
		err = json.Unmarshal(resp.Kvs[0].Value, &etcdConfigConf)
		if err != nil {
			fmt.Printf("\n etcd get success unmarshal failure, err: %s \n", err.Error())
			return
		}
	}
	return
}
