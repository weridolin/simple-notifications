package schedulers

import "fmt"

type Period struct {
}

type Scheduler struct {
	Period   Period
	PlatForm string
	Ups      []string
	Status   int8 // 0 停止 1 启动  2 暂停
	DBIndex  int  //唯一索引
}

func (s *Scheduler) Start() {
	s.Status = 1
	fmt.Println("start scheduler...", s)
}

func (s *Scheduler) Stop() {
	s.Status = 0
	fmt.Println("stop scheduler...", s)
}

func (s *Scheduler) AddUp() {
	fmt.Println("add up...", s)
}

func (s *Scheduler) RemoveUp() {
	fmt.Println("remove up...", s)
}

func (s *Scheduler) Delete() {
	s.Status = 0
	fmt.Println("delete scheduler...", s)
}
