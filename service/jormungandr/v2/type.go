package jormungandr

import (
	"jormungandr/v2/proto/base"
)

type Handler interface {
	// 输入的是一个Space的引用
	// 将输入复制一份
	// 为Space触发所有的Tick方法
	// 返回新的Space
	Tick(*base.Space) *base.Space

	// 一次性运行很多个Tick
	MultiTick(*base.Space, uint32) []*base.Space
}

type Runner interface {
	Tick(*base.Space) *base.Space
}
