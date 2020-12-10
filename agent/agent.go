package agent

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/linqiurong2021/go-gateway/config"
	"github.com/linqiurong2021/go-gateway/logic"
	"github.com/linqiurong2021/go-gateway/model"
)

// Agent 代理服务
type Agent struct {
	server []*model.Service // 需要动态添加
}

// ServeHTTP 服务代理
func (s *Agent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//
	var remote *url.URL
	var hasRouter bool
	// fmt.Println("ServeHTTP")
	for _, server := range s.server {
		// 搜索
		if strings.Contains(r.RequestURI, server.URL) {
			fmt.Printf("Agent:%#v\n", s.server)
			remote, _ = url.Parse("http://" + server.Host + ":" + server.Port) // 有可能https
			hasRouter = true
		}
	}
	// 需求判断未找到时提示路由未找到
	if !hasRouter {
		fmt.Fprintf(w, "404 Not Found")
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

// Start 启动转发
func (s *Agent) Start() {
	var serviceLogic = new(logic.Service)
	// 服务列表(数据库中取出)
	serviceList, err := serviceLogic.GetAllService()
	if err != nil {
		fmt.Sprintln(" Error: ", err.Error())
		return
	}
	fmt.Printf("%#v\n", serviceList)
	// 服务列表
	service := &Agent{
		server: serviceList,
	}
	addr := fmt.Sprintf(":%d", config.Conf.Port)
	err = http.ListenAndServe(addr, service)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
