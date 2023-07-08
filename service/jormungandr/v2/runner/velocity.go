package runner

import (
	"jormungandr/v2/proto/base"
	velo "jormungandr/v2/proto/velo"
)

type VelocityRunner struct {
	Tickable
}

func (r *VelocityRunner) tick(ent *base.Entity) *base.Entity {
	vel := ent.GetVelo()
	rtn := &(base.Entity{
		ID: ent.ID,
		Velo: &velo.Fragment{
			X:  vel.X + vel.XV,
			Y:  vel.Y + vel.YV,
			XV: vel.XV + vel.Delta.XA,
			YV: vel.YV + vel.Delta.YA,
			Delta: &velo.Delta{
				XA: 0,
				YA: 0,
			},
		},
	})
	return rtn
}
