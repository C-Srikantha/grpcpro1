package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"grpc.com/grpc/mangment"
)

type server struct {
	mangment.UnimplementedUsermanagementServer
}
type user struct {
	id    int32
	name  string
	age   int32
	phone int64
}

func (se *server) CreateNewUser(ctx context.Context, u *mangment.NewUser) (*mangment.User, error) {
	fmt.Printf("%s called the server\n", u.Name)
	return &mangment.User{Name: u.GetName(), Age: u.GetAge(), Id: "1"}, nil
}
func (se *server) GetaUser(ctx context.Context, i *mangment.UserId) (*mangment.UserInfo, error) {
	us := []user{
		{id: 1, name: "srikantha", age: 15, phone: 98765432122},
		{id: 2, name: "anil", age: 20, phone: 98767624526},
	}
	for _, val := range us {
		if val.id == i.Id {
			return &mangment.UserInfo{Id: val.id, Name: val.name, Age: val.age, Phone: val.phone}, nil
		}
	}
	return nil, errors.New("No user found")
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
