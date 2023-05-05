/*
	ticker是scheduler的执行者.每个ticker对应一个cron,可以同时运行多个scheduler.ticker的最小运行单位为scheduler.每次一个scheduler运行完毕,调整tickerPool的大小
	tickerPool是对所有ticker的管理，采用的类似线程池的实现方式和策略,这里操作的粒度精确到ticker,而不是scheduler

	## TODO逻辑修改
	每个ticker会对应一个platform。只能运行属于同一个platform的scheduler
	当ticker达到pool的最大数量限制，且所有ticker的scheduler达到最大数量限制。会暂时先把scheduler加入到tickerPool的等待队列中，等待ticker空闲后再执行
	当初ticker中的scheduler都停止或者出错时,ticker会去判断tickerPool中是否有属于同个platform的等待的scheduler,如果有则会把等待的scheduler加入到ticker中执行，
	如没有，则会再去遍历tp的等待队列，拿到第一个platform不同且该platform还未达到上限的scheduler,添加到该ticker并且作为该platform下的ticker运行


*/
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
	tp                *TickerPool
}

func NewTicker(platform string, maxSchedulerCount int, schedulers []*Scheduler) *Ticker {
	t := &Ticker{platform, maxSchedulerCount, schedulers, cron.New(), nil}
	return t
}

func (t *Ticker) Start() {
	// ch := make(chan bool)
	// var wg sync.WaitGroup
	for _, s := range t.ScheduLerCache {
		// wg.Add(1)
		fmt.Println("add and start scheduler...", s)
		t.Executor.AddFunc(s.Period.Cron, func() {
			fmt.Println("start run scheduler... platform", s.PlatForm, "period", s.Period.Cron, "status")
			s.Start()
			fmt.Println("run scheduler end...")
			// wg.Done()
		})
		s.ticker = t
		s.Status = 1
	}
	fmt.Println("start ticker...")
	t.Executor.Start()
	// wg.Wait()
	fmt.Println("ticker run finish...begin to adjust pool ")
	t.tp.AdjustPoolSize(t)
	// finish := <-ch
	// if finish {
	// 	fmt.Println("ticker run finish...")
	// }
}

func (t *Ticker) AddScheduler(s *Scheduler) {
	fmt.Println("add and start scheduler...", s)
	t.Executor.AddFunc(s.Period.Cron, func() {
		fmt.Println("start run scheduler... platform", s.PlatForm, "period", s.Period.Cron, "status")
		s.Start()
		fmt.Println("run scheduler end...")
	})
	s.ticker = t
	s.Status = 1
}

func (t *Ticker) Stop() {
	for _, s := range t.ScheduLerCache {
		s.Status = 0
		s.ticker = nil
	}
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
		fmt.Println("add scheduler to cache...")
		tp.SchedulerCache[s.DBIndex] = s
	}

	if _, ok := tp.TickerCache[s.PlatForm]; !ok {
		fmt.Println("create a new ticker to pool...")
		t := &Ticker{s.PlatForm, 2, []*Scheduler{s}, cron.New(), tp}
		tp.AddTicker(t)
		tp.RunningScheduler = append(tp.RunningScheduler, s)

	} else {
		// 每个platform暂时最多对应两个ticker,每个ticker最多对应2个scheduler
		// 其他策略 //TODO
		if len(tp.TickerCache[s.PlatForm]) > 2 {
			fmt.Println("add scheduler to waiting scheduler...")
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
			fmt.Println("ticker running queue is full...  add scheduler to waiting scheduler...")
			tp.WaitingScheduler = append(tp.WaitingScheduler, s)
		}
	}
}

func (tp *TickerPool) AdjustPoolSize(ticker *Ticker) {
	fmt.Println("adjust pool size...")
	// 先从waitingScheduler中取出等待执行的scheduler放入该ticker执行
	for i, s := range tp.WaitingScheduler {
		if s.PlatForm == ticker.PlatForm {
			ticker.AddScheduler(s)
			tp.WaitingScheduler = append(tp.WaitingScheduler[:i], tp.WaitingScheduler[i+1:]...)
		}
	}
	// 没有对应platform的task 删除该ticker
	ticker.Stop()
}
