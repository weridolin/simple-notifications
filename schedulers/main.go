package main

import (
	"flag"
	"fmt"
)

var configFile string = ""

func init() {
	flag.StringVar(&configFile, "f", "", "scheduler配置文件路径")
	flag.Parse()
}

func main() {
	//这里要主要的是，如果是在main函数中，运行要使用 go run ./directory
	config := &SchedulerConfig{}
	if configFile == "" {
		fmt.Println("请使用 -f 指定配置文件目录")
		return
	}
	config = config.FromYamlFile(configFile)
	fmt.Println(config, ">>>")
}
