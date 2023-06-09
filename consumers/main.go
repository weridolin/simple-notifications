package main

import (
	"flag"
	"fmt"

	"github.com/weridolin/simple-vedio-notifications/monitor"
)

var configFile string = ""

func init() {
	flag.StringVar(&configFile, "f", "", "consumer配置文件路径")
	flag.Parse()
}

func main() {
	//这里要主要的是，如果是在main函数中，运行要使用 go run ./directory
	config := &ConsumerConfig{}
	if configFile == "" {
		fmt.Println("请使用 -f 指定配置文件目录")
		return
	}
	config = config.FromYamlFile(configFile)

	// 启动监控
	fmt.Println("consumer监控启动...", config.Prometheus.Host, ":", config.Prometheus.Port, config.Prometheus.Path)
	monitor.Start(config.Prometheus.Host, config.Prometheus.Port, config.Prometheus.Path)
	// emailConsumer := NewEmailConsumer(tools.GetUUID(), *config)
	// emailConsumer.Start()
}
