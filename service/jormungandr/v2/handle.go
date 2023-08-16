package jormungandr

import (
	"jormungandr/v2/proto/base"
	"jormungandr/v2/runner"
	"sync"

	"google.golang.org/protobuf/proto"
)

// 处理器，对每次的宇宙进行处理
type handler struct {
	IsolateRunner   []runner.Runner // 独立的Runner，可任意并发执行
	AssociateRunner []runner.Runner // 有关联的Runner，需要顺序执行
}

// 对s运行一次全量的tick，修改元数据
func (hand *handler) UpdateTick(s *base.Space) {
	// 判断属性是否存在，如不存在则进行维护

	// 将这些属性当作可选的
	if s.Entity == nil{
		// 宇宙内源死亡，不再有实体意味着不再tick
		s.Entity = make(map[uint64]*base.Entity)
		return
	}

	// Isolate异步执行
	wg := new(sync.WaitGroup)
	for _, ir := range hand.IsolateRunner {
		wg.Add(1)
		go func(r runner.Runner) {
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
}

// 对src宇宙进行一次全量Tick，不修改src，返回新宇宙的指针
func (hand *handler) Tick(src *base.Space) *base.Space {
	// 这样做可以保持修改的Space都不是原本的，而是内部的，可以直接返回这个Space的指针
	// 这里空的默认值将被proto删除，并置为空指针
	s := proto.Clone(src).(*base.Space)

	hand.UpdateTick(s)
	return s
}

func (hand *handler) MultiTick(s *base.Space, times uint32) []*base.Space {
	var rtn []*base.Space
	for i := 0; i < int(times); i++ {
		s = hand.Tick(s)
		rtn = append(rtn, s)
	}
	return rtn
}

// 新建一个Handler，半动态地将Runner入系统
func NewHandler() *handler {
	return &handler{
		IsolateRunner: []runner.Runner{
			runner.NewTimeRunner(),
			runner.NewVeloRunner(),
			runner.NewStructRunner(),
			// TODO: 这里添加其他可独立运行没有依赖的Runner
			// 未添加的Runner将不会运行
		},
		AssociateRunner: []runner.Runner{},
		// TODO: 这里添加其他需要关联运行或互相依赖的Runner
	}
}
