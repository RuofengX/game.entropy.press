package runner

import (
	"jormungandr/v2/errors"
	ctum "jormungandr/v2/proto/continuum"
	meta "jormungandr/v2/proto/meta"
	time "jormungandr/v2/proto/time"
)

type TimeRunner struct{}

func (r *TimeRunner) tick(ent *meta.Entity) *meta.Entity {
	time_ent := ent.GetTime()
	rtn := &(meta.Entity{
		Id: ent.Id,
		Field: &meta.Entity_Time{
			Time: &time.Time{
				Age:   time_ent.Age + time_ent.Speed,
				Speed: time_ent.Speed,
			},
		},
	})
	return rtn
}

func (r *TimeRunner) fieldTick(space *meta.Field) *meta.Field {
	rtn := new(meta.Field)
	for _, ent := range space.Entity {
		rtn.Entity = append(rtn.Entity, r.tick(ent))
	}
	return rtn
}

func (r *TimeRunner) Handle(req *ctum.Request) (*ctum.Result, error) {
	epoch := 1
	limit := int(req.NestTick)
	time_field := req.GetField()
	if time_field == nil {
		return nil, errors.RequestError
	}
	rtn := new(ctum.Result)
	for epoch <= limit {
		time_field = r.fieldTick(time_field)
		rtn.History = append(rtn.History, time_field)
		epoch += 1
	}
	return rtn, nil
}
