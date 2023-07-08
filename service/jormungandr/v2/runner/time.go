package runner

import (
	"jormungandr/v2/proto/base"
	"jormungandr/v2/proto/time"
)

type TimeRunner struct {
	Tickable
	// 这里可以添加Runner级别的状态/属性，例如缓存等
}

// 单个实体碎片的tick方法
func (r *TimeRunner) tick(ent *base.Entity) *base.Entity {
	time_ent := ent.Time
	rtn := &(base.Entity{
		ID: ent.ID,
		Time: &time.Fragment{
			Age:   time_ent.Age + time_ent.Speed,
			Speed: time_ent.Speed + time_ent.GetDelta().GetTimeA(),
			Delta: &time.Delta{}, // 这里如果不写，delta对象**不会**自动置空
			// 消极创建其他碎片字段
		},
	})
	return rtn
}
