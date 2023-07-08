package runner

import (
	"sync"

	"jormungandr/v2/errors"
	"jormungandr/v2/proto/base"
	ctum "jormungandr/v2/proto/continuum"
)

// 一个系统的执行器的接口
type Tickable interface {
	tick(ent *base.Entity) *base.Entity // 一个实体的tick
}

// 顺序对Field全部的实体进行tick
func SyncFieldTick(self Tickable, space *base.Field) *base.Field {
	rtn := new(base.Field)
	for _, ent := range space.Entity {
		rtn.Entity = append(rtn.Entity, self.tick(ent))
	}
	return rtn
}

// 对一个Field全部的实体多进程并发tick
func AsyncFieldTick(self Tickable, space *base.Field) *base.Field {
	var entities []*base.Entity
	wg := &sync.WaitGroup{}

	mux := new(sync.Mutex)

	for _, ent := range space.Entity {
		wg.Add(1)
		go func(ent *base.Entity) {
			mux.Lock()
			defer mux.Unlock()
			defer wg.Done()
			entities = append(entities, self.tick(ent))
		}(ent)
	}
	wg.Wait()

	return &base.Field{
		Entity: entities,
	}
}

func Handle(self Tickable, req *ctum.Request) (*ctum.Result, error) {
	epoch := 1
	limit := int(req.NestTick)
	fragment_field := req.GetField()
	if fragment_field == nil {
		return nil, errors.RequestError
	}
	rtn := new(ctum.Result)
	for epoch <= limit {
		fragment_field = AsyncFieldTick(self, fragment_field)
		rtn.History = append(rtn.History, fragment_field)
		epoch += 1
	}
	return rtn, nil
}

// 编译时检查
var _ Tickable = new(TimeRunner)
var _ Tickable = new(VelocityRunner)
