package api

import (
	"golang.org/x/net/context"

	proto "github.com/megaspace/messaging/stars"
)

type StarServiceServer struct{}

func (s *StarServiceServer) GetStars(ctx context.Context, req *proto.GetStarsRequest) (res *proto.GetStarsResponse, err error) {
	res = new(proto.GetStarsResponse)
	return
}
