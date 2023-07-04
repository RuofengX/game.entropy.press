package runner

import (
	"jormungandr/v2/errors"
	ctum "jormungandr/v2/proto/continuum"
	meta "jormungandr/v2/proto/meta"
	velo "jormungandr/v2/proto/velo"
)

type VelocityRunner struct{}

func (r *VelocityRunner) tick(ent *meta.Entity) *meta.Entity {
	vel := ent.GetVelo()
	rtn := &(meta.Entity{
		ID: ent.ID,
		Velo: &velo.Fragment{
			X:  vel.X + vel.XV,
			Y:  vel.Y + vel.YV,
			XV: vel.XV + vel.Delta.XA,
			YV: vel.YV + vel.Delta.YA,
		},
	})
	return rtn
}

func (r *VelocityRunner) fieldTick(space *meta.Field) *meta.Field {
	rtn := new(meta.Field)
	for _, ent := range space.Entity {
		rtn.Entity = append(rtn.Entity, r.tick(ent))
	}
	return rtn
}

func (r *VelocityRunner) Handle(req *ctum.Request) (*ctum.Result, error) {
	epoch := 1
	limit := int(req.NestTick)
	velo_field := req.GetField()
	if velo_field == nil {
		return nil, errors.RequestError
	}
	rtn := new(ctum.Result)
	for epoch <= limit {
		velo_field = r.fieldTick(velo_field)
		rtn.History = append(rtn.History, velo_field)
		epoch += 1
	}
	return rtn, nil
}
