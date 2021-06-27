package grpc

import (
	"context"
	"github.com/AnyKey/service/user"
	proto "github.com/AnyKey/sub-service/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type grpcServer struct {
	usecase user.Usecase
	proto.SubServer
}

// Launch will create new an server object and register in grpc server
func Launch(s *grpc.Server, uuc user.Usecase) {
	proto.RegisterSubServer(s, &grpcServer{
		usecase: uuc,
	})
}

func (s *grpcServer) GetToken(ctx context.Context, in *proto.GetTokenRequest) (*proto.GetTokenResponse, error) {
	log.Debugf("Received: %v", in.GetToken())
	s.usecase.SetToken(in.Token)
	return &proto.GetTokenResponse{Message: in.GetToken() + " Received"}, nil
}
