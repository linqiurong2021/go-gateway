package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/linqiurong2021/go-gateway/config"
	"github.com/linqiurong2021/go-gateway/etcd"
)

//

func TestInit(t *testing.T) {

	if err := config.Init("config.ini"); err != nil {
		fmt.Printf("load config from file falure !, err:%v\n", err)
		return
	}

}

func TestEtcdPutServices(t *testing.T) {

	if err := config.Init("config.ini"); err != nil {
		fmt.Printf("load config from file falure !, err:%v\n", err)
		return
	}
	etcd.Init(config.Conf.EtcdConfig)

	services := &etcd.EtcdProxyConf{
		{ID: 1, Name: "测试", Port: 8899, Alive: true, Host: "127.0.0.1", URL: "/service/register"},
		{ID: 2, Name: "测试", Port: 8899, Alive: true, Host: "127.0.0.1", URL: "/service/unregister"},
		{ID: 3, Name: "测试", Port: 8899, Alive: true, Host: "127.0.0.1", URL: "/api/service"},
		{ID: 4, Name: "测试", Port: 8899, Alive: true, Host: "127.0.0.1", URL: "/api2/service"},
	}

	bytes, err := json.Marshal(services)
	if err != nil {
		fmt.Printf("marshal err 444 %s", err.Error())
		return
	}
	value := string(bytes)
	//
	resp, err := etcd.Set("/services", value)
	if err != nil {
		fmt.Println("etcd set failure , err : ", err.Error())
		return
	}
	fmt.Printf("%v", resp)
	// assert.Equal(t, string(resp.PrevKv[0].Value), value)
}

func TestEtcdPut(t *testing.T) {

	if err := config.Init("config.ini"); err != nil {
		fmt.Printf("load config from file falure !, err:%v\n", err)
		return
	}
	etcd.Init(config.Conf.EtcdConfig)
	value := "world22"
	//
	resp, err := etcd.Set("Hello", value)
	if err != nil {
		fmt.Println("etcd set failure , err : ", err.Error())
		return
	}
	fmt.Printf("%v", resp)
	// assert.Equal(t, string(resp.PrevKv[0].Value), value)
}

func TestEtcdGet(t *testing.T) {

	if err := config.Init("config.ini"); err != nil {
		fmt.Printf("load config from file falure !, err:%v\n", err)
		return
	}
	etcd.Init(config.Conf.EtcdConfig)
	value := "world"
	//
	resp, err := etcd.Get("/services")
	if err != nil {
		fmt.Println("etcd set failure , err : ", err.Error())
		return
	}
	fmt.Printf("%v", resp)
	// 未获取到数据
	if resp.Kvs != nil {
		assert.Equal(t, string(resp.Kvs[0].Value), value)
	}
}

func TestEtcdClear(t *testing.T) {

	if err := config.Init("config.ini"); err != nil {
		fmt.Printf("load config from file falure !, err:%v\n", err)
		return
	}
	etcd.Init(config.Conf.EtcdConfig)
	//
	resp, err := etcd.Delete("/services")
	if err != nil {
		fmt.Println("etcd set failure , err 22 : ", err.Error())
		return
	}
	fmt.Printf("%v", resp)
	// assert.Equal(t, string(resp.Kvs[0].Value), value)
}
