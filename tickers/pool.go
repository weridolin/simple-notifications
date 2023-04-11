package tickers

import (
	"fmt"

	"github.com/robfig/cron/v3"
	schedulers "github.com/weridolin/simple-vedio-notifications/schedulers"
)

type Ticker struct {
	PlatForm          string
	MaxSchedulerCount int
	ScheduLerCache    []*schedulers.Scheduler
	Executor          *cron.Cron
}

func (t *Ticker) Start() {
	t.Executor.Start()
}

func (t *Ticker) AddScheduler(s *schedulers.Scheduler) {
	// t.Executor.Stop()
}

func (t *Ticker) Stop() {
	t.Executor.Stop()
}

type TickerPool struct {
	MaxTickerCount int
	SchedulerCache map[int]*schedulers.Scheduler
	TickerCache    map[string][]*Ticker
	RunningTicker  []*Ticker
	WaitingTicker  []*Ticker
}

func NewTickerPool(maxTickerCount int) *TickerPool {
	return &TickerPool{
		maxTickerCount,
		make(map[int]*schedulers.Scheduler),
		make(map[string][]*Ticker),
		make([]*Ticker, 0),
		make([]*Ticker, 0),
	}
}

func (tp *TickerPool) AddTicker(ticker *Ticker) {
	fmt.Println("add ticker...")
	if len(tp.RunningTicker) >= tp.MaxTickerCount {
		tp.WaitingTicker = append(tp.WaitingTicker, ticker)
	} else {
		tp.RunningTicker = append(tp.RunningTicker, ticker)
		ticker.Start()
	}
	if _, ok := tp.TickerCache[ticker.PlatForm]; !ok {
		tp.TickerCache[ticker.PlatForm] = make([]*Ticker, 0)
	}
	tp.TickerCache[ticker.PlatForm] = append(tp.TickerCache[ticker.PlatForm], ticker)
}

func (tp *TickerPool) RemoveTicker(ticker *Ticker) {
	if _, ok := tp.TickerCache[ticker.PlatForm]; ok {
		for i, t := range tp.TickerCache[ticker.PlatForm] {
			if t == ticker {
				tp.TickerCache[ticker.PlatForm] = append(tp.TickerCache[ticker.PlatForm][:i], tp.TickerCache[ticker.PlatForm][i+1:]...)
			}
		}
	}
	// 从RunningTicker移除
	for i, t := range tp.RunningTicker {
		if t == ticker {
			tp.RunningTicker = append(tp.RunningTicker[:i], tp.RunningTicker[i+1:]...)
			ticker.Stop()
			return
		}
	}
	//	从RunningTicker移除
	for i, t := range tp.WaitingTicker {
		if t == ticker {
			tp.WaitingTicker = append(tp.WaitingTicker[:i], tp.WaitingTicker[i+1:]...)
			return
		}
	}
	fmt.Println("remove ticker...")
}

func (tp *TickerPool) SubmitScheduler(s *schedulers.Scheduler) {
	if _, ok := tp.SchedulerCache[s.DBIndex]; !ok {
		tp.SchedulerCache[s.DBIndex] = s
	}

	if _, ok := tp.TickerCache[s.PlatForm]; !ok {
		t := &Ticker{s.PlatForm, 10, []*schedulers.Scheduler{s}, cron.New()}
		tp.AddTicker(t)
	} else {
		// 其他策略 //TODO
		tp.TickerCache[s.PlatForm][0].AddScheduler(s)
	}
}

func (tp *TickerPool) AdjustPoolSize() {
	// ticker.Executor.Stop(
	fmt.Println("adjust pool size...")
}
