package jormungandr

import (
	"jormungandr/v2/proto/base"
	"jormungandr/v2/proto/time"
	"jormungandr/v2/proto/velo"
)

func NewEmptySpace() *base.Space {
	return &base.Space{
		Entity: make(map[uint64]*base.Entity),
	}
}

func NewEmptyEntity(ID uint64) *base.Entity {
	return &base.Entity{
		ID: ID,
		Time: &time.Property{
			Age:   0,
			Speed: 0,
			Delta: &time.Delta{
				TimeA: 1,
			},
		},
		Velo: &velo.Property{
			X:  0,
			Y:  0,
			XV: 0,
			YV: 0,
			Delta: &velo.Delta{
				XA: 0,
				YA: 0,
			},
		},
	}
}
