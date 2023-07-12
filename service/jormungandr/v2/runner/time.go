package runner

import (
	"jormungandr/v2/proto/base"
)

type TimeRunner struct {
	// 这里可以添加Runner级别的状态/属性，例如缓存等
}

// 单个实体碎片的tick方法
func (r *TimeRunner) tick(ent *base.Entity) (*base.Entity, ) {
	ent.Time.Age = ent.Time.Age + ent.Time.Speed
	ent.Time.Speed = ent.Time.Speed + ent.TimeD.TimeA
	ent.TimeD.TimeA = 0
	// 不可以创建其他属性字段
	// 可以创建其他D字段
	return ent
}

func (r *TimeRunner) flat() 
