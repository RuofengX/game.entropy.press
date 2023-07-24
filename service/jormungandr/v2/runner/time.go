package runner

import (
	"jormungandr/v2/proto/base"
)

type timeRunner struct {
	// 这里可以添加Runner级别的状态/属性，例如缓存等
}

func (t *timeRunner) Tick(s *base.Space) *base.Space {
	ents := s.Entity
	for _, ent := range ents {
		ent.Time.Speed += ent.Time.Delta.TimeA
		ent.Time.Age += ent.Time.Speed
		ent.Time.Delta.TimeA = 0
	}
	return s
}

func NewTimeRunner() *timeRunner {
	return &timeRunner{}
}
