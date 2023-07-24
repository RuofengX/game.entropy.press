package jormungandr_test

import (
	jor "jormungandr/v2"
	"jormungandr/v2/proto/base"
	"jormungandr/v2/proto/time"
	"jormungandr/v2/proto/velo"
	"testing"
)

func TestAge(t *testing.T) {
	space := jor.NewEmptySpace()

	for i := 0; i < 10; i++ {
		test_ent := jor.NewEmptyEntity(uint64(i))
		test_ent.Time = &time.Property{
			Age:   0,
			Speed: 0,
			Delta: &time.Delta{
				TimeA: 1,
			},
		}

		space.Entity[uint64(i)] = test_ent
	}

	result := jor.NewHandler().MultiTick(
		space,
		30,
	)
	for i := 0; i < 10; i++ {
		lastSpace := result[len(result)-1]
		resultEntity := lastSpace.Entity[uint64(i)]
		resultAge := resultEntity.Time.Age
		resultSpeed := resultEntity.Time.Speed
		resultTimeA := resultEntity.Time.Delta.TimeA
		if resultAge != 30 {
			t.Errorf("不正确的年龄 %d", resultAge)
		}
		if resultSpeed != 1 {
			t.Errorf("不正确的时间流逝 %d", resultSpeed)
		}
		if resultTimeA != 0 {
			t.Errorf("不正确的时间变换速度 %d", resultTimeA)
		}
	}
}

func TestVelo(t *testing.T) {
	space := &base.Space{
		Entity: make(map[uint64]*base.Entity),
	}

	// 创建测试对象
	for i := 0; i < 10; i++ {
		ID := uint64(i)
		test_ent := jor.NewEmptyEntity(ID)
		test_ent.ID = ID
		test_ent.Velo = &velo.Property{
			X:  0,
			Y:  1,
			XV: 1,
			YV: -1,
			Delta: &velo.Delta{
				XA: 0,
				YA: -1000,
			},
		}

		space.Entity[ID] = test_ent
	}
	result := jor.NewHandler().MultiTick(
		space,
		30,
	)
	targetX := float32(30)
	targetY := float32(-30029)

	for i := 0; i < 10; i++ {
		ID := uint64(i)
		resultVelo := result[len(result)-1].Entity[ID].GetVelo()

		if resultVelo.X != targetX || resultVelo.Y != targetY {
			t.Errorf("不正确的位置 %f,%f", resultVelo.X, resultVelo.Y)
            // FIXME: BUG here
		}
	}
}
