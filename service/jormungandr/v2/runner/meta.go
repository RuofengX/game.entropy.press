package runner

import (
	"jormungandr/v2/proto/base"
	"sync"
)

// 对s中的所有实体异步调用f方法
func AsyncTick(s *base.Space, f func(*base.Entity)) {
	wg := new(sync.WaitGroup)
	for _, ent := range s.Entity {
		wg.Add(1)
		go func(ent *base.Entity) {
			f(ent)
			wg.Done()
		}(ent)
	}
	wg.Wait()

}

// 对s中的所有实体有序调用f方法
func SyncTick(s *base.Space, f func(*base.Entity)) {
	for _, ent := range s.Entity {
		f(ent)
	}
}
