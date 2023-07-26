package jormungandr

import (
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"jormungandr/v2/errors"
	ctum "jormungandr/v2/proto/continuum"
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
type continuumServer struct {
	ctum.UnimplementedContinuumServer

	// 实现有状态的运行器，兼容未来可能的状态缓存
	hand *handler
}

func NewService() *continuumServer {
	return &continuumServer{
		hand: NewHandler(),
	}
}

// 对请求进行应用级预处理
func (s *continuumServer) preParse(in *ctum.TickRequest) error {
	if in.Iteration >= MAX_ITERATION {
		return errors.RequestTickTooBigError
	}
	return nil
}

// 对所有runner传入所有信息
func (s *continuumServer) Tick(ctx context.Context, in *ctum.TickRequest) (*ctum.HistoryResult, error) {
	err := s.preParse(in)
	if err != nil {
		return nil, err
	}
	rtn := &ctum.HistoryResult{
		History: s.hand.MultiTick(in.Space, in.Iteration),
	}
	return rtn, nil
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
			println(err.Error())
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
			println(err.Error())
		}
	}()

	// 调用实际的 gRPC 处理程序
	return handler(srv, ss)
}

// 启动并阻塞运行
func (s *continuumServer) Start(lis net.Listener, use_reflection bool) {

	server := grpc.NewServer(
	// grpc.UnaryInterceptor(RecoveryInterceptor),
	// grpc.StreamInterceptor(StreamRecoveryInterceptor),
	)

	ctum.RegisterContinuumServer(server, s)
	if use_reflection {
		reflection.Register(server)
	}

	server.Serve(lis)
}
