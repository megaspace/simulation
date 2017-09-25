package api

import (
	"fmt"
	"net"

	"github.com/megaspace/messaging/stars"

	"google.golang.org/grpc"
)

func StartServer(port string) (err error) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return
	}

	s := grpc.NewServer()
	stars.RegisterStarServiceServer(s, &StarServiceServer{})

	fmt.Println("Server started...")

	if err = s.Serve(lis); err != nil {
		return
	}

	return
}
