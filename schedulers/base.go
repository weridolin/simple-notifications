package main

import (
	"os"

	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/yaml.v3"
)

type Storage struct {
	StorageType   string                 `json:"StorageType" yaml:"StorageType"` // 存储类型
	StorageParams map[string]interface{} `json:"StorageParams" yaml:"StorageParams"`
}

type SchedulerConfig struct {
	DefaultTickerMaxSchedulerCount int    `json:"DefaultTickerMaxSchedulerCount" yaml:"DefaultTickerMaxSchedulerCount"` // 执行池最大执行scheduler数量`
	DefaultMaxTickerCount          int    `json:"DefaultMaxTickerCount" yaml:"DefaultMaxTickerCount"`                   // 执行池最大执行ticker数量
	RabbitMQUri                    string `json:"RabbitMQUri" yaml:"RabbitMQUri"`

	EmailConsumerCount          int    `json:"EmailConsumerCount" yaml:"EmailConsumerCount"`                   // EmailConsumer的数量
	EmailMessageAckTimeOut      int    `json:"EmailMessageAckTimeOut" yaml:"EmailMessageAckTimeOut"`           // EmailConsumer消息ack超市时间
	EmailMessageDlxExchangeName string `json:"EmailMessageDlxExchangeName" yaml:"EmailMessageDlxExchangeName"` // EmailConsumer消息死信交换机名称
	EmailMessageDlxQueueName    string `json:"EmailMessageDlxQueueName" yaml:"EmailMessageDlxQueueName"`       // EmailConsumer消息死信队列名称
	EmailMessageExchangeName    string `json:"EmailMessageExchangeName" yaml:"EmailMessageExchangeName"`       // EmailConsumer消息交换机名称
	EmailMessageQueueName       string `json:"EmailMessageQueueName" yaml:"EmailMessageQueueName"`             // EmailConsumer消息队列名称

	// StorageFileRelativePath     string // 结果存储文件路径
	Storage Storage `json:"Storage" yaml:"Storage"`
}

func (s *SchedulerConfig) FromYamlFile(configPath string) *SchedulerConfig {
	dataBytes, err := os.ReadFile(configFile)
	if err != nil {
		logx.Error("读取文件失败：", err)
		return nil
	}
	var ConfigInstance = &SchedulerConfig{}
	err = yaml.Unmarshal(dataBytes, ConfigInstance)
	if err != nil {
		logx.Error("yaml反序列化失败:", err)
		return nil
	}
	return ConfigInstance
}

type Scheduler struct {
	Period   string        //定时周期（cron格式）
	PlatForm string        //	scheduler对应的平台名称
	Tasks    []interface{} // scheduler对应的task列表
	Status   int8          // 0 停止 1 启动  2 暂停
	DBIndex  int           //唯一索引
	ticker   *Ticker       //绑定的ticker
}

type Task struct {
	PlatForm    string //	task对应的平台名称
	DBIndex     int    //唯一索引
	Name        string
	Description string
	Ups         map[string]interface{} // task对应的ups
	// EmailNotifiers []any  //默认自带邮件通知
}

type ITask interface {
	Run()
	GetUpInfo()
	Stop()
}
