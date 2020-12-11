package test

import (
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

	var client = new(etcd.Client)
	client.Init(config.Conf.EtcdConfig)
	put, err := client.Set("hello", "world")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v", put)
	get, err := client.Get("hello")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%#v", get)
	assert.Equal(t, "world", string(get.Kvs[0].Value))
}
