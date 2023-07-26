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
			if e.Structure == nil {
				return
			}
			if e.Structure.Destroy {
				return
			}
			if e.Structure.Health <= 0 {
				e.Structure.Destroy = true
				return
			}

			e.Structure.Shield += e.Structure.ShieldRecovery // 护盾自回复

			e.Structure.Health += e.Structure.Delta.HealthA // 生命值回复
			e.Structure.Structure += e.Structure.Delta.StructureA // 结构回复

			// 将Delta归零
			e.Structure.Delta.HealthA = 0
			e.Structure.Delta.StructureA = 0

		},
	)
	return s
}

func NewStructRunner() *structRunner {
	return &structRunner{}
}
