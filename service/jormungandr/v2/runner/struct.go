package runner

import (
	"jormungandr/v2/proto/base"
)

type structRunner struct {
	// 这里可以添加Runner级别的状态/属性，例如缓存等
}

func (t *structRunner) Tick(s *base.Space) *base.Space {
	AsyncTick(
		s,
		func(e *base.Entity) {
			if e.Struct.Destroy{
				return
			}
			if e.Struct.Health <= 0{
				e.Struct.Destroy = true
				return
			}

			e.Struct.ShieldRecovery += e.Struct.Delta.ShieldRecoveryA

			e.Struct.Health += e.Struct.Delta.HealthA
			e.Struct.Struct += e.Struct.Delta.StructA
			e.Struct.Shield += e.Struct.ShieldRecovery
		},
	)
	return s
}

func NewStructRunner() *structRunner {
	return &structRunner{}
}
