package runner

import (
	"jormungandr/v2/proto/base"
	"jormungandr/v2/proto/structure"
)

type structRunner struct {
	// 这里可以添加Runner级别的状态/属性，例如缓存等
}

func (t *structRunner) Tick(s *base.Space) *base.Space {
	// 同步tick伤害系统
	SyncTick(
		s,
		injuryTick,
	)
	// 异步Tick所有结构体
	AsyncTick(
		s,
		structTick,
	)
	return s
}

func NewStructRunner() *structRunner {
	return &structRunner{}
}

// 造成伤害
func doDamage(s *base.Space, one uint64, damage *structure.Health) {
	ent, ok := s.Entity[one]
	if !ok {
		return
	}
	health := ent.Structure.Health
	health.Body -= damage.Body
	health.Armor -= damage.Armor
	health.Shield -= damage.Shield
}

// 结构tick
func structTick(e *base.Entity) {
	// Structure是可选字段
	// 包含可选字段是常态，没有可选字段是特例
	if e.Structure == nil {
		return
	}
	if e.Structure.Destroy {
		return
	}
	if e.Structure.Health.Body <= 0 {
		e.Structure.Destroy = true
		return
	}

	// 生命数值修改
	e.Structure.Health.Body += e.Structure.Delta.HealthA.Body     // 生命值回复
	e.Structure.Health.Armor += e.Structure.Delta.HealthA.Armor   // 生命值回复
	e.Structure.Health.Shield += e.Structure.Delta.HealthA.Shield // 生命值回复

	// 将Delta归零
	e.Structure.Delta.HealthA = &structure.Health{
		Body:   0,
		Armor:  0,
		Shield: 0,
	}

}

// 伤害他人tick
func injuryTick(s *base.Space, e *base.Entity) {
	if e.Structure == nil {
		return
	}

	i := e.Structure.Injury // 伤害属性

	if !i.Injurable {
		return
	}

	switch content := i.Way.(type) {
	case *structure.Injury_Direct:
		doDamage(s, content.Direct, i.Damage)
	case *structure.Injury_Radius:
		targetList, err := GetNearby(s, e.ID, content.Radius)
		if err != nil {
			return
		}
		for _, target := range targetList {
			doDamage(s, target, i.Damage)
		}
	}

}
