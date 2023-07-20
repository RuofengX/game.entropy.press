package jormungandr

import (
	"google.golang.org/protobuf/proto"
	"jormungandr/v2/proto/base"
	"jormungandr/v2/runner"
)

// 编译时检查
var _ Handler = new(DefaultHandler)

type DefaultHandler struct {
	// 有状态的变量
	RunnerList []Runner
	TickDone   chan bool
}

func (hand *DefaultHandler) Tick(src *base.Space) *base.Space {
	// 这样做可以保持修改的Space都不是原本的，而是内部的，可以直接返回这个Space的指针
	s := proto.Clone(src).(*base.Space)

	rs := hand.RunnerList

	for _, r := range rs {
		r.Tick(s)
	}

	return s
}

func (hand DefaultHandler) MultiTick(s *base.Space, times uint32) []*base.Space {
	var rtn []*base.Space
	for i := 0; i < int(times); i++ {
		s = hand.Tick(s)
		rtn = append(rtn, s)
	}
	return rtn
}

func NewHandler() Handler {
	return &DefaultHandler{
		RunnerList: []Runner{
			runner.NewTimeRunner(),
			// 这里添加其他的Runner
		},
		TickDone: make(chan bool),
	}
}
