/*
manager.go: 统一管理所有scheduler,对scheduler的操作都必须经过manager
每个manager会对应一个tickerPool,用于管理所有的ticker
*/
package scheduler

import (
	"context"
	"sync"

	// config "github.com/weridolin/simple-vedio-notifications/configs"
	"github.com/weridolin/simple-vedio-notifications/tools"
	"honnef.co/go/tools/config"
)

type SchedulerManager struct {
	Schedulers             map[int]*Scheduler // key为数据库ID
	PlatFormSchedulerCache map[string][]int   // 平台名称: scheduler id列表（数据库索引）
	Config                 *config.Config     // 配置信息
	lock                   sync.RWMutex
	StartAfterAdd          bool            // 是否默认添加后启动
	Ctx                    context.Context //上下文
	Key                    string          //唯一标识
}

func NewSchedulerManager(ctx context.Context, key string) *SchedulerManager {
	return &SchedulerManager{Ctx: ctx, Schedulers: make(map[int]*Scheduler), Key: key, PlatFormSchedulerCache: make(map[string][]int), StartAfterAdd: false}
}

func (sm *SchedulerManager) AddScheduler(s *Scheduler, startAtOnce bool) (*SchedulerManager, error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	// scheduler缓存
	_, ok := sm.Schedulers[s.DBIndex]
	if ok {
		return sm, tools.SchedulerIsExistError
	}
	sm.Schedulers[s.DBIndex] = s
	// 平台缓存
	_, ok = sm.PlatFormSchedulerCache[s.PlatForm]
	if !ok {
		sm.PlatFormSchedulerCache[s.PlatForm] = []int{s.DBIndex}
	} else {
		sm.PlatFormSchedulerCache[s.PlatForm] = append(sm.PlatFormSchedulerCache[s.PlatForm], s.DBIndex)
	}

	if startAtOnce {
		tp := sm.Ctx.Value("tp").(*TickerPool)
		sm.lock.RLock()
		defer sm.lock.RUnlock()
		tp.SubmitScheduler(s)
	}
	return sm, nil
}

func (sm *SchedulerManager) RemoveScheduler(s *Scheduler) (*SchedulerManager, error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	// scheduler缓存移除
	_, ok := sm.Schedulers[s.DBIndex]
	if !ok {
		return sm, tools.SchedulerIsExistError
	} else {
		delete(sm.Schedulers, s.DBIndex)
	}
	// 平台缓存移除
	_, ok = sm.PlatFormSchedulerCache[s.PlatForm]
	if !ok {
		return sm, tools.SchedulerIsExistError
	} else {
		for i, v := range sm.PlatFormSchedulerCache[s.PlatForm] {
			if v == s.DBIndex {
				sm.PlatFormSchedulerCache[s.PlatForm] = append(sm.PlatFormSchedulerCache[s.PlatForm][:i], sm.PlatFormSchedulerCache[s.PlatForm][i+1:]...)
				break
			}
		}
	}

	if s.Status == 1 {
		s.Stop()
	}
	return sm, nil
}

func (sm *SchedulerManager) StartAll() (*SchedulerManager, error) {
	tp := sm.Ctx.Value("tp").(*TickerPool)
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	for _, s := range sm.Schedulers {
		tp.SubmitScheduler(s)
	}
	return sm, nil
}

func (sm *SchedulerManager) StopAll() (*SchedulerManager, error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	for _, s := range sm.Schedulers {
		s.Stop()
	}
	return sm, nil
}
