package main

import (
	"context"
	"fmt"
	"net"

	"github.com/go-pg/pg"
	"google.golang.org/grpc"
	"grpc.com/grpc/createtable"
	"grpc.com/grpc/dbcon"
	"grpc.com/grpc/entity"
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

func (se *server) PostaUser(ctx context.Context, info *mangment.UserInfo) (*mangment.Useroutput, error) {
	var det entity.Information
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
	var det entity.Information
	err := db.Model(&det).Where("id=?", userid.Id).Select()
	if err != nil {
		return nil, err
	}
	return &mangment.UserInfo{Name: det.Name, Age: det.Age, Phone: det.Phone}, nil
}
func (se *server) UpdateUser(ctx context.Context, info *mangment.UserInfo) (*mangment.Useroutput, error) {
	var det entity.Information
	det.Id = info.Id
	det.Name = info.Name
	det.Age = info.Age
	det.Phone = info.Phone
	_, err := db.Model(&det).WherePK().Update()
	if err != nil {
		return nil, err
	}
	return &mangment.Useroutput{Mess: "Update Success"}, nil

}
func (se *server) GetAllUser(ctx context.Context, in *mangment.Empty) (*mangment.AllUser, error) {
	var d []*mangment.UserInfo
	var det1 []entity.Information
	err := db.Model(&det1).Select()
	if err != nil {
		return nil, err
	}
	for _, i := range det1 {
		d = append(d, &mangment.UserInfo{Name: i.Name, Age: i.Age, Phone: i.Phone}) //deepcopier
	}
	return &mangment.AllUser{Info: d}, nil
}
func (se *server) DeleteaUser(ctx context.Context, userid *mangment.UserId) (*mangment.Useroutput, error) {
	var det entity.Information
	_, err := db.Model(&det).Where("id=?", userid.Id).Delete()
	if err != nil {
		return &mangment.Useroutput{Mess: "Invalid User Id"}, err
	}
	return &mangment.Useroutput{Mess: "Deleted"}, nil

}
func (se *server) PostCollegeDet(ctx context.Context, info *mangment.CollegeInfo) (*mangment.Useroutput, error) {
	var det entity.CollegeDetails
	det.Id = info.Id
	det.CollegeCode = info.Collagecode
	det.CollegeName = info.Collegename
	det.CollegeLocation = info.Collegelocation
	det.CollegeContactInfo.Phone = info.Contact.Phone
	det.CollegeContactInfo.Email = info.Contact.Email
	_, err := db.Model(&det).Insert()
	if err != nil {
		return nil, err
	}
	return &mangment.Useroutput{Mess: "Success"}, nil
}
func (se *server) GetaCollegeDet(ctx context.Context, infoid *mangment.UserId) (*mangment.CollegeInfo, error) {
	var det entity.CollegeDetails
	err := db.Model(&det).Where("id=?", infoid.Id).Select()
	if err != nil {
		return nil, err
	}
	return &mangment.CollegeInfo{Collagecode: det.CollegeCode, Collegename: det.CollegeName, Collegelocation: det.CollegeLocation,
		Contact: &mangment.Collegecontact{Email: det.CollegeContactInfo.Email, Phone: det.CollegeContactInfo.Phone}}, nil
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
