package jormungandr

import (
	"jormungandr/v2/proto/base"
	"jormungandr/v2/runner"
	"sync"

	"google.golang.org/protobuf/proto"
)

// 处理器，对每次的宇宙进行处理
type handler struct {
	IsolateRunner   []Runner // 独立的Runner，可任意并发执行
	AssociateRunner []Runner // 有关联的Runner，需要顺序执行
}

// 运行一次，返回新宇宙的指针
func (hand *handler) Tick(src *base.Space) *base.Space {
	// 这样做可以保持修改的Space都不是原本的，而是内部的，可以直接返回这个Space的指针
	s := proto.Clone(src).(*base.Space)

	// Isolate异步执行
	wg := new(sync.WaitGroup)
	for _, ir := range hand.IsolateRunner {
		wg.Add(1)
		go func(r Runner) {
			r.Tick(s)
			wg.Done()
		}(ir)
	}
	wg.Wait()
	// Associate可能也会依赖Isolate的类型
	// 所以这里直接等待Isolate全部完成

	// Associate顺序执行
	for _, r := range hand.AssociateRunner {
		r.Tick(s)
	}

	return s
}

func (hand handler) MultiTick(s *base.Space, times uint32) []*base.Space {
	var rtn []*base.Space
	for i := 0; i < int(times); i++ {
		s = hand.Tick(s)
		rtn = append(rtn, s)
	}
	return rtn
}

// 新建一个Handler，半动态地将Runner入系统
func NewHandler() Handler {
	return &handler{
		IsolateRunner: []Runner{
			runner.NewTimeRunner(),
			runner.NewVeloRunner(),
			// TODO: 这里添加其他可独立运行没有依赖的Runner
			// 未添加的Runner将不会运行
		},
		AssociateRunner: []Runner{},
		// TODO: 这里添加其他需要关联运行或互相依赖的Runner
	}
}
