package jormungandr

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	"jormungandr/v2/errors"
	ctum "jormungandr/v2/proto/continuum"
	"jormungandr/v2/runner"
)

const (
	/*
		MAX_NESTED_TICK
		单次请求最长向后计算的tick数量
		太长的数量没有意义，因为单个请求的操作是同步的，
		在结果返回前，开头的一部分tick已经失效
	*/
	MAX_NESTED_TICK = 1000
)

// 一个服务代理类，将grpc来流分发给函数计算
type Service struct {
	ctum.UnimplementedContinuumServer

	// 实现有状态的运行器，兼容状态缓存
	// TODO: 实现初始化
	time_runner *runner.TimeRunner
	velo_runner *runner.VelocityRunner
}

// 对请求进行应用级预处理
func (s Service) preParse(in *ctum.Request) error {
	if in.NestTick >= MAX_NESTED_TICK {
		return errors.RequestTickTooBigError
	}
	return nil
}

// 处理时间
func (s Service) TimePass(ctx context.Context, in *ctum.Request) (*ctum.Result, error) {
	err := s.preParse(in)
	if err != nil {
		return nil, err
	}
	return runner.Handle(s.time_runner, in)

}

// 处理位移
func (s Service) VelocityMove(ctx context.Context, in *ctum.Request) (*ctum.Result, error) {
	return runner.Handle(s.velo_runner, in)
}

// ...
// 在这里添加其他接口
// ...

// 启动并阻塞运行
func (s Service) Start(lis net.Listener, use_reflection bool) {
	server := grpc.NewServer()
	ctum.RegisterContinuumServer(server, s)
	if use_reflection {
		reflection.Register(server)
	}
	server.Serve(lis)
}
