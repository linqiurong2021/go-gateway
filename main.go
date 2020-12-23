package main

import (
	"fmt"

	"github.com/linqiurong2021/go-gateway/agent"
	"github.com/linqiurong2021/go-gateway/config"
)

func main() {
	//
	if err := config.Init("./config/config.ini"); err != nil {
		fmt.Printf("load config from file falure !, err:%v\n", err)
		return
	}

	// 实例化agent并启用
	// new(agent.Agent).Start()
	new(agent.Agent).StartWatch()
}
