package runner

import (
	"jormungandr/v2/errors"
	ctum "jormungandr/v2/proto/continuum"
	meta "jormungandr/v2/proto/meta"
	time "jormungandr/v2/proto/time"
)

type TimeRunner struct{}

// 单个实体碎片的tick方法
func (r *TimeRunner) tick(ent *meta.Entity) *meta.Entity {
	time_ent := ent.Time
	rtn := &(meta.Entity{
		ID: ent.ID,
		Time: &time.Fragment{
			Age:   time_ent.Age + time_ent.Speed,
			Speed: time_ent.Speed + time_ent.Delta.TimeA,
			// Delta: &time.Delta{},  // 这里直接不写，delta对象自动置空
		// 如果没必要，也不去创建同ID的其他属性
		// 如果需要改变其他属性，则通过表象实现，尽可能将实现包装在自己的场中
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
