package jormungandr

import (
	ctum "jormungandr/v2/proto/continuum"
	"jormungandr/v2/proto/velo"
	"jormungandr/v2/runner"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type ContinuumService struct {
	ctum.UnimplementedContinuumServer
}

func (ct ContinuumService) VelocityCalculate(ctx context.Context, req *velo.Request) (*velo.Result, error) {
	return runner.VelocityHandle(req), nil
}

// ...
// 在这里添加其他接口

func (ct ContinuumService) Start(lis net.Listener) {
	s := grpc.NewServer()
	ctum.RegisterContinuumServer(s, ct)
	s.Serve(lis)
}
