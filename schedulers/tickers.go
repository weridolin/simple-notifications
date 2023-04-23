package schedulers

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

type Ticker struct {
	PlatForm          string
	MaxSchedulerCount int
	ScheduLerCache    []*Scheduler
	Executor          *cron.Cron
}

func NewTicker(platform string, maxSchedulerCount int, schedulers []*Scheduler) *Ticker {
	t := &Ticker{platform, maxSchedulerCount, schedulers, cron.New()}
	return t
}

func (t *Ticker) Start() {
	fmt.Println("start ticker...")
	t.Executor.Start()
}

func (t *Ticker) AddScheduler(s *Scheduler) {
	fmt.Println("add and start scheduler...", s)
	t.Executor.AddFunc(s.Period.Cron, func() {
		fmt.Println("start run scheduler... platform", s.PlatForm, "ups", s.Ups, "period", s.Period.Cron, "status")
		s.Start()
		fmt.Println("run scheduler end...")
	})
}

func (t *Ticker) Stop() {
	t.Executor.Stop()
}

type TickerPool struct {
	MaxTickerCount   int
	SchedulerCache   map[int]*Scheduler
	TickerCache      map[string][]*Ticker
	RunningTicker    []*Ticker
	WaitingTicker    []*Ticker
	RunningScheduler []*Scheduler
	WaitingScheduler []*Scheduler
}

func NewTickerPool(maxTickerCount int) *TickerPool {
	return &TickerPool{
		maxTickerCount,
		make(map[int]*Scheduler),
		make(map[string][]*Ticker),
		make([]*Ticker, 0),
		make([]*Ticker, 0),
		make([]*Scheduler, 0),
		make([]*Scheduler, 0),
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

func (tp *TickerPool) SubmitScheduler(s *Scheduler) {
	if _, ok := tp.SchedulerCache[s.DBIndex]; !ok {
		tp.SchedulerCache[s.DBIndex] = s
	}

	if _, ok := tp.TickerCache[s.PlatForm]; !ok {
		t := &Ticker{s.PlatForm, 2, []*Scheduler{s}, cron.New()}
		tp.AddTicker(t)
		t.AddScheduler(s)
		tp.RunningScheduler = append(tp.RunningScheduler, s)

	} else {
		// 每个platform暂时最多对应两个ticker,每个ticker最多对应2个scheduler
		// 其他策略 //TODO
		if len(tp.TickerCache[s.PlatForm]) > 2 {
			tp.WaitingScheduler = append(tp.WaitingScheduler, s)
		} else {
			for _, t := range tp.TickerCache[s.PlatForm] {
				// 每个ticker最多对应2个scheduler，判断scheduler是否已经满了
				if len(t.ScheduLerCache) < t.MaxSchedulerCount {
					t.AddScheduler(s)
					tp.RunningScheduler = append(tp.RunningScheduler, s)
					return
				}
			}
			tp.WaitingScheduler = append(tp.WaitingScheduler, s)
		}
	}
}

func (tp *TickerPool) AdjustPoolSize() {
	// ticker.Executor.Stop(
	fmt.Println("adjust pool size...")
}
