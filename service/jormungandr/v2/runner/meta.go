package runner

import (
	"jormungandr/v2/proto/base"
	"reflect"
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

// 对f的name字段进行可选的调用，如果name字段不存在，则不调用f方法
// 实验性，暂不使用
func OptionalTick(name string, f func(*base.Entity)) func(*base.Entity) {
	return func(ent *base.Entity) {
		// 判断数值是否存在
		v := reflect.ValueOf(ent)
		if v.FieldByName(name).IsNil() {
			return
		}
		// 正式调用
		f(ent)
	}
}
