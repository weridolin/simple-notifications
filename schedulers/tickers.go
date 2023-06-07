/*
ticker是scheduler的执行者.每个ticker对应一个cron,可以同时运行多个scheduler.ticker的最小运行单位为scheduler.每次一个scheduler运行完毕,调整tickerPool的大小
tickerPool是对所有ticker的管理，采用的类似线程池的实现方式和策略,这里操作的粒度精确到ticker,而不是scheduler

## TODO逻辑修改
每个ticker会对应一个platform。只能运行属于同一个platform的scheduler
当ticker达到pool的最大数量限制，且所有ticker的scheduler达到最大数量限制。会暂时先把scheduler加入到tickerPool的等待队列中，等待ticker空闲后再执行
当初ticker中的scheduler都停止或者出错时,ticker会去判断tickerPool中是否有属于同个platform的等待的scheduler,如果有则会把等待的scheduler加入到ticker中执行，
如没有，则会再去遍历tp的等待队列，拿到第一个platform不同且该platform还未达到上限的scheduler,添加到该ticker并且作为该platform下的ticker运行
*/
package main

import (
	"context"

	"github.com/robfig/cron/v3"
	"github.com/weridolin/simple-vedio-notifications/storage"
	"github.com/weridolin/simple-vedio-notifications/tools"
	"github.com/zeromicro/go-zero/core/logx"
)

type Ticker struct {
	PlatForm          string       // ticker所属平台，每个ticker只能对应一个平台
	MaxSchedulerCount int          // ticker 对应的scheduler最大监听数量
	ScheduLerCache    []*Scheduler // ticker里面监听的scheduler
	Executor          *cron.Cron   // ticker 对应的执行器
	tp                *TickerPool  // ticker绑定的ticker pool
	id                string       // ticker的唯一标识
}

func NewTicker(platform string, maxSchedulerCount int, schedulers []*Scheduler) *Ticker {
	t := &Ticker{platform, maxSchedulerCount, schedulers, cron.New(), nil, tools.GetUUID()}
	return t
}

func (t *Ticker) Start() {
	// ch := make(chan bool)
	// var wg sync.WaitGroup
	for _, s := range t.ScheduLerCache {
		t.Executor.AddFunc(s.Period, func() {
			logx.Info("start run scheduler... platform", s.PlatForm, "period", s.Period, "ticker id", t.id)
			s.Start()
			logx.Info("run scheduler end...", s.PlatForm, "period", s.Period)
			// wg.Done()
		})
		s.ticker = t
		s.Status = 1
	}
	logx.Info("start ticker,ticker scheduler count -> ", len(t.ScheduLerCache))
	t.Executor.Start()
	// wg.Wait()
	// logger.Println("ticker run finish...begin to adjust pool ")
	// t.tp.AdjustPoolSize(t)
	// finish := <-ch
	// if finish {
	// 	fmt.Println("ticker run finish...")
	// }
}

func (t *Ticker) AddScheduler(s *Scheduler) {
	logx.Info("add a scheduler to ticker and start")
	t.Executor.AddFunc(s.Period, func() {
		logx.Info("start run scheduler... platform", s.PlatForm, "period", s.Period, "ticker id", t.id)
		s.Start()
		logx.Info("run scheduler end...", s.PlatForm, "period", s.Period)
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
	MaxTickerCount   int //这里是指每个platform对应的最大ticker数量
	SchedulerCache   map[int]*Scheduler
	TickerCache      map[string][]*Ticker
	RunningTicker    []*Ticker
	WaitingTicker    []*Ticker
	RunningScheduler []*Scheduler
	WaitingScheduler []*Scheduler
	ID               string
	Storage          storage.StorageInterface
	AppConfig        SchedulerConfig
}

func NewTickerPool(config SchedulerConfig) *TickerPool {
	ctx := context.TODO()
	return &TickerPool{
		config.DefaultMaxTickerCount,
		make(map[int]*Scheduler),
		make(map[string][]*Ticker),
		make([]*Ticker, 0),
		make([]*Ticker, 0),
		make([]*Scheduler, 0),
		make([]*Scheduler, 0),
		tools.GetUUID(),
		storage.NewStorage(ctx, config.Storage.StorageType, config.Storage.StorageParams),
		config,
	}
}

func (tp *TickerPool) AddTicker(ticker *Ticker) {
	logx.Info("add ticker to pool ... pool id -> ", tp.ID, "ticker id -> ", ticker.id)
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
	logx.Info("remove ticker id", ticker.id, "from pool id", tp.ID)
}

func (tp *TickerPool) SubmitScheduler(s *Scheduler) {
	if _, ok := tp.SchedulerCache[s.DBIndex]; !ok {
		logx.Info("add scheduler to cache...")
		tp.SchedulerCache[s.DBIndex] = s
	}

	if _, ok := tp.TickerCache[s.PlatForm]; !ok {
		if len(tp.TickerCache) < tp.MaxTickerCount {
			logx.Info("create a new ticker to pool...")
			t := &Ticker{s.PlatForm, tp.AppConfig.DefaultTickerMaxSchedulerCount, []*Scheduler{s}, cron.New(), tp, tools.GetUUID()}
			tp.AddTicker(t)
			tp.RunningScheduler = append(tp.RunningScheduler, s)
		} else {
			tp.WaitingScheduler = append(tp.WaitingScheduler, s)
		}

	} else {
		// 先判断tickerPool中的ticker是否已经满了，是的话scheduler添加到等待队列
		if len(tp.TickerCache[s.PlatForm]) > tp.MaxTickerCount {
			logx.Info("ticker count is full , add scheduler to waiting scheduler...")
			tp.WaitingScheduler = append(tp.WaitingScheduler, s)
		} else {
			// 在判断对应的平台的所有的ticker中，是否有ticker的scheduler数量未达到最大值，是的话添加到该ticker中执行，都没有的话添加到等待队列
			for _, t := range tp.TickerCache[s.PlatForm] {
				if len(t.ScheduLerCache) < t.MaxSchedulerCount {
					t.AddScheduler(s)
					tp.RunningScheduler = append(tp.RunningScheduler, s)
					return
				}
			}
			logx.Info("ticker running queue is full...  add scheduler to waiting scheduler...")
			tp.WaitingScheduler = append(tp.WaitingScheduler, s)
		}
	}
}

func (tp *TickerPool) AdjustPoolSize(ticker *Ticker) {
	logx.Info("adjust pool size...")
	// 先从waitingScheduler中取出等待执行的scheduler放入该ticker执行
	for i, s := range tp.WaitingScheduler {
		if s.PlatForm == ticker.PlatForm {
			ticker.AddScheduler(s)
			tp.WaitingScheduler = append(tp.WaitingScheduler[:i], tp.WaitingScheduler[i+1:]...)
			return
		}
	}
	// 没有对应platform的task 删除该ticker
	ticker.Stop()
}
