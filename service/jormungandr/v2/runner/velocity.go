package runner

import (
	velo "jormungandr/v2/proto/velo"
)

func tick(ent *velo.Entity) *velo.Entity {
	rtn := new(velo.Entity)
	rtn.Id = ent.Id
	rtn.XV = ent.XV + ent.XA
	rtn.YV = ent.YV + ent.YA
	rtn.X = ent.X + ent.XV
	rtn.Y = ent.Y + ent.YV
	return rtn
}

func spaceTick(space *velo.Space) *velo.Space {
	rtn := new(velo.Space)
	for _, ent := range space.Entity {
		rtn.Entity = append(rtn.Entity, tick(ent))
	}
	return rtn
}

func VelocityHandle(req *velo.Request) *velo.Result {
	epoch := 1
	limit := int(req.NestTick)
	space := req.Space
	rtn := new(velo.Result)
	for epoch <= limit {
		space = spaceTick(space)
		rtn.History = append(rtn.History, space)
		epoch += 1
	}
	return rtn
}
