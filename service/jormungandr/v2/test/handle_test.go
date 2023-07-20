package jormungandr_test

import (
	jor "jormungandr/v2"
	"jormungandr/v2/proto/base"
	"jormungandr/v2/proto/time"
	"testing"
)

func TestAdd(t *testing.T) {
    space := &base.Space{
        Entity: make(map[uint64]*base.Entity),
    }
    test_ent := &base.Entity{
        ID: 1,
        Time: &time.Property{
            Age: 0,
            Speed: 0,
            Delta: &time.Delta{
                TimeA: 1,
            },
        },
    }
    space.Entity[1] = test_ent
	result := jor.NewHandler().MultiTick(
        space,
        30,
    )
    println(result[2].Entity[1].Time.Speed)
}

// TODO: 写一个多Delta的测试样例
