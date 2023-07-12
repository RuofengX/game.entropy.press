// 粒度为单个实体的系统的执行器
// 通过同步或并发的方法，对所有entity运行一次tick方法
// 将结果合并并返回
package runner

import (
	"sync"

	"jormungandr/v2/errors"
	"jormungandr/v2/proto/base"
	ctum "jormungandr/v2/proto/continuum"
)

type Runner interface {
	tick(ent *base.Entity) (*base.Entity, deltaGather)
	deltaAdd(one *base.Entity, other *base.Entity) *base.Entity
}

type IsolateRunner interface {
	tick(ent *base.Entity) (*base.Entity)
}

// 顺序对Field全部的实体进行tick
// func SyncFieldTick(self IsolateRunner, space *base.Space) *base.Space {
// 	rtn := new(base.Space)
// 	for _, ent := range space.Entity {
// 		result := self.tick(ent)
// 		rtn.Entity = append(rtn.Entity, result)
// 	}
// 	return rtn
// }

// 对一个Field全部的实体多进程并发tick
func asyncFieldTick(self IsolateRunner, space *base.Space) {
	channel := make(chan *base.Entity, 16)
	var entities []*base.Entity
	wg := &sync.WaitGroup{}

	go func() {
		for ent := range channel {
			entities = append(entities, ent)
			wg.Done()
		}
	}()

	for _, ent := range space.Entity {
		wg.Add(1)
		go func(ent *base.Entity) {
			channel <- self.tick(ent)
		}(ent)
	}
	wg.Wait()
	close(channel)
}

func Handle(runner []Runner, req *ctum.Request) *ctum.Result {
	epoch := 1
	limit := int(req.Iteration)
	space := req.GetSpace()
	if space == nil {
		panic(errors.RequestError)
	}
	rtn := new(ctum.Result)
	wg := new(sync.WaitGroup)

	for epoch <= limit {
		for _, ticker := range runner {
			wg.Add(1)
			go func(t Runner) {
				defer wg.Done()
				switch r := t.(type) {
				case IsolateRunner:
					asyncFieldTick(
						r,
						space,
					)
				case Runner:
					r.fieldTick(space)
				}
			}(ticker)
			// TODO: 对属性那一部分保持不变
			// 对delta那一部分每个ticker.tick之后返回不同的delta，每次迭代后相加
		}
		wg.Wait()
		spaceCopy := space
		rtn.History = append(rtn.History, spaceCopy)
		epoch += 1
	}
	return rtn
}

// 编译时检查
var _ IsolateRunner = new(TimeRunner)
var _ IsolateRunner = new(VelocityRunner)
