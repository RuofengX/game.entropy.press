package runner

import (
	"fmt"
	"math"
	"jormungandr/v2/proto/base"
)

type veloRunner struct {
	// 这里可以添加Runner级别的状态/属性，例如缓存等
}

func (t *veloRunner) Tick(s *base.Space) *base.Space {
	AsyncTick(
		s,
		entTick,
	)
	SyncTick(
		s,
		gridTick,
	)
	return s
}

func NewVeloRunner() *veloRunner {
	return &veloRunner{}
}

// 实体本身tick
func entTick(ent *base.Entity) {

	// 更新速度信息
	ent.Velo.XV += ent.Velo.Delta.XA
	ent.Velo.YV += ent.Velo.Delta.YA
	ent.Velo.X += ent.Velo.XV
	ent.Velo.Y += ent.Velo.YV
	ent.Velo.Delta.XA = 0
	ent.Velo.Delta.YA = 0

	// 更新chunkIndex
	cx := unify(ent.Velo.X)
	cy := unify(ent.Velo.Y)
	ent.Velo.ChunkIndex = fmt.Sprintf("%d,%d", cx,cy)
}

// 将ent.Velo.ChunkIndex更新至s.Grid中去
func gridTick(s *base.Space, ent *base.Entity) {
	if ent.Velo == nil{
		return
	}

	ci := ent.Velo.ChunkIndex
	if s.Grid == nil{
		// 宇宙的Grid为空，意味着之前不存在有Velo属性的实体
		s.Grid = make(map[string]*base.Chunk)
	}

	chunk, ok := s.Grid[ci]
	if !ok {
		s.Grid[ci] = &base.Chunk{
			EntityIndex: []uint64{ent.ID},
		}
	} else {
		chunk.EntityIndex = append(chunk.EntityIndex, ent.ID)
	}
}

func unify(f float64) int64 {
	return int64(math.Floor(f / 1024))
}
