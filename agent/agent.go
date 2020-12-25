package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/linqiurong2021/go-gateway/config"
	"github.com/linqiurong2021/go-gateway/etcd"
	"go.etcd.io/etcd/clientv3"
)

// Agent 代理服务
type Agent struct {
	ProxyConfList []*etcd.EtcdProxyConfItem
}

// etcdProxyConfChan 配置项通道
var etcdProxyConfChan chan []*etcd.EtcdProxyConfItem

func (a *Agent) print(config []*etcd.EtcdProxyConfItem) {
	fmt.Printf("get new Config %#v\n", config)
	for _, conf := range config {
		// 搜索(有可能会有多个 需要负载均衡)
		fmt.Printf("Agent:%#v\n", conf.URL)
	}
}

// ServeHTTP 服务代理
func (a *Agent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var remote *url.URL
	var hasRouter bool
	go a.GetNewConf(etcdProxyConfChan)
	for _, conf := range a.ProxyConfList {
		// 搜索(有可能会有多个 需要负载均衡)
		if strings.Contains(r.RequestURI, conf.URL) {
			// fmt.Printf("Agent:%#v\n", s.server)
			add := fmt.Sprintf("http://%s:%d", conf.Host, conf.Port)
			remote, _ = url.Parse(add) // 有可能https
			hasRouter = true
		}
	}

	if !hasRouter {
		fmt.Fprintf(w, "404 Not Found")
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

// GetProxyConfList 获取代理列表
func (a *Agent) GetProxyConfList() []*etcd.EtcdProxyConfItem {
	return etcd.GetProxyConfList(config.Conf.EtcdConfig.Key)
}

// Start 启动转发
func (a *Agent) Start() {
	// 初始化
	etcd.Init(config.Conf.EtcdConfig)
	etcdProxyConfChan = make(chan []*etcd.EtcdProxyConfItem, 1)
	addr := fmt.Sprintf(":%d", config.Conf.Port)
	// 先把默认的给
	etcdProxyConfChan <- a.GetProxyConfList()
	// 开启监听
	go a.Watch("/services", etcdProxyConfChan)
	err := http.ListenAndServe(addr, a)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

// Watch Watch
func (a *Agent) Watch(key string, newProxyConfList chan []*etcd.EtcdProxyConfItem) {
	for true {
		rch := etcd.EtcdClient.Watch(context.Background(), key)
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("Type:%s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				var newConfList []*etcd.EtcdProxyConfItem
				if ev.Type != clientv3.EventTypeDelete {
					err := json.Unmarshal(ev.Kv.Value, &newConfList)
					if err != nil {
						// fmt.Println("unmarshal new configuration failed,err:", err)
						return
					}
				}
				// fmt.Printf("Watch: %#v \n\n", newProxyConfList)
				newProxyConfList <- newConfList
			}
		}
	}
}

// GetNewConf GetNewConf
func (a *Agent) GetNewConf(newProxyConfList chan []*etcd.EtcdProxyConfItem) {
	select {
	case config := <-newProxyConfList:
		{
			fmt.Printf("GetNewConf%#v\n\n", config)
			a.ProxyConfList = config
		}
	}
}
