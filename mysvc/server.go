package main

import (
	"context"
	"fmt"
	"net"

	"github.com/go-pg/pg"
	"google.golang.org/grpc"
	"grpc.com/grpc/createtable"
	"grpc.com/grpc/dbcon"
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

var db *pg.DB

/*func (se *server) CreateNewUser(ctx context.Context, u *mangment.NewUser) (*mangment.User, error) {
	fmt.Printf("%s called the server\n", u.Name)
	return &mangment.User{Name: u.GetName(), Age: u.GetAge(), Id: "1"}, nil
}
func (se *server) GetaUser(ctx context.Context, i *mangment.UserId) (*mangment.UserInfo, error) {
	err := errors.New("No User found")
	us := []user{
		{id: 1, name: "srikantha", age: 15, phone: 98765432122},
		{id: 2, name: "anil", age: 20, phone: 98767624526},
	}
	for _, val := range us {
		if val.id == i.Id {
			return &mangment.UserInfo{Id: val.id, Name: val.name, Age: val.age, Phone: val.phone}, nil
		}
	}
	return nil, err

}*/
var det createtable.Information

func (se *server) PostaUser(ctx context.Context, info *mangment.UserInfo) (*mangment.Useroutput, error) {
	det.Name = info.Name
	det.Age = info.Age
	det.Phone = info.Phone
	_, err := db.Model(&det).Insert()
	if err != nil {
		return nil, err
	}
	return &mangment.Useroutput{Mess: "Success"}, nil
}
func (se *server) GetaUser(ctx context.Context, userid *mangment.UserId) (*mangment.UserInfo, error) {
	err := db.Model(&det).Where("id=?", userid.Id).Select()
	if err != nil {
		return nil, err
	}
	return &mangment.UserInfo{Name: det.Name, Age: det.Age, Phone: det.Phone}, nil
}
func (se *server) DeleteaUser(ctx context.Context, userid *mangment.UserId) (*mangment.Useroutput, error) {
	_, err := db.Model(&det).Where("id=?", userid.Id).Delete()
	if err != nil {
		return &mangment.Useroutput{Mess: "Invalid User Id"}, err
	}
	return &mangment.Useroutput{Mess: "Deleted"}, nil

}

func (se *server) GetAllUser(ctx context.Context, userid *mangment.UserId) (*mangment.AllUser, error) {
	var det1 []createtable.Information
	err := db.Model(&det1).Select()
	if err != nil {
		return nil, err
	}
	for _, i := range det1 {
		mangment.AllUser{Info: []*mangment.UserInfo{{Name: i.Name, Age: i.Age, Phone: i.Phone}}}

	}
}

func main() {
	conn, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}
	db, err = dbcon.Connection()
	if err != nil {
		fmt.Println(err)
	}
	err = createtable.CreateTab(db)
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	mangment.RegisterUsermanagementServer(s, &server{})
	if err := s.Serve(conn); err != nil {
		fmt.Println(err)
	}

}
