package jormungandr

import (
	"google.golang.org/protobuf/proto"
	"jormungandr/v2/proto/base"
	"jormungandr/v2/runner"
)

// 处理器，对每次的宇宙进行处理
type handler struct {
	RunnerList []Runner  // Runner可以是有状态的变量
}

// 运行一次，返回新宇宙的指针
func (hand *handler) Tick(src *base.Space) *base.Space {
	// 这样做可以保持修改的Space都不是原本的，而是内部的，可以直接返回这个Space的指针
	s := proto.Clone(src).(*base.Space)

	rs := hand.RunnerList

	for _, r := range rs {
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
		RunnerList: []Runner{
			runner.NewTimeRunner(),
			runner.NewVeloRunner(),
			// TODO: 这里添加其他的Runner，未添加的Runner将不会运行
		},
	}
}
