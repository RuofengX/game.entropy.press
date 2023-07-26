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
			// 可选字段
			// 包含可选字段是常态，没有可选字段是特例
			if e.Structure == nil{
				return
			}
			if e.Structure.Destroy{
				return
			}
			if e.Structure.Health <= 0{
				e.Structure.Destroy = true
				return
			}

			e.Structure.ShieldRecovery += e.Structure.Delta.ShieldRecoveryA

			e.Structure.Health += e.Structure.Delta.HealthA
			e.Structure.Structure += e.Structure.Delta.StructureA
			e.Structure.Shield += e.Structure.ShieldRecovery
		},
	)
	return s
}

func NewStructRunner() *structRunner {
	return &structRunner{}
}
