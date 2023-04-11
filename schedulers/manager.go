package schedulers

import (
	"context"
	"sync"

	config "github.com/weridolin/simple-vedio-notifications/configs"
)

type SchedulerManager struct {
	Schedulers map[int]*Scheduler // key为数据库ID
	Config     *config.Config
	lock       sync.RWMutex
	start      bool
	Ctx        context.Context
}

func NewSchedulerManager(ctx context.Context) *SchedulerManager {
	return &SchedulerManager{Ctx: ctx, Schedulers: make(map[int]*Scheduler)}
}

func (sm *SchedulerManager) AddScheduler(s *Scheduler) (*SchedulerManager, error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	_, ok := sm.Schedulers[s.DBIndex]
	if ok {
		return sm, SchedulerIsExistError
	}
	sm.Schedulers[s.DBIndex] = s
	if sm.start {
		s.Start()
	}
	return sm, nil
}

func (sm *SchedulerManager) StopScheduler(s *Scheduler) (*SchedulerManager, error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	_, ok := sm.Schedulers[s.DBIndex]
	if !ok {
		return sm, SchedulerIsExistError
	}
	if s.Status == 1 {
		s.Stop()
	}
	delete(sm.Schedulers, s.DBIndex)
	return sm, nil
}

func (sm *SchedulerManager) StartAll() (*SchedulerManager, error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	// s, ok := sm.Schedulers[s.DBIndex]
	for _, s := range sm.Schedulers {
		s.Start()
	}
	return sm, nil
}

func (sm *SchedulerManager) StopAll() (*SchedulerManager, error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	// s, ok := sm.Schedulers[s.DBIndex]
	for _, s := range sm.Schedulers {
		s.Stop()
	}
	return sm, nil
}
