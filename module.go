package module

import (
	"context"
	"sync"
)

type Module interface {
	Run(ctx context.Context) //当
}


type ModuleList struct {
	ctx context.Context
	lock sync.RWMutex
	waitGrop sync.WaitGroup
}


func (hub *ModuleList)RegisterModule(module Module)  {
	hub.lock.Lock()
	defer hub.lock.Unlock()

	if hub.ctx == nil {
		panic("注册模块调用顺序不正确")
	}
	ctx, _ := context.WithCancel(hub.ctx)
	hub.waitGrop.Add(1)
	go func() {
		module.Run(ctx)
		hub.waitGrop.Done()
	}()

}

func (hub *ModuleList) Run(ctx context.Context)  {
	hub.ctx = ctx
	<-ctx.Done()
	hub.waitGrop.Wait()
}


type ModuleHub struct {
	ModuleList
	cancel context.CancelFunc
}

func (hub *ModuleHub)Cancel()  {
	hub.Cancel()
	hub.ModuleList.waitGrop.Wait()
}

func NewModuleHub() *ModuleHub {
	ctx, f := context.WithCancel(context.TODO())
	hub := &ModuleHub{}
	hub.cancel = f
	go hub.Run(ctx)
	return hub
}