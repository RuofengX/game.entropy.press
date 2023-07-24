package runner

import (
	"jormungandr/v2/proto/base"
)

type veloRunner struct {
	// 这里可以添加Runner级别的状态/属性，例如缓存等
}

func (t *veloRunner) Tick(s *base.Space) *base.Space {
	AsyncTick(
		s,
		func(ent *base.Entity) {
			ent.Velo.XV += ent.Velo.Delta.XA
			ent.Velo.YV += ent.Velo.Delta.YA
			ent.Velo.X += ent.Velo.XV
			ent.Velo.Y += ent.Velo.YV

			ent.Velo.Delta.XA = 0
			ent.Velo.Delta.YA = 0

		},
	)
	return s
}

func NewVeloRunner() *veloRunner {
	return &veloRunner{}
}
