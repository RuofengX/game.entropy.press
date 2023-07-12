package runner

import (
	"jormungandr/v2/proto/base"
)

type VelocityRunner struct {}

func (r *VelocityRunner) tick(ent *base.Entity) *base.Entity {
	ent.Velo.X = ent.Velo.X + ent.Velo.XV
	ent.Velo.Y = ent.Velo.Y + ent.Velo.YV
	ent.Velo.XV = ent.Velo.XV + ent.VeloD.XA
	ent.Velo.YV = ent.Velo.YV + ent.VeloD.YA
	ent.VeloD.XA = 0
	ent.VeloD.YA = 0
	return ent
}
