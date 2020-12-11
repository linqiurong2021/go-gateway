package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/linqiurong2021/go-gateway/config"
	"go.etcd.io/etcd/clientv3"
)

func TestEtcd(t *testing.T) {
	if err := config.Init("config.ini"); err != nil {
		fmt.Printf("load config from file falure !, err:%v\n", err)
		return
	}
	//
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Conf.Endpoints,
		DialTimeout: config.Conf.DialTimeout * time.Second,
	})
	//
	if err != nil {
		// handle error!
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := cli.Put(ctx, "sample", "sample_value")
	cancel()
	if err != nil {
		// handle error!
	}
	fmt.Printf("%#v\n", resp)
	fmt.Println("ok")
}
