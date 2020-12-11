package etcd

import (
	"context"
	"time"

	"github.com/linqiurong2021/go-gateway/config"
	"go.etcd.io/etcd/clientv3"
)

// Client Client
type Client struct {
	Client      *clientv3.Client
	DialTimeout time.Duration
}

// Init Init
func (c *Client) Init(cfg *config.EtcdConfig) (*clientv3.Client, error) {
	//
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: cfg.DialTimeout * time.Second,
	})
	c.Client = cli
	c.DialTimeout = cfg.DialTimeout
	return cli, err
}

// Set 设置
func (c *Client) Set(key string, value string) (*clientv3.PutResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := c.Client.Put(ctx, key, value)
	cancel()
	if err != nil {
		// handle error!
	}
	return resp, err
}

// Get 设置
func (c *Client) Get(key string) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := c.Client.Get(ctx, key)
	cancel()
	if err != nil {
		// handle error!
	}
	return resp, err
}
