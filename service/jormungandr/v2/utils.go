package jormungandr

import (
	"jormungandr/v2/proto/base"
	"jormungandr/v2/proto/structure"
	"jormungandr/v2/proto/time"
	"jormungandr/v2/proto/velo"
)

func NewEmptySpace() *base.Space {
	rtn := &base.Space{
		Entity: make(map[uint64]*base.Entity),
		Grid: make(map[string]*base.Chunk),
	}
	return rtn
}

func NewEmptyEntity(ID uint64) *base.Entity {
	rtn := &base.Entity{
		ID: ID,
		Time: &time.Property{
			Age:   0,
			Speed: 0,
			Delta: &time.Delta{
				TimeA: 1,
			},
		},
		Velo: &velo.Property{
			Delta: new(velo.Delta),
			X:  0,
			Y:  0,
			XV: 0,
			YV: 0,
			ChunkIndex: "0,0",
		},
		Structure: &structure.Property{
			Destroy: false,
			Health: new(structure.Health),
			Injury: &structure.Injury{
				Injurable: false,
				Damage: new(structure.Health),
				Way: &structure.Injury_Direct{
					Direct: 0,
				},
			},
			Delta: &structure.Delta{
				HealthA: &structure.Health{
				Body:   0,
				Armor:  0,
				Shield: 0,
			},
			},
		},
	}

	return rtn
}
