package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"grpc.com/grpc/mangment"
)

type details struct {
	Name  string
	Age   int32
	Phone int64
}

func main() {
	det := []details{
		{Name: "Srikantha", Age: 22, Phone: 8766565722},
		{Name: "Anil", Age: 21, Phone: 8766567622},
	}
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	c := mangment.NewUsermanagementClient(conn)
	/*var m = map[string]int64{"Srikantha": 22, "anil": 23}
	var id = []int32{3, 1}
	for k, v := range m {
		res, err := c.CreateNewUser(context.Background(), &mangment.NewUser{Name: k, Age: v})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Name:%s\n Age:%d \n Id:%s\n", res.Name, res.Age, res.Id)

	}
	for _, i := range id {
		res1, err := c.GetaUser(context.Background(), &mangment.UserId{Id: i})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Id=%d\tName:%s\tAge:%d\tPhone:%d\n", res1.Id, res1.Name, res1.Age, res1.Phone)
	}*/
	for _, i := range det {
		postres, err := c.PostaUser(context.Background(), &mangment.UserInfo{Name: i.Name, Age: i.Age, Phone: i.Phone})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(postres)
		}
	}
	time.Sleep(2 * time.Second)
	getres, err := c.GetaUser(context.Background(), &mangment.UserId{Id: 1})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getres)
	}
	time.Sleep(1 * time.Second)
	getallres, err := c.GetAllUser(context.Background(), &mangment.UserId{Id: 1})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getallres)
	}
	//time.Sleep(2 * time.Second)
	/*delres, err := c.DeleteaUser(context.Background(), &mangment.UserId{Id: 1})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(delres)
	}*/

}
