package runner

import (
	ctum "jormungandr/v2/proto/continuum"
	meta "jormungandr/v2/proto/meta"
)

// 一个系统的执行器的接口
type Runner interface {
	Handle(req *ctum.Request) (*ctum.Result, error) // 处理请求
	fieldTick(field *meta.Field) *meta.Field        // 一个包含全部实体的场的tick
	tick(ent *meta.Entity) *meta.Entity             // 一个实体的tick
}

// 编译时检查
var _ Runner = new(TimeRunner)
var _ Runner = new(VelocityRunner)
