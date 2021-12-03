package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"grpc.com/grpc/mangment"
)

type server struct {
	mangment.UnimplementedUsermanagementServer
}

func (se *server) CreateNewUser(ctx context.Context, u *mangment.NewUser) (*mangment.User, error) {
	fmt.Printf("%s called the server\n", u.Name)
	return &mangment.User{Name: u.GetName(), Age: u.GetAge(), Id: "1"}, nil
}
func main() {
	conn, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	mangment.RegisterUsermanagementServer(s, &server{})
	if err := s.Serve(conn); err != nil {
		fmt.Println(err)
	}

}
