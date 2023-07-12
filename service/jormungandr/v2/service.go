package jormungandr

import (
	"net"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"jormungandr/v2/errors"
	ctum "jormungandr/v2/proto/continuum"
	"jormungandr/v2/runner"
)

const (
	/*
		MAX_ITERATION
		单次请求最长向后计算的tick数量
		太长的数量没有意义，因为单个请求的操作是同步的，
		在结果返回前，开头的一部分tick已经失效
	*/
	MAX_ITERATION = 1000
)

// 一个服务代理类，将grpc来流分发给函数计算
type Service struct {
	ctum.UnimplementedContinuumServer

	// 实现有状态的运行器，兼容状态缓存
	runner []*runner.Ticker
}

func NewService()(*Service){
	return &Service{
		runner: {
			runner.TimeRunner{}, 
			runner.VelocityRunner{},
		},
	}
}

// 对请求进行应用级预处理
func (s Service) preParse(in *ctum.Request) error {
	if in.Iteration >= MAX_ITERATION {
		return errors.RequestTickTooBigError
	}
	return nil
}

// 对所有runner传入所有信息
func (s Service) Tick(ctx context.Context, in *ctum.Request) (*ctum.Result, error) {
	spaceFragment := new(chan *base.Space)
	wg := sync.WaitGroup{}
	// 对所有runner，分别传入space
	for _, r := range s.runner{
		wg.Add(1)
		go func(r runner.Ticker) {
			defer wg.Done()
			epoch := 1
			limit := int(req.Iteration)
			space := in.Space
			runner.Handle(

			)
		}(r)
	return rtn
}

// ...
// 在这里添加其他接口
// ...

// 中间件函数
func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			// 将 panic 转换为 gRPC 错误
			err := status.Errorf(codes.Internal, "Internal server error")
			panic(err)
		}
	}()

	// 调用实际的 gRPC 处理程序
	return handler(ctx, req)
}

// 流中间件函数
func StreamRecoveryInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	defer func() {
		if r := recover(); r != nil {
			// 将 panic 转换为 gRPC 错误
			err := status.Errorf(codes.Internal, "Internal server error")
			panic(err)
		}
	}()

	// 调用实际的 gRPC 处理程序
	return handler(srv, ss)
}

// 启动并阻塞运行
func (s Service) Start(lis net.Listener, use_reflection bool) {

	server := grpc.NewServer(
		grpc.UnaryInterceptor(RecoveryInterceptor),
		grpc.StreamInterceptor(StreamRecoveryInterceptor),
	)

	ctum.RegisterContinuumServer(server, s)
	if use_reflection {
		reflection.Register(server)
	}

	server.Serve(lis)
}
