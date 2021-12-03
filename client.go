package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"grpc.com/grpc/mangment"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	c := mangment.NewUsermanagementClient(conn)
	var m = map[string]int64{"Srikantha": 22, "anil": 23}
	var id = []int32{1, 2}
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
		fmt.Printf("Id=%d\tName:%s\tAge:%d\tPhone:%d", res1.Id, res1.Name, res1.Age, res1.Phone)
	}

}
